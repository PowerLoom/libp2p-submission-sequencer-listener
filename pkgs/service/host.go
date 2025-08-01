package service

import (
	"Listen/config"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/core/routing"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/p2p/discovery/util"
	routing_discovery "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-libp2p/p2p/net/connmgr"
	rcmgr "github.com/libp2p/go-libp2p/p2p/host/resource-manager"
	"github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
)

// NewHost creates a new libp2p host and connects to bootstrap peers.
func NewHost(ctx context.Context, bootstrapPeers string, listenerPort string) (h host.Host, kademliaDHT *dht.IpfsDHT, err error) {
	listenAddr := fmt.Sprintf("/ip4/0.0.0.0/tcp/%s", listenerPort)

	// 1. Create a new resource manager with custom limits.
	scalingLimits := rcmgr.DefaultLimits
	libp2p.SetDefaultServiceLimits(&scalingLimits)

		limits := rcmgr.ResourceLimits{
		StreamsOutbound: rcmgr.Unlimited,
		StreamsInbound:  rcmgr.Unlimited,
		Streams:         rcmgr.Unlimited,
		Conns:           rcmgr.Unlimited,
		ConnsOutbound:   rcmgr.Unlimited,
		ConnsInbound:    rcmgr.Unlimited,
		FD:              rcmgr.Unlimited,
		Memory:          rcmgr.LimitVal64(rcmgr.Unlimited),
	}

	cfg := rcmgr.PartialLimitConfig{
		System:    limits,
		Transient: limits,
	}

	limiter := rcmgr.NewFixedLimiter(cfg.Build(scalingLimits.AutoScale()))
	rscMgr, err := rcmgr.NewResourceManager(limiter, rcmgr.WithMetricsDisabled())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create resource manager: %w", err)
	}

	// 2. Create the connection manager
	cm, err := connmgr.NewConnManager(
		100, // Lowwater
		400, // Highwater
		connmgr.WithGracePeriod(time.Minute),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create connection manager: %w", err)
	}

	// 3. Parse bootstrap peers
	bootstrapAddrInfos, err := parseBootstrapPeers(bootstrapPeers)
	if err != nil {
		log.Warnf("Failed to parse bootstrap peers: %v", err)
	}

	var kadDHT *dht.IpfsDHT
	// 4. Build the libp2p host
	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(listenAddr),
		libp2p.ResourceManager(rscMgr),
		libp2p.ConnectionManager(cm),
		libp2p.Routing(func(h host.Host) (routing.PeerRouting, error) {
			var err error
			kadDHT, err = dht.New(ctx, h, dht.Mode(dht.ModeClient), dht.BootstrapPeers(bootstrapAddrInfos...))
			return kadDHT, err
		}),
		libp2p.EnableRelay(),
		libp2p.EnableHolePunching(),
	}

	if config.SettingsObj.PublicIP != "" {
		publicAddr, err := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/%s/tcp/%s", config.SettingsObj.PublicIP, listenerPort))
		if err != nil {
			log.Errorf("Failed to create public multiaddr: %v", err)
		} else {
			opts = append(opts, libp2p.AddrsFactory(func(addrs []multiaddr.Multiaddr) []multiaddr.Multiaddr {
				return append(addrs, publicAddr)
			}))
		}
	}

	h, err = libp2p.New(opts...)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create libp2p host: %w", err)
	}

	// 5. Bootstrap the DHT in the background
	go func() {
		log.Info("Starting DHT bootstrap process...")
		if err := kadDHT.Bootstrap(ctx); err != nil {
			log.Errorf("DHT bootstrap failed: %v", err)
		} else {
			log.Info("DHT bootstrap completed.")
		}
	}()

	// 6. Announce our presence
	go func() {
		log.Info("Starting rendezvous announcement loop...")
		routingDiscovery := routing_discovery.NewRoutingDiscovery(kadDHT)
		util.Advertise(ctx, routingDiscovery, config.SettingsObj.RendezvousPoint)
	}()

	log.Infof("Libp2p host created with ID: %s", h.ID())
	return h, kadDHT, nil
}

func parseBootstrapPeers(peers string) ([]peer.AddrInfo, error) {
	if peers == "" {
		return nil, nil
	}
	peerStrings := strings.Split(peers, ",")
	addrInfos := make([]peer.AddrInfo, 0, len(peerStrings))
	for _, peerString := range peerStrings {
		addr, err := multiaddr.NewMultiaddr(peerString)
		if err != nil {
			return nil, fmt.Errorf("failed to parse multiaddr '%s': %w", peerString, err)
		}
		peerInfo, err := peer.AddrInfoFromP2pAddr(addr)
		if err != nil {
			return nil, fmt.Errorf("failed to get peer info from multiaddr '%s': %w", peerString, err)
		}
		addrInfos = append(addrInfos, *peerInfo)
	}
	return addrInfos, nil
}

// This function is not actively used but kept for potential debugging.
func DiscoverPeers(ctx context.Context, h host.Host, dht *dht.IpfsDHT, rendezvousPoint string) {
	log.Infof("Discovering peers for rendezvous point: %s", rendezvousPoint)

	routingDiscovery := routing_discovery.NewRoutingDiscovery(dht)
	peerChan, err := routingDiscovery.FindPeers(ctx, rendezvousPoint)
	if err != nil {
		log.Errorf("Failed to find peers: %v", err)
		return
	}

	for p := range peerChan {
		if p.ID == h.ID() {
			continue
		}
		log.Infof("Found peer: %s", p.ID.String())
		if err := h.Connect(ctx, p); err != nil {
			log.Warnf("Failed to connect to peer %s: %s", p.ID.String(), err)
		} else {
			log.Infof("Connected to peer: %s", p.ID.String())
		}
	}
}

