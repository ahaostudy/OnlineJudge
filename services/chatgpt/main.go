package chatgpt

import (
	"fmt"
	"log"
	nacosclient "main/common/nacos_client"
	nacosconfig "main/common/nacos_config"
	"main/kitex_gen/chatgpt/chatgptservice"
	"main/services/chatgpt/config"
	"net"

	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	nacosserver "github.com/kitex-contrib/config-nacos/server"
	"github.com/kitex-contrib/registry-nacos/registry"
	r "github.com/cloudwego/kitex/pkg/registry"
)

const DataId = "chatgpt"

func Run() {
	cli, err := nacosclient.NewNamingClient()
	if err != nil {
		panic(err)
	}

	nacosClient, err := nacosconfig.NewNacosConfig(DataId, &config.Config)
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%d", config.Config.Host, config.Config.Port))
	if err != nil {
		panic(err)
	}

	svr := chatgptservice.NewServer(
		new(ChatGPTServiceImpl),
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
