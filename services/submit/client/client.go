package client

import (
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"

	"main/kitex_gen/judge/judgeservice"
)

var (
	JudgeCli judgeservice.Client
)

func InitClient(cli naming_client.INamingClient) {
	JudgeCli = judgeservice.MustNewClient(
		"judge",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)
}
