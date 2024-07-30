package service

import (
	"Listen/pkgs/redis"
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p/core/network"
	log "github.com/sirupsen/logrus"
	"io"
)

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
		submissionId := uuid.New()
		err = submissionId.UnmarshalText(buf[:36])
		if err != nil {
			log.Debugln("Unable to unmarshal uuid for submission: ", string(buf[:36]))
			return
		}

		// Add submission to Redis queue
		queueData := map[string]interface{}{
			"submission_id": string(buf[:36]),
			"data":          string(buf[36:length]),
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
	}
}

func StartCollectorServer() {
	collectorHost.SetStreamHandler("/collect", handleStream)
}
