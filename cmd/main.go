package main

import (
	"Listen/config"
	"Listen/pkgs/prost"
	"Listen/pkgs/redis"
	"Listen/pkgs/service"
	"Listen/pkgs/utils"
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

func main() {
	utils.InitLogger()
	config.LoadConfig()

	service.InitializeReportingService(config.SettingsObj.SlackReportingUrl, 5*time.Second)

	var wg sync.WaitGroup

	prost.ConfigureClient()
	prost.ConfigureContractInstance()
	redis.RedisClient = redis.NewRedisClient()

	service.ConfigureRelayer()

	// Setup signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGTERM, syscall.SIGINT)

	// Start collector in background
	wg.Add(1)
	go func() {
		defer wg.Done()
		service.StartCollectorServer()
	}()

	// Wait for shutdown signal
	go func() {
		sig := <-sigChan
		log.Infof("Received signal %v, initiating graceful shutdown", sig)

		// Step 1: Stop accepting new connections
		log.Info("Closing libp2p host to stop accepting new connections...")
		if service.RelayerHost != nil {
			// This will close all streams and connections
			if err := service.RelayerHost.Close(); err != nil {
				log.Errorf("Error closing host: %v", err)
			} else {
				log.Info("Successfully closed all libp2p connections")
			}
		}

		// Step 2: Give a small grace period for any in-flight messages
		log.Info("Waiting for in-flight messages to complete...")
		time.Sleep(5 * time.Second)

		// Step 3: Close Redis connection
		log.Info("Closing Redis connection...")
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if redis.RedisClient != nil {
			if err := redis.RedisClient.Shutdown(ctx).Err(); err != nil {
				log.Errorf("Error closing Redis: %v", err)
			}
		}

		log.Info("Graceful shutdown complete")
		os.Exit(0)
	}()

	wg.Wait()
}
