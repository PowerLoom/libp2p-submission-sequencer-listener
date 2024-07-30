package service

import (
	"context"
	"github.com/cenkalti/backoff/v4"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-libp2p/p2p/protocol/circuitv2/client"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func DiscoverAtInterval(ctx context.Context, routingDiscovery *routing.RoutingDiscovery, rendezvousString string, host host.Host, ticker *time.Ticker) {
	//ticker := time.NewTicker(10 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			// Discover peers every 10 minutes
			ConnectToPeers(ctx, routingDiscovery, rendezvousString, host)
		case <-ctx.Done():
			// Exit the loop if the context is canceled
			return
		}
	}
}

func ConnectToPeers(ctx context.Context, routingDiscovery *routing.RoutingDiscovery, rendezvousString string, host host.Host) {
	peerChan, err := routingDiscovery.FindPeers(ctx, rendezvousString)
	if err != nil {
		log.Errorf("Failed to find peers: %s", err)
		return
	}
	log.Debugln("Triggering periodic relayer peer discovery")

	for peer := range peerChan {
		if peer.ID == host.ID() || len(peer.Addrs) == 0 {
			continue // Skip self or peers with no addresses
		}

		operation := func() error {
			if host.Network().Connectedness(peer.ID) != network.Connected {
				if err := host.Connect(ctx, peer); err != nil {
					log.Errorln("Peer connection error: ", err.Error())
					return err
				}
			}

			// Attempt to reserve a slot once connected
			if _, err := client.Reserve(ctx, host, peer); err != nil {
				log.Errorln("Reservation error: ", err.Error())
				return err
			}
			return nil
		}

		err = backoff.Retry(operation, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 3))
		if err != nil {
			log.Errorf("Failed to connect and reserve with peer %s after retries: %s", peer.ID, err)
		} else {
			log.Infof("Successfully connected and reserved with peer %s", peer.ID)
		}
	}
	log.Debugln("Active connections : ", activeConnections)
}

func ConfigureDHT(ctx context.Context, host host.Host) *dht.IpfsDHT {
	// Set up a Kademlia DHT for the service host
	kademliaDHT, err := dht.New(ctx, host)
	if err != nil {
		log.Fatalf("Failed to create DHT: %s", err)
	}

	// Bootstrap the DHT
	if err = kademliaDHT.Bootstrap(ctx); err != nil {
		log.Fatalf("Failed to bootstrap DHT: %s", err)
	}

	var wg sync.WaitGroup
	for _, peerAddr := range dht.DefaultBootstrapPeers {
		peerinfo, _ := peer.AddrInfoFromP2pAddr(peerAddr)
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := host.Connect(ctx, *peerinfo); err != nil {
				log.Warning(err)
			} else {
				log.Debugln("Connection established with bootstrap node:", *peerinfo)
			}
		}()
	}
	wg.Wait()

	return kademliaDHT
}
