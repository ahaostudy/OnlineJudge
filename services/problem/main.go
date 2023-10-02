package problem

import (
	"log"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	nacosserver "github.com/kitex-contrib/config-nacos/server"
	"github.com/kitex-contrib/registry-nacos/registry"

	"main/kitex_gen/problem/problemservice"
	"main/pkg/common"
	"main/services/problem/client"
	"main/services/problem/config"
	"main/services/problem/dal/db"
)

const DataId = "problem"

func Run() {
	cli, err := common.NewNamingClient()
	if err != nil {
		panic(err)
	}

	nacosClient, err := common.NewNacosConfig(DataId, &config.Config)
	if err != nil {
		panic(err)
	}

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
