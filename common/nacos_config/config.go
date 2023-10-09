package nacosconfig

import (
	"fmt"
	"log"
	"main/common/config"

	"github.com/kitex-contrib/config-nacos/nacos"
	"github.com/nacos-group/nacos-sdk-go/vo"
)

func NewNacosConfig(DataId string, cfg any) (nacos.Client, error) {
	// 配置中心
	nacosClient, err := nacos.New(nacos.Options{
		Address:     config.NacosHost,
		Port:        uint64(config.NacosPort),
		NamespaceID: config.NamespaceID,
	})
	if err != nil {
		return nil, err
	}

	nacosClient.RegisterConfigCallback(
		vo.ConfigParam{
			DataId: DataId,
			Group:  config.Group,
		},
		func(data string, c nacos.ConfigParser) {
			fmt.Printf("data: %v\n", data)
			err := c.Decode(vo.YAML, data, cfg)
			if err != nil {
				log.Printf("nacos config error: %v\n", err)
			}
		},
	)

	return nacosClient, nil
}