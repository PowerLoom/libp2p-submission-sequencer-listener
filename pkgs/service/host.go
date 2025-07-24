package service

import (
	"Listen/config"
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	dutil "github.com/libp2p/go-libp2p/p2p/discovery/util"
	"github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
)

// NewHost creates a new libp2p host and connects to bootstrap peers.
func NewHost(ctx context.Context, bootstrapPeers string, listenerPort string) (h host.Host, kademliaDHT *dht.IpfsDHT, err error) {
	listenAddr := fmt.Sprintf("/ip4/0.0.0.0/tcp/%s", listenerPort)

	opts := []libp2p.Option{
		libp2p.ListenAddrStrings(listenAddr),
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

	kademliaDHT, err = dht.New(ctx, h)
	if err != nil {
		return
	}

	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		return
	}

	if bootstrapPeers != "" {
		ConnectToBootstrapPeers(ctx, h, bootstrapPeers)
	}

	// Announce our presence using the rendezvous point
	log.Info("Announcing ourselves...")
	routingDiscovery := routing.NewRoutingDiscovery(kademliaDHT)
	dutil.Advertise(ctx, routingDiscovery, config.SettingsObj.RendezvousPoint)
	log.Info("Successfully announced!")

	log.Infof("Libp2p host created with ID: %s", h.ID())
	return
}

// ConnectToBootstrapPeers connects the host to a list of bootstrap peers.
func ConnectToBootstrapPeers(ctx context.Context, h host.Host, peers string) {
	peerStrings := strings.Split(peers, ",")
	var wg sync.WaitGroup
	for _, peerString := range peerStrings {
		if peerString == "" {
			continue
		}
		wg.Add(1)
		go func(peerString string) {
			defer wg.Done()
			addr, err := multiaddr.NewMultiaddr(peerString)
			if err != nil {
				log.Errorf("Failed to parse multiaddr: %v", err)
				return
			}
			peerInfo, err := peer.AddrInfoFromP2pAddr(addr)
			if err != nil {
				log.Errorf("Failed to get peer info from multiaddr: %v", err)
				return
			}
			if err := h.Connect(ctx, *peerInfo); err != nil {
				log.Errorf("Failed to connect to bootstrap peer %s: %v", peerString, err)
			} else {
				log.Infof("Successfully connected to bootstrap peer: %s", peerString)
			}
		}(peerString)
	}
	wg.Wait()
}
