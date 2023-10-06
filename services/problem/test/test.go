package test

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"

	nacosclient "main/common/nacos_client"
	"main/kitex_gen/problem/problemservice"
)

var (
	ProblemCli problemservice.Client
	Ctx        = context.Background()
)

func InitClient() {
	cli, err := nacosclient.NewNamingClient()
	if err != nil {
		panic(err)
	}

	ProblemCli = problemservice.MustNewClient(
		"problem",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)
}

func init() {
	InitClient()
}
