package main

import (
	"main/config"
	"main/internal/gateway/route"
	"main/rpc"
)

func init() {
	if err := config.InitConfig(); err != nil {
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
