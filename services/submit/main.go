package submit

import (
	"log"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	nacosserver "github.com/kitex-contrib/config-nacos/server"
	"github.com/kitex-contrib/registry-nacos/registry"

	"main/kitex_gen/submit/submitservice"
	"main/pkg/common"
	"main/services/problem/client"
	"main/services/submit/config"
	"main/services/submit/dal/cache"
	"main/services/submit/dal/db"
)

const DataId = "submit"

func Run() {
	cli, err := common.NewNamingClient()
	if err != nil {
		panic(err)
	}

	nacosClient, err := common.NewNacosConfig(DataId, &config.Config)
	if err != nil {
		panic(err)
	}

	// 连接rpc服务
	client.InitClient(cli)

	// 连接数据库
	if err := db.InitMySQL(); err != nil {
		panic(err)
	}

	// 连接缓存
	if err := cache.InitRedis(); err != nil {
		panic(err)
	}

	svr := submitservice.NewServer(
		new(SubmitServiceImpl),
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
