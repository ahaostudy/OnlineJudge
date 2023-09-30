package problem

import (
	"fmt"
	"log"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/config-nacos/nacos"
	nacosserver "github.com/kitex-contrib/config-nacos/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"

	"main/kitex_gen/problem/problemservice"
	"main/services/problem/client"
	"main/services/problem/config"
	"main/services/problem/dal/db"
)

const (
	NacosHost = "127.0.0.1"
	NacosPort = 8848

	NamespaceID = "8f995970-806d-4063-a5e7-f3f9b6a4d141"
	DataId      = "problem"
)

func Run() {
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

	// 配置中心
	nacosClient, err := nacos.New(nacos.Options{
		NamespaceID: NamespaceID,
	})
	if err != nil {
		panic(err)
	}
	nacosClient.RegisterConfigCallback(
		vo.ConfigParam{
			DataId: DataId,
			Group:  "DEFAULT_GROUP",
		},
		func(data string, c nacos.ConfigParser) {
			err := c.Decode(vo.YAML, data, &config.Config)
			if err != nil {
				fmt.Printf("nacos config error: %v\n", err)
			}
		},
	)

	// 初始化客户端连接
	client.InitClient(cli)

	// 连接数据库
	if err := db.InitMySQL(); err != nil {
		panic(err)
	}

	svr := problemservice.NewServer(
		new(ProblemServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Config.Name}),
		server.WithRegistry(registry.NewNacosRegistry(cli)),
		server.WithSuite(nacosserver.NewSuite(config.Config.Name, nacosClient)),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
