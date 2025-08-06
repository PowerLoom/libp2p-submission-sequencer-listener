package service

import (
	"Listen/config"
	"context"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	pubsub "github.com/libp2p/go-libp2p-pubsub"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-libp2p/p2p/discovery/util"
	log "github.com/sirupsen/logrus"
)

// GossipsubManager manages pub-sub subscriptions and message handling.

type GossipsubManager struct {
	host        host.Host
	pubsub      *pubsub.PubSub
	redisClient *redis.Client
	dht         *dht.IpfsDHT
	topics      map[string]*pubsub.Subscription
	mu          sync.Mutex
}

// NewGossipsubManager creates a new manager for gossip subscriptions.
func NewGossipsubManager(ctx context.Context, cfg *config.Settings, redisClient *redis.Client, h host.Host, kademliaDHT *dht.IpfsDHT) (*GossipsubManager, error) {
	// Configure gossipsub with discovery and flood publishing for better peer discovery
	ps, err := pubsub.NewGossipSub(
		ctx, 
		h,
		pubsub.WithDiscovery(routing.NewRoutingDiscovery(kademliaDHT)),
		pubsub.WithFloodPublish(true), // Ensure messages reach all peers
		pubsub.WithDirectPeers([]peer.AddrInfo{}), // Can be populated with known peers
	)
	if err != nil {
		return nil, err
	}

	return &GossipsubManager{
		host:        h,
		pubsub:      ps,
		redisClient: redisClient,
		dht:         kademliaDHT,
		topics:      make(map[string]*pubsub.Subscription),
	}, nil
}

// Start begins the Redis subscription and message handling loop.
func (m *GossipsubManager) Start(ctx context.Context) {
	// Join the two-level topic architecture immediately
	// 1. Discovery topic (epoch 0) for peer finding
	m.joinTopic(ctx, "/powerloom/snapshot-submissions/0")
	
	// 2. Main submissions topic for all epochs
	m.joinTopic(ctx, "/powerloom/snapshot-submissions/all")
	
	log.Info("Two-level topic architecture initialized: discovery (epoch 0) and submissions (all)")
	
	// Still listen to Redis for any additional topics if needed
	redisSub := m.redisClient.Subscribe(ctx, "epoch-topics")
	defer redisSub.Close()

	log.Info("Listening for additional epoch topics from Redis...")

	for {
		msg, err := redisSub.ReceiveMessage(ctx)
		if err != nil {
			log.Errorf("Error receiving message from Redis: %v", err)
			continue
		}

		topicName := msg.Payload
		// Skip if it's one of our standard topics
		if topicName == "/powerloom/snapshot-submissions/0" || 
		   topicName == "/powerloom/snapshot-submissions/all" {
			continue
		}
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

	// Advertise ourselves on this specific topic for discovery
	go func() {
		routingDiscovery := routing.NewRoutingDiscovery(m.dht)
		// Advertise on the topic name as the rendezvous point
		// This allows publishers to find us when they're about to publish to this topic
		log.Infof("Advertising on rendezvous point: %s", topicName)
		util.Advertise(ctx, routingDiscovery, topicName)
		log.Infof("Successfully advertised on topic-specific rendezvous: %s", topicName)
	}()

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