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

		var actualSubmission pkgs.SnapshotSubmission
		err = json.Unmarshal(msg.Data, &actualSubmission)
		if err != nil {
			log.Errorf("Error unmarshalling submission on topic %s: %v", sub.Topic(), err)
			continue
		}

		// Use a unique ID for the submission in the queue to avoid duplicates if needed,
		// here we are just using the message ID.
		submissionID := msg.ID

		// Add submission to Redis queue
		queueData := map[string]interface{}{
			"submission_id":       submissionID,
			"data_market_address": actualSubmission.DataMarket,
			"data":                string(msg.Data),
		}
		queueDataJSON, err := json.Marshal(queueData)
		if err != nil {
			log.Errorf("Error marshalling queue data for topic %s: %v", sub.Topic(), err)
			continue
		}

		err = redis.RedisClient.LPush(context.Background(), "submissionQueue", queueDataJSON).Err()
		if err != nil {
			log.Errorf("Error adding to Redis queue for topic %s: %v", sub.Topic(), err)
			continue
		}
		log.Infof("Queued snapshot from topic %s: %s", sub.Topic(), submissionID)

		// Increment submission count
		count, err := redis.Incr(context.Background(), redis.EpochSubmissionCountsReceivedInSlotKey(actualSubmission.DataMarket, actualSubmission.Request.SlotId, actualSubmission.Request.EpochId))
		if err != nil {
			log.Errorf("Error incrementing submission count for topic %s: %v", sub.Topic(), err)
		}
		log.Infof("Submission count for slot %d and epoch %d on topic %s is %d", actualSubmission.Request.SlotId, actualSubmission.Request.EpochId, sub.Topic(), count)
		redis.RedisClient.Expire(context.Background(), redis.EpochSubmissionCountsReceivedInSlotKey(actualSubmission.DataMarket, actualSubmission.Request.SlotId, actualSubmission.Request.EpochId), 5*time.Minute)
	}
}
