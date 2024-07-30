package main

import (
	"Listen/config"
	"Listen/pkgs/redis"
	"Listen/pkgs/service"
	"Listen/pkgs/utils"
	"sync"
)

func main() {
	utils.InitLogger()
	config.LoadConfig()

	var wg sync.WaitGroup

	redis.RedisClient = redis.NewRedisClient()

	service.ConfigureRelayer()
	service.StartCollectorServer()

	wg.Add(1)
	go service.StartApiServer()
	wg.Wait()
}
