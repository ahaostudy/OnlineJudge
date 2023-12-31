package judge

import (
	"fmt"
	"log"
	"net"
	"os"

	r "github.com/cloudwego/kitex/pkg/registry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	nacosserver "github.com/kitex-contrib/config-nacos/server"
	"github.com/kitex-contrib/registry-nacos/registry"

	nacosclient "main/common/nacos_client"
	nacosconfig "main/common/nacos_config"
	"main/kitex_gen/judge/judgeservice"
	"main/services/judge/client"
	"main/services/judge/config"
	"main/services/judge/dal/cache"
	"main/services/judge/dal/mq"
)

const DataId = "judge"

func InitDir() {
	if err := os.MkdirAll(config.Config.File.CodePath, os.ModePerm); err != nil {
		panic(err)
	}
	if err := os.MkdirAll(config.Config.File.TempPath, os.ModePerm); err != nil {
		panic(err)
	}
}

func Run() {
	cli, err := nacosclient.NewNamingClient()
	if err != nil {
		panic(err)
	}

	nacosClient, err := nacosconfig.NewNacosConfig(DataId, &config.Config)
	if err != nil {
		panic(err)
	}

	// 初始化一些目录
	InitDir()

	// 连接rpc服务
	client.InitClient(cli)

	// 连接缓存
	if err := cache.InitRedis(); err != nil {
		panic(err)
	}

	// 连接并启动MQ
	defer mq.Run().Destroy()

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", config.Config.Host, config.Config.Port))
	if err != nil {
		panic(err)
	}

	svr := judgeservice.NewServer(
		new(JudgeServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Config.Name}),
		server.WithRegistry(registry.NewNacosRegistry(cli)),
		server.WithRegistryInfo(&r.Info{
			ServiceName: config.Config.Name,
			Addr:        addr,
		}),
		server.WithSuite(nacosserver.NewSuite(config.Config.Name, nacosClient)),
		server.WithServiceAddr(addr),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
