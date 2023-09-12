package main

import (
	"main/config"
	"main/server/dao"
	"main/server/middleware/mq"
	"main/server/middleware/redis"
	"main/server/route"
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
