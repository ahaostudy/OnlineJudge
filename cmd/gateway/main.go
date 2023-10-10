package main

import (
	nacosclient "main/common/nacos_client"
	nacosconfig "main/common/nacos_config"
	"main/gateway/client"
	"main/gateway/route"
	"main/gateway/config"
)

func main() {
	cli, err := nacosclient.NewNamingClient()
	if err != nil {
		panic(err)
	}

	nacosconfig.NewNacosConfig("gateway", &config.Config)

	client.InitClient(cli)

	if err := route.InitRoute().Run(); err != nil {
		panic(err)
	}
}
