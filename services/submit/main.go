package submit

import (
	"fmt"
	"log"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	nacosserver "github.com/kitex-contrib/config-nacos/server"
	"github.com/kitex-contrib/registry-nacos/registry"

	nacosclient "main/common/nacos_client"
	nacosconfig "main/common/nacos_config"
	"main/kitex_gen/submit/submitservice"
	"main/services/submit/client"
	"main/services/submit/config"
	"main/services/submit/dal/cache"
	"main/services/submit/dal/db"
	"main/services/submit/dal/mq"
	"main/services/submit/jobs"
)

const DataId = "submit"

func Run() {
	cli, err := nacosclient.NewNamingClient()
	if err != nil {
		panic(err)
	}

	nacosClient, err := nacosconfig.NewNacosConfig(DataId, &config.Config)
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

	// 连接并启动MQ
	defer mq.Run().Destroy()

	// 运行启动任务
	jobs.RunSubmitJobs()

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf(":%d", config.Config.Port))
	if err != nil {
		panic(err)
	}

	svr := submitservice.NewServer(
		new(SubmitServiceImpl),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Config.Name}),
		server.WithRegistry(registry.NewNacosRegistry(cli)),
		server.WithSuite(nacosserver.NewSuite(config.Config.Name, nacosClient)),
		server.WithServiceAddr(addr),
	)
	if err := svr.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
