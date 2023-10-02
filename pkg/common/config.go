package common

import (
	"log"

	"github.com/kitex-contrib/config-nacos/nacos"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

const (
	NamespaceID = "8f995970-806d-4063-a5e7-f3f9b6a4d141"
	Group = "DEFAULT_GROUP"
)

func NewNacosConfig(DataId string, config any) (nacos.Client, error) {
	// 配置中心
	nacosClient, err := nacos.New(nacos.Options{
		NamespaceID: NamespaceID,
	})
	if err != nil {
		return nil, err
	}

	nacosClient.RegisterConfigCallback(
		vo.ConfigParam{
			DataId: DataId,
			Group:  Group,
		},
		func(data string, c nacos.ConfigParser) {
			err := c.Decode(vo.YAML, data, config)
			if err != nil {
				log.Printf("nacos config error: %v\n", err)
			}
		},
	)

	return nacosClient, nil
}