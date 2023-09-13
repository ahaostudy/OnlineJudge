package main

import (
	"main/config"
	"main/internal/gateway/dao"
	"main/internal/gateway/route"
	"main/internal/middleware/redis"
	"main/rpc"
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
	if err := rpc.InitGRPCClients(); err != nil {
		panic(err)
	}
	defer rpc.CloseGPRCClients()

	if err := route.InitRoute().Run(); err != nil {
		panic(err)
	}
}
