package client

import (
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients/naming_client"

	"main/kitex_gen/chatgpt/chatgptservice"
	"main/kitex_gen/contest/contestservice"
	"main/kitex_gen/judge/judgeservice"
	"main/kitex_gen/problem/problemservice"
	"main/kitex_gen/submit/submitservice"
	"main/kitex_gen/user/userservice"
)

var (
	ProblemCli problemservice.Client
	JudgeCli   judgeservice.Client
	SubmitCli  submitservice.Client
	UserCli    userservice.Client
	ContestCli contestservice.Client
	ChatGPTCli chatgptservice.Client
)

func InitClient(cli naming_client.INamingClient) {
	ProblemCli = problemservice.MustNewClient(
		"problem",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)
	JudgeCli = judgeservice.MustNewClient(
		"judge",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)
	SubmitCli = submitservice.MustNewClient(
		"submit",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)
	UserCli = userservice.MustNewClient(
		"user",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)
	ContestCli = contestservice.MustNewClient(
		"contest",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)
	ChatGPTCli = chatgptservice.MustNewClient(
		"chatgpt",
		client.WithResolver(resolver.NewNacosResolver(cli)),
		client.WithRPCTimeout(time.Second*3),
	)
}
