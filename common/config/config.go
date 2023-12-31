package config

import (
	"flag"
	"fmt"

	"github.com/spf13/viper"
)

var configPath string

var (
	NacosHost   string
	NacosPort   int
	NamespaceID string
	Group       string
)

func init() {
	flag.StringVar(&configPath, "cp", "config/config.yaml", "config path")
	flag.Parse()

	viper.SetConfigType("yaml")
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	initConfigVar()
}

func initConfigVar() {
	NacosHost = viper.GetString("nacosHost")
	NacosPort = viper.GetInt("nacosPort")
	NamespaceID = viper.GetString("namespaceID")
	Group = viper.GetString("group")

	fmt.Printf("NacosHost: %v\n", NacosHost)
	fmt.Printf("NacosPort: %v\n", NacosPort)
	fmt.Printf("NamespaceID: %v\n", NamespaceID)
	fmt.Printf("Group: %v\n", Group)
}
