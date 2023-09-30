package test

import (
	"context"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/registry-nacos/resolver"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"main/kitex_gen/problem/problemservice"
)

var (
	ProblemCli problemservice.Client
	Ctx        = context.Background()
)

const (
	NacosHost = "127.0.0.1"
	NacosPort = 8848

	NamespaceID = "8f995970-806d-4063-a5e7-f3f9b6a4d141"
	DataId      = "problem"
)

func InitClient() {
	sc := []constant.ServerConfig{
		*constant.NewServerConfig(NacosHost, NacosPort),
	}

	// 注册中心
	cc := constant.ClientConfig{
		NamespaceId:         "public",
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogDir:              "nacos/log",
		CacheDir:            "nacos/cache",
		LogLevel:            "info",
	}
	cli, err := clients.NewNamingClient(
		vo.NacosClientParam{
			ClientConfig:  &cc,
			ServerConfigs: sc,
		},
	)
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
