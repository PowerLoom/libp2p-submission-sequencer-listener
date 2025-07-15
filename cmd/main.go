package main

import (
	"Listen/config"
	"Listen/pkgs/prost"
	"Listen/pkgs/redis"
	"Listen/pkgs/service"
	"Listen/pkgs/utils"
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

	gossipManager, err := service.NewGossipsubManager(context.Background(), config.SettingsObj, redis.RedisClient)
	if err != nil {
		log.Fatalf("Failed to create gossipsub manager: %v", err)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		gossipManager.Start(context.Background())
	}()

	wg.Wait()
}
