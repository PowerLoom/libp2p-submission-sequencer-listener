package service

import (
	"Listen/config"
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/libp2p/go-libp2p/core/host"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	log "github.com/sirupsen/logrus"
)

// GossipsubManager manages pub-sub subscriptions and message handling.

type GossipsubManager struct {
	host        host.Host
	pubsub      *pubsub.PubSub
	redisClient *redis.Client
	topics      map[string]*pubsub.Subscription
	mu          sync.Mutex
}

// NewGossipsubManager creates a new manager for gossip subscriptions.
func NewGossipsubManager(ctx context.Context, cfg *config.Settings, redisClient *redis.Client, h host.Host, kademliaDHT *dht.IpfsDHT) (*GossipsubManager, error) {
	ps, err := pubsub.NewGossipSub(ctx, h, pubsub.WithDiscovery(routing.NewRoutingDiscovery(kademliaDHT)))
	if err != nil {
		return nil, err
	}

	return &GossipsubManager{
		host:        h,
		pubsub:      ps,
		redisClient: redisClient,
		topics:      make(map[string]*pubsub.Subscription),
	}, nil
}

// Start begins the Redis subscription and message handling loop.
func (m *GossipsubManager) Start(ctx context.Context) {
	redisSub := m.redisClient.Subscribe(ctx, "epoch-topics")
	defer redisSub.Close()

	log.Info("Listening for new epoch topics from Redis...")

	for {
		msg, err := redisSub.ReceiveMessage(ctx)
		if err != nil {
			log.Errorf("Error receiving message from Redis: %v", err)
			continue
		}

		topicName := msg.Payload
		m.joinTopic(ctx, topicName)
	}
}

// joinTopic joins a new pub-sub topic and starts a message handler goroutine.
func (m *GossipsubManager) joinTopic(ctx context.Context, topicName string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.topics[topicName]; exists {
		log.Infof("Already subscribed to topic: %s", topicName)
		return
	}

	topic, err := m.pubsub.Join(topicName)
	if err != nil {
		log.Errorf("Failed to join topic %s: %v", topicName, err)
		return
	}

	sub, err := topic.Subscribe()
	if err != nil {
		log.Errorf("Failed to subscribe to topic %s: %v", topicName, err)
		return
	}

	m.topics[topicName] = sub
	go GossipsubMessageHandler(ctx, sub)

	log.Infof("Successfully joined and subscribed to topic: %s. Host ID: %s", topicName, m.host.ID())

	// Start periodic diagnostic logging for the new topic
	go func() {
		ticker := time.NewTicker(30 * time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				peersInTopic := m.pubsub.ListPeers(topicName)
				log.WithFields(log.Fields{
					"topic":      topicName,
					"peer_count": len(peersInTopic),
					"peers":      peersInTopic,
				}).Info("DIAGNOSTIC: Periodic check of peers in topic")
			case <-ctx.Done():
				return
			}
		}
	}()
}