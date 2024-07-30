package main

import (
	"Listen/config"
	"Listen/pkgs/prost"
	"Listen/pkgs/redis"
	"Listen/pkgs/service"
	"Listen/pkgs/utils"
	"sync"
)

func main() {
	utils.InitLogger()
	config.LoadConfig()

	var wg sync.WaitGroup

	prost.ConfigureClient()
	prost.ConfigureContractInstance()
	redis.RedisClient = redis.NewRedisClient()

	service.ConfigureRelayer()
	service.StartCollectorServer()

	wg.Add(1)
	go service.StartApiServer()
	wg.Wait()
}
