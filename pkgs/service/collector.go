package service

import (
	"Listen/pkgs"
	"Listen/pkgs/redis"
	"context"
	"encoding/json"
	"time"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
	log "github.com/sirupsen/logrus"
)

// P2PSnapshotSubmission represents the data structure for snapshot submissions
// sent over the P2P network by the collector.
type P2PSnapshotSubmission struct {
	EpochID       uint64                     `json:"epoch_id"`
	Submissions   []*pkgs.SnapshotSubmission `json:"submissions"`
	SnapshotterID string                     `json:"snapshotter_id"`
	Signature     []byte                     `json:"signature"`
}

func GossipsubMessageHandler(ctx context.Context, sub *pubsub.Subscription) {
	for {
		msg, err := sub.Next(ctx)
		if err != nil {
			log.Errorf("Error getting next message from topic %s: %v", sub.Topic(), err)
			continue
		}
		log.Infof("GossipsubMessageHandler: Received a message from topic %s", sub.Topic())

		

		log.Infof("Received message from %s on topic %s", msg.GetFrom(), sub.Topic())
		log.Infof("RAW_MESSAGE_DATA: %s", string(msg.Data))

		// First try to unmarshal as P2PSnapshotSubmission (from collector)
		var p2pSubmission P2PSnapshotSubmission
		err = json.Unmarshal(msg.Data, &p2pSubmission)
		if err != nil {
			// If that fails, try direct SnapshotSubmission for backward compatibility
			var actualSubmission pkgs.SnapshotSubmission
			err = json.Unmarshal(msg.Data, &actualSubmission)
			if err != nil {
				log.Errorf("Error unmarshalling submission on topic %s: %v", sub.Topic(), err)
				log.Errorf("Raw data that failed to unmarshal: %s", string(msg.Data))
				continue
			}
			// Process single submission
			processSubmission(ctx, &actualSubmission, msg, sub.Topic())
			continue
		}

		// Process each submission in the P2P message
		log.Infof("Processing P2P submission with %d submissions from snapshotter %s", 
			len(p2pSubmission.Submissions), p2pSubmission.SnapshotterID)
		
		for _, actualSubmission := range p2pSubmission.Submissions {
			processSubmission(ctx, actualSubmission, msg, sub.Topic())
		}
	}
}

func processSubmission(ctx context.Context, actualSubmission *pkgs.SnapshotSubmission, msg *pubsub.Message, topicName string) {
	// Use a unique ID for the submission in the queue to avoid duplicates if needed,
	// here we are just using the message ID.
	submissionID := string(msg.ID)

	// Marshal the actual submission data (not the original message data)
	submissionData, err := json.Marshal(actualSubmission)
	if err != nil {
		log.Errorf("Error marshalling submission data for topic %s: %v", topicName, err)
		return
	}

	// Add submission to Redis queue
	queueData := map[string]interface{}{
		"submission_id":       submissionID,
		"data_market_address": actualSubmission.DataMarket,
		"data":                string(submissionData),
	}
	queueDataJSON, err := json.Marshal(queueData)
	if err != nil {
		log.Errorf("Error marshalling queue data for topic %s: %v", topicName, err)
		return
	}

	err = redis.RedisClient.LPush(context.Background(), "submissionQueue", queueDataJSON).Err()
	if err != nil {
		log.Errorf("Error adding to Redis queue for topic %s: %v", topicName, err)
		return
	}
	log.Infof("Queued snapshot from topic %s: %s", topicName, submissionID)

	// Increment submission count
	count, err := redis.Incr(context.Background(), redis.EpochSubmissionCountsReceivedInSlotKey(actualSubmission.DataMarket, actualSubmission.Request.SlotId, actualSubmission.Request.EpochId))
	if err != nil {
		log.Errorf("Error incrementing submission count for topic %s: %v", topicName, err)
	}
	log.Infof("Submission count for slot %d and epoch %d on topic %s is %d", actualSubmission.Request.SlotId, actualSubmission.Request.EpochId, topicName, count)
	redis.RedisClient.Expire(context.Background(), redis.EpochSubmissionCountsReceivedInSlotKey(actualSubmission.DataMarket, actualSubmission.Request.SlotId, actualSubmission.Request.EpochId), 5*time.Minute)
}
