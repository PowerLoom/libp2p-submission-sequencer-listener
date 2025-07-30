package service

import (
	"Listen/config"
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/multiformats/go-multiaddr"
	"github.com/libp2p/go-libp2p/p2p/net/connmgr"
	rcmgr "github.com/libp2p/go-libp2p/p2p/host/resource-manager"
	log "github.com/sirupsen/logrus"
)

// NewHost creates a new libp2p host and connects to bootstrap peers.
func NewHost(ctx context.Context, bootstrapPeers string, listenerPort string) (h host.Host, kademliaDHT *dht.IpfsDHT, err error) {
	listenAddr := fmt.Sprintf("/ip4/0.0.0.0/tcp/%s", listenerPort)
	
	// Create a new resource manager with scaled limits.
	limiter := rcmgr.NewFixedLimiter(rcmgr.DefaultLimits.AutoScale())
	rscMgr, err := rcmgr.NewResourceManager(limiter)
	if err != nil {
		return
	}

	cm, err := connmgr.NewConnManager(
		100, // Lowwater
		400, // Highwater
		connmgr.WithGracePeriod(time.Minute),
	)
	if err != nil {
		return
	}

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(listenAddr),
		libp2p.ConnectionManager(cm),
		libp2p.ResourceManager(rscMgr),
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
		return
	}

	// Parse bootstrap peers
	bootstrapAddrInfos, err := parseBootstrapPeers(bootstrapPeers)
	if err != nil {
		// Log the error but continue, the node might discover peers through other means
		log.Warnf("Failed to parse bootstrap peers: %v", err)
	}

	// Create a new Kademlia DHT in client mode, providing the bootstrap peers.
	// The DHT will automatically use these peers to bootstrap itself.
	kademliaDHT, err = dht.New(ctx, h, dht.Mode(dht.ModeClient), dht.BootstrapPeers(bootstrapAddrInfos...))
	if err != nil {
		return
	}

	// It's good practice to trigger a bootstrap process in the background.
	// This ensures the DHT actively seeks to connect and populate its routing table.
	go func() {
		if err := kademliaDHT.Bootstrap(ctx); err != nil {
			log.Warnf("Initial DHT bootstrap failed: %v", err)
		}
	}()

	// Announce our presence using the rendezvous point
	go func() {
		log.Info("Starting rendezvous announcement loop...")
		routingDiscovery := routing.NewRoutingDiscovery(kademliaDHT)
		// Give the DHT a moment to connect to bootstrap peers before starting to advertise.
		time.Sleep(5 * time.Second)
		ticker := time.NewTicker(15 * time.Second) // Advertise every 15 seconds
		defer ticker.Stop()

		for {
			log.Infof("Advertising our presence for rendezvous point: %s", config.SettingsObj.RendezvousPoint)

			ttl, err := routingDiscovery.Advertise(ctx, config.SettingsObj.RendezvousPoint)
			if err != nil {
				log.Errorf("Failed to advertise rendezvous point: %v", err)
			} else {
				log.Infof("Successfully advertised! Time to live for advertisement: %s", ttl)
			}

			select {
			case <-ticker.C:
				// Continue to next iteration
			case <-ctx.Done():
				log.Info("Stopping rendezvous announcement loop.")
				return
			}
		}
	}()

	log.Infof("Libp2p host created with ID: %s", h.ID())
	return
}

// parseBootstrapPeers converts a comma-separated string of multiaddresses into a slice of AddrInfo.
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

// ConnectToBootstrapPeers connects the host to a list of bootstrap peers.
// This function is kept for potential future debugging but is not actively used in the NewHost flow.
func ConnectToBootstrapPeers(ctx context.Context, h host.Host, addrInfos []peer.AddrInfo) {
	var wg sync.WaitGroup
	for _, pi := range addrInfos {
		wg.Add(1)
		go func(peerInfo peer.AddrInfo) {
			defer wg.Done()
			if err := h.Connect(ctx, peerInfo); err != nil {
				log.Errorf("Failed to connect to bootstrap peer %s: %v", peerInfo.ID, err)
			} else {
				log.Infof("Successfully connected to bootstrap peer: %s", peerInfo.ID)
			}
		}(pi)
	}
	wg.Wait()
}


// DiscoverPeers finds peers for a given rendezvous point.
func DiscoverPeers(ctx context.Context, h host.Host, dht *dht.IpfsDHT, rendezvousPoint string) {
	log.Infof("Discovering peers for rendezvous point: %s", rendezvousPoint)

	routingDiscovery := routing.NewRoutingDiscovery(dht)
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
