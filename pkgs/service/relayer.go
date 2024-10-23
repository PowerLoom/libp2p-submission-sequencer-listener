package service

import (
	"Listen/config"
	"context"
	"encoding/base64"
	"os"
	"syscall"
	"time"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/network"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	rcmgr "github.com/libp2p/go-libp2p/p2p/host/resource-manager"
	"github.com/libp2p/go-libp2p/p2p/muxer/yamux"
	"github.com/libp2p/go-libp2p/p2p/net/connmgr"
	"github.com/libp2p/go-libp2p/p2p/security/noise"
	libp2ptls "github.com/libp2p/go-libp2p/p2p/security/tls"
	ma "github.com/multiformats/go-multiaddr"
	log "github.com/sirupsen/logrus"
)

var RelayerHost host.Host
var activeConnections int

func handleConnectionEstablished(network network.Network, conn network.Conn) {
	activeConnections++
}

func handleConnectionClosed(network network.Network, conn network.Conn) {
	activeConnections--
}

func generateAndSaveKey() (crypto.PrivKey, error) {
	priv, _, err := crypto.GenerateKeyPair(crypto.RSA, 2048)
	if err != nil {
		return nil, err
	}
	encodedKey := PrivateKeyToString(priv)
	err = os.WriteFile("/keys/key.txt", []byte(encodedKey), 0644)
	if err != nil {
		log.Fatal(err)
	}
	return priv, nil
}

func PrivateKeyToString(privKey crypto.PrivKey) string {
	privBytes, err := crypto.MarshalPrivateKey(privKey)
	if err != nil {
		log.Errorln("Unable to convert private key to string")
	}

	encodedKey := base64.StdEncoding.EncodeToString(privBytes)
	return encodedKey
}

func ConfigureRelayer() {
	var err error

	tcpAddr, err := ma.NewMultiaddr("/ip4/0.0.0.0/tcp/9100")
	if err != nil {
		log.Debugln(err)
	}

	connManager, _ := connmgr.NewConnManager(
		409600,
		819200,
		connmgr.WithGracePeriod(5*time.Minute))

	scalingLimits := rcmgr.DefaultLimits

	libp2p.SetDefaultServiceLimits(&scalingLimits)

	scaledDefaultLimits := scalingLimits.AutoScale()

	var rlim syscall.Rlimit
	if err = syscall.Getrlimit(syscall.RLIMIT_NOFILE, &rlim); err != nil {
		log.Debugln("failed to fetch RLimit")
	}

	err = syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rlim)
	if err != nil {
		log.Debugln("failed to set RLimit")
	}

	//TODO: convert these to config
	cfg := rcmgr.PartialLimitConfig{
		System: rcmgr.ResourceLimits{
			Streams:         rcmgr.Unlimited,
			StreamsOutbound: rcmgr.Unlimited,
			StreamsInbound:  rcmgr.Unlimited,
			Conns:           rcmgr.Unlimited,
			ConnsInbound:    rcmgr.Unlimited,
			ConnsOutbound:   rcmgr.Unlimited,
			FD:              rcmgr.Unlimited,
			Memory:          rcmgr.LimitVal64(rcmgr.Unlimited),
		},
	}

	limits := cfg.Build(scaledDefaultLimits)

	limiter := rcmgr.NewFixedLimiter(limits)

	rm, err := rcmgr.NewResourceManager(limiter, rcmgr.WithMetricsDisabled())
	if err != nil {
		log.Debugln("Unable to build resource manager")
	}

	var pk crypto.PrivKey
	if config.SettingsObj.RelayerPrivateKey == "" {
		pk, err = generateAndSaveKey()
		if err != nil {
			log.Errorln("Unable to generate private key")
		} else {
			log.Debugln("Libp2p host private key: ", PrivateKeyToString(pk))
		}
	} else {
		pkBytes, err := base64.StdEncoding.DecodeString(config.SettingsObj.RelayerPrivateKey)
		if err != nil {
			log.Errorln("Unable to decode private key")
		} else {
			pk, err = crypto.UnmarshalPrivateKey(pkBytes)
			if err != nil {
				log.Errorln("Unable to unmarshal private key")
			}
		}
	}

	RelayerHost, err = libp2p.New(
		libp2p.EnableRelay(),
		libp2p.Identity(pk),
		libp2p.ConnectionManager(connManager),
		libp2p.ListenAddrs(tcpAddr),
		libp2p.Security(libp2ptls.ID, libp2ptls.New),
		libp2p.Security(noise.ID, noise.New),
		libp2p.DefaultTransports,
		libp2p.NATPortMap(),
		libp2p.EnableNATService(),
		libp2p.ResourceManager(rm),
		libp2p.EnableHolePunching(),
		libp2p.Muxer(yamux.ID, yamux.DefaultTransport),
	)
	RelayerHost.Network().Notify(&network.NotifyBundle{
		ConnectedF:    handleConnectionEstablished,
		DisconnectedF: handleConnectionClosed,
	})

	kademliaDHT := ConfigureDHT(context.Background(), RelayerHost)

	// Create a discovery service using the DHT
	routingDiscovery := routing.NewRoutingDiscovery(kademliaDHT)

	log.Debugln("Listener ID:", RelayerHost.ID().String())
	log.Debugln("Peerable addresses: ", RelayerHost.Addrs())

	// Form initial connections to all peers available
	go ConnectToPeers(context.Background(), routingDiscovery, config.SettingsObj.RelayerRendezvousPoint, RelayerHost)

	// Keep checking for new peers at a set interval
	go DiscoverAtInterval(context.Background(), routingDiscovery, config.SettingsObj.RelayerRendezvousPoint, RelayerHost, time.NewTicker(5*time.Minute))

	log.Debugf("Listener host info: %s", RelayerHost.ID().String())
}
