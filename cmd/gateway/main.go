package main

import (
	"github.com/opentracing/opentracing-go"

	"main/config"
	"main/internal/middleware/tracing"
	"main/internal/gateway/route"
	"main/rpc"
)

func init() {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}
}

func main() {
	// 链路追踪
	tracer, closer := tracing.InitTracer("gateway")
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	if err := rpc.InitGRPCClients(); err != nil {
		panic(err)
	}
	defer rpc.CloseGPRCClients()

	if err := route.InitRoute().Run(); err != nil {
		panic(err)
	}
}
