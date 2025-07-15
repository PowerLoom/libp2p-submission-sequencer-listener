package service

import (
	"context"
	"strings"
	"sync"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
)

// NewHost creates a new libp2p host and connects to bootstrap peers.
func NewHost(ctx context.Context, bootstrapPeers string) (host.Host, error) {
	h, err := libp2p.New()
	if err != nil {
		return nil, err
	}

	kademliaDHT, err := dht.New(ctx, h)
	if err != nil {
		return nil, err
	}

	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		return nil, err
	}

	if bootstrapPeers != "" {
		ConnectToBootstrapPeers(ctx, h, bootstrapPeers)
	}

	log.Infof("Libp2p host created with ID: %s", h.ID())
	return h, nil
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
				log.Warningf("Failed to connect to bootstrap peer %s: %v", peerString, err)
			} else {
				log.Infof("Successfully connected to bootstrap peer: %s", peerString)
			}
		}(peerString)
	}
	wg.Wait()
}
