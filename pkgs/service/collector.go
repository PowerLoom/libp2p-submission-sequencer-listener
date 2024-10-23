package service

import (
	"context"
	"encoding/json"
	"io"

	"github.com/ethereum/go-ethereum/common"
	"github.com/google/uuid"
	"github.com/libp2p/go-libp2p/core/network"
	log "github.com/sirupsen/logrus"

	"Listen/pkgs/redis"
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

		submissionID, err := uuid.ParseBytes(buf[:36])
		if err != nil {
			log.Debugln("Unable to parse UUID for submission: ", string(buf[:36]))
			return
		}

		// Extract data market address
		// EVM address is a fixed length of 42 characters (0x followed by 40 hex digits)
		dataMarketAddress := common.HexToAddress(string(buf[36:78])).Hex()

		// Add submission to Redis queue
		queueData := map[string]interface{}{
			"submission_id":       submissionID.String(),
			"data_market_address": dataMarketAddress,
			"data":                string(buf[78:length]),
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
