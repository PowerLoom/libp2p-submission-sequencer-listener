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
		log.Errorf("Failed to parse bootstrap peers: %v", err)
		// Continue without bootstrap peers if parsing fails
	}

	// Create a new Kademlia DHT
	kademliaDHT, err = dht.New(ctx, h, dht.Mode(dht.ModeClient), dht.BootstrapPeers(bootstrapAddrInfos...))
	if err != nil {
		return
	}

	// Connect to the bootstrap peers
	if len(bootstrapAddrInfos) > 0 {
		ConnectToBootstrapPeers(ctx, h, bootstrapAddrInfos)
		if err = kademliaDHT.Bootstrap(ctx); err != nil {
			log.Warnf("DHT bootstrap failed: %v", err)
			// This is not a fatal error, the node can still discover peers over time
		}
	}

	// Add a delay to allow the DHT to populate before we start advertising.
	// This is a temporary diagnostic step.
	log.Info("Waiting for 5 seconds for DHT to populate...")
	time.Sleep(5 * time.Second)


	// Announce our presence using the rendezvous point
	go func() {
		log.Info("Starting rendezvous announcement loop...")
		routingDiscovery := routing.NewRoutingDiscovery(kademliaDHT)
		ticker := time.NewTicker(10 * time.Second) // Advertise more frequently for debugging
		defer ticker.Stop()

		for {
			log.Infof("Advertising our presence for rendezvous point: %s", config.SettingsObj.RendezvousPoint)

			// Retry advertisement until successful or context is done
			for i := 0; i < config.SettingsObj.AdvertiseRetries; i++ { // Try up to configurable times
				ttl, err := routingDiscovery.Advertise(ctx, config.SettingsObj.RendezvousPoint)
				if err == nil {
					log.Infof("Successfully advertised! Time to live for advertisement: %s", ttl)
					break // Exit retry loop on success
				} else {
					log.Errorf("Failed to advertise rendezvous point (attempt %d/%d): %v", i+1, config.SettingsObj.AdvertiseRetries, err)
					if i < config.SettingsObj.AdvertiseRetries-1 { // Don't sleep after last attempt
						time.Sleep(time.Duration(config.SettingsObj.AdvertiseRetryDelaySec) * time.Second) // Wait before retrying
					}
				}
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
