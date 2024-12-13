package service

import (
	"Listen/config"
	"Listen/pkgs"
	"context"
	"encoding/json"
	"io"
	"time"

	"github.com/libp2p/go-libp2p/core/network"
	log "github.com/sirupsen/logrus"

	"Listen/pkgs/redis"
)

func parseSubmissionBytes(data []byte) (uuid string, submission []byte, err error) {
	currentPos := 0
	uuid = string(data[currentPos : currentPos+36])
	currentPos += 36
	log.Debugln("Data market address found for submission with ID: ", uuid)

	// Rest is the submission JSON
	submission = data[currentPos:]

	return uuid, submission, nil
}

func handleStream(stream network.Stream) {
	defer stream.Close()
	for {
		buf := make([]byte, 1024)
		length, err := stream.Read(buf)
		if err != nil {
			if err == io.EOF {
				log.Debugln("End of stream reached")
				break
			}
			log.Debugln("Error reading:", err)
			return
		}

		submissionID, submission, err := parseSubmissionBytes(buf[:length])
		if err != nil {
			log.Debugln("Unable to parse submission: ", err)
			return
		}
		var actualSubmission pkgs.SnapshotSubmission
		err = json.Unmarshal(submission, &actualSubmission)
		if err != nil {
			log.Debugln("Error unmarshalling submission", err, "with body: ", string(submission))
			continue
		}
		// Add submission to Redis queue
		queueData := map[string]interface{}{
			"submission_id":       submissionID,
			"data_market_address": actualSubmission.DataMarket,
			"data":                string(submission),
		}
		queueDataJSON, err := json.Marshal(queueData)
		if err != nil {
			log.Debugln("Error marshalling queue data:", err)
			continue
		}

		err = redis.RedisClient.LPush(context.Background(), "submissionQueue", queueDataJSON).Err()
		if err != nil {
			log.Debugln("Error adding to Redis queue:", err)
			continue
		}
		log.Debugln("Queued snapshot: ", queueData)
		
		count, err := redis.Incr(context.Background(), redis.EpochSubmissionCountsReceivedInSlotKey(address, actualSubmission.Request.SlotId, actualSubmission.Request.EpochId))
		if err != nil {
			log.Debugln("Error incrementing submission count", err)
		}
		log.Debugln("Submission count for slot", actualSubmission.Request.SlotId, "and epoch", actualSubmission.Request.EpochId, "is", count)
		redis.RedisClient.Expire(context.Background(), redis.EpochSubmissionCountsReceivedInSlotKey(address, actualSubmission.Request.SlotId, actualSubmission.Request.EpochId), 5*time.Minute)
	}
}

func StartCollectorServer() {
	if RelayerHost == nil {
		log.Fatal("RelayerHost is not initialized. Make sure ConfigureRelayer() is called before StartCollectorServer()")
	}
	RelayerHost.SetStreamHandler("/collect", handleStream)
}
