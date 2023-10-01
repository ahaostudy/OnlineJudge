package test

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"

	"main/kitex_gen/problem/problemservice"
	"main/pkg/common"
)

var (
	ProblemCli problemservice.Client
	Ctx        = context.Background()
)

func InitClient() {
	cli, err := common.NewNamingClient()
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
