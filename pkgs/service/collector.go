package service

import (
	"Listen/config"
	"context"
	"encoding/json"
	"errors"
	"io"

	"github.com/libp2p/go-libp2p/core/network"
	log "github.com/sirupsen/logrus"

	"Listen/pkgs/redis"
)

func parseSubmissionBytes(data []byte) (address, uuid string, submission []byte, err error) {
	currentPos := 0
	// If DataMarketInRequest is true
	if config.SettingsObj.DataMarketInRequest {
		if len(data) < 42 {
			return "", "", nil, errors.New("data too short for address")
		}
		address = string(data[:42])
		currentPos = 42
	} else {
		address = config.SettingsObj.DataMarketAddress
	}

	// Parse UUID
	if len(data[currentPos:]) < 36 {
		return "", "", nil, errors.New("data too short for UUID")
	}
	uuid = string(data[currentPos : currentPos+36])
	currentPos += 36
	log.Debugln("Data market address found for submission with ID: ", uuid, " and address: ", address)

	// Rest is the submission JSON
	submission = data[currentPos:]

	return address, uuid, submission, nil
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

		address, submissionID, submission, err := parseSubmissionBytes(buf[:length])
		if err != nil {
			log.Debugln("Unable to parse submission: ", err)
			return
		}
		// Add submission to Redis queue
		queueData := map[string]interface{}{
			"submission_id":       submissionID,
			"data_market_address": address,
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
	}
}

func StartCollectorServer() {
	if RelayerHost == nil {
		log.Fatal("RelayerHost is not initialized. Make sure ConfigureRelayer() is called before StartCollectorServer()")
	}
	RelayerHost.SetStreamHandler("/collect", handleStream)
}
