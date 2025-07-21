package watcher

import (
	"Listen/config"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-redis/redis/v8"
	log "github.com/sirupsen/logrus"
)

// ContractArtifact is the structure of the JSON artifact file
type ContractArtifact struct {
	ABI json.RawMessage `json:"abi"`
}

// Watcher struct holds the state for the watcher service
type Watcher struct {
	config      config.Settings
	ethClient   *ethclient.Client
	redisClient *redis.Client
	contractABI abi.ABI
}

// NewWatcher creates and initializes a new Watcher service
func NewWatcher(cfg config.Settings) (*Watcher, error) {
	client, err := ethclient.Dial(cfg.ClientUrl)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to Ethereum node: %w", err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: cfg.RedisHost + ":" + cfg.RedisPort,
	})

	abiBytes, err := ioutil.ReadFile(cfg.ABIFilePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read ABI file: %w", err)
	}

	var artifact ContractArtifact
	if err := json.Unmarshal(abiBytes, &artifact); err != nil {
		return nil, fmt.Errorf("failed to unmarshal contract artifact: %w", err)
	}

	contractABI, err := abi.JSON(strings.NewReader(string(artifact.ABI)))
	if err != nil {
		return nil, fmt.Errorf("failed to parse contract ABI: %w", err)
	}

	return &Watcher{
		config:      cfg,
		ethClient:   client,
		redisClient: rdb,
		contractABI: contractABI,
	}, nil
}

// Start begins the message handling loop for the watcher
func (w *Watcher) Start(ctx context.Context) {
	log.Info("Watcher started. Waiting for EpochReleased events...")

	if w.config.OnlyEpoch0 {
		// Publish epoch 0 for simulation submissions and exit
		epoch0Topic := fmt.Sprintf("/powerloom/snapshot-submissions/%d", 0)
		log.Infof("ONLY_EPOCH_0 is true. Publishing static topic for epoch 0 to Redis: %s", epoch0Topic)
		err := w.redisClient.Publish(ctx, "epoch-topics", epoch0Topic).Err()
		if err != nil {
			log.Errorf("Failed to publish epoch 0 topic to Redis: %v", err)
		}
		log.Info("ONLY_EPOCH_0 is true. Watcher will not poll for further events.")
		return
	}

	contractAddress := common.HexToAddress(w.config.ContractAddress)
	lastBlock, err := w.ethClient.BlockNumber(ctx)
	if err != nil {
		log.Fatalf("Failed to get latest block number: %v", err)
	}
	log.Infof("Starting to poll for logs from block %d", lastBlock)

	ticker := time.NewTicker(15 * time.Second) // Poll every 15 seconds
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			latestBlock, err := w.ethClient.BlockNumber(ctx)
			if err != nil {
				log.Errorf("Failed to get latest block number: %v", err)
				continue
			}

			if latestBlock <= lastBlock {
				continue
			}

			query := ethereum.FilterQuery{
				FromBlock: new(big.Int).SetUint64(lastBlock + 1),
				ToBlock:   new(big.Int).SetUint64(latestBlock),
				Addresses: []common.Address{contractAddress},
			}

			logs, err := w.ethClient.FilterLogs(ctx, query)
			if err != nil {
				log.Errorf("Failed to filter logs: %v", err)
				continue
			}

			for _, vLog := range logs {
				w.handleLog(ctx, vLog)
			}

			lastBlock = latestBlock

		case <-ctx.Done():
			log.Info("Stopping watcher...")
			return
		}
	}
}

func (w *Watcher) handleLog(ctx context.Context, vLog types.Log) {
	epochReleasedEventID := w.contractABI.Events["EpochReleased"].ID

	if vLog.Topics[0] != epochReleasedEventID {
		return // Not an EpochReleased event
	}

	epochReleasedEvent := struct {
		Begin     *big.Int
		End       *big.Int
		Timestamp *big.Int
	}{}

	err := w.contractABI.UnpackIntoInterface(&epochReleasedEvent, "EpochReleased", vLog.Data)
	if err != nil {
		log.Errorf("Failed to unpack EpochReleased event: %v", err)
		return
	}

	// The epochId is the second indexed topic
	epochId := vLog.Topics[2].Big()

	topic := fmt.Sprintf("/powerloom/snapshot-submissions/%s", epochId.String())
	log.Infof("New epoch released: %s. Publishing topic to Redis: %s", epochId.String(), topic)

	err = w.redisClient.Publish(ctx, "epoch-topics", topic).Err()
	if err != nil {
		log.Errorf("Failed to publish topic to Redis: %v", err)
	}
}