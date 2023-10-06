package client

import (
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"

	"main/kitex_gen/contest/contestservice"
	"main/kitex_gen/judge/judgeservice"
	"main/kitex_gen/problem/problemservice"
)

var (
	JudgeCli   judgeservice.Client
	ContestCli contestservice.Client
	ProblemCli problemservice.Client
)

func InitClient(cli naming_client.INamingClient) {
	JudgeCli = judgeservice.MustNewClient(
		"judge",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)
	ContestCli = contestservice.MustNewClient(
		"contest",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)
	ProblemCli = problemservice.MustNewClient(
		"problem",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)
}
