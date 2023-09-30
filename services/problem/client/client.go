package client

import (
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"

	"main/kitex_gen/contest/contestservice"
	"main/kitex_gen/submit/submitservice"
)

var (
	ContestCli contestservice.Client
	SubmitCli submitservice.Client
)

func InitClient(cli naming_client.INamingClient) {
	ContestCli = contestservice.MustNewClient(
		"contest",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)

	SubmitCli = submitservice.MustNewClient(
		"submit",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)
}
