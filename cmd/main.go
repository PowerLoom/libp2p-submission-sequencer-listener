package main

import (
	"Listen/config"
	"Listen/pkgs/prost"
	"Listen/pkgs/redis"
	"Listen/pkgs/service"
	"Listen/pkgs/utils"
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

	service.ConfigureRelayer()

	wg.Add(1)
	go service.StartCollectorServer()
	wg.Wait()
}
