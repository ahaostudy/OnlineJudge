package main

import (
	"main/config"
	"main/internal/gateway/dao"
	"main/internal/gateway/middleware/mq"
	"main/internal/gateway/middleware/redis"
	"main/internal/gateway/route"
)

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	if err := dao.InitMySQL(); err != nil {
		panic(err)
	}

	if err := redis.InitRedis(); err != nil {
		panic(err)
	}
}

func main() {
	mq.InitRabbitMQ()
	defer mq.DestroyRabbitMQ()

	if err := route.InitRoute().Run(); err != nil {
		panic(err)
	}
}
