package main

import (
	nacosclient "main/common/nacos_client"
	"main/gateway/client"
	"main/gateway/route"
)

func main() {
	cli, err := nacosclient.NewNamingClient()
	if err != nil {
		panic(err)
	}

	client.InitClient(cli)

	if err := route.InitRoute().Run(); err != nil {
		panic(err)
	}
}
