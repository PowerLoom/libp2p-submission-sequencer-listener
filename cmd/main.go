package main

import (
	"Listen/config"
	"Listen/pkgs/prost"
	"Listen/pkgs/redis"
	"Listen/pkgs/service"
	"Listen/pkgs/utils"
	"Listen/pkgs/watcher"
	"context"
	log "github.com/sirupsen/logrus"
	"sync"
	"time"
)

func main() {
	utils.InitLogger()
	config.LoadConfig()

	service.InitializeReportingService(config.SettingsObj.SlackReportingUrl, 5*time.Second)

	var wg sync.WaitGroup

	prost.ConfigureClient()
	prost.ConfigureContractInstance()
	redis.RedisClient = redis.NewRedisClient()

	host, kademliaDHT, err := service.NewHost(context.Background(), config.SettingsObj.BootstrapPeers, config.SettingsObj.ListenerP2PPort)
	if err != nil {
		log.Fatalf("Failed to create host: %v", err)
	}

	gossipManager, err := service.NewGossipsubManager(context.Background(), config.SettingsObj, redis.RedisClient, host, kademliaDHT)
	if err != nil {
		log.Fatalf("Failed to create gossipsub manager: %v", err)
	}

	watcherService, err := watcher.NewWatcher(*config.SettingsObj)
	if err != nil {
		log.Fatalf("Failed to create watcher: %v", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		gossipManager.Start(context.Background())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		watcherService.Start(context.Background())
	}()

	wg.Wait()
}
