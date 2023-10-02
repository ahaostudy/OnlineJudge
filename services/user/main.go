package user

import (
	"log"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	nacosserver "github.com/kitex-contrib/config-nacos/server"
	"github.com/kitex-contrib/registry-nacos/registry"

	"main/kitex_gen/user/userservice"
	"main/pkg/common"
	"main/services/user/config"
	"main/services/user/dal/cache"
	"main/services/user/dal/db"
)

const DataId = "user"

func Run() {
	cli, err := common.NewNamingClient()
	if err != nil {
		panic(err)
	}

	nacosClient, err := common.NewNacosConfig(DataId, &config.Config)
	if err != nil {
		panic(err)
	}

	// 连接数据库
	if err := db.InitMySQL(); err != nil {
		panic(err)
	}

	// 连接缓存
	if err := cache.InitRedis(); err != nil {
		panic(err)
	}

	svr := userservice.NewServer(
		new(UserServiceImpl),
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
