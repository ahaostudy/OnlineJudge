package client

import (
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"

	"main/kitex_gen/problem/problemservice"
)

var (
	ProblemCli problemservice.Client
)

func InitClient(cli naming_client.INamingClient) {
	ProblemCli = problemservice.MustNewClient(
		"submit",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)
}
