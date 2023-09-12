package config

import "github.com/spf13/viper"

const configPath = "/projects/OnlineJudge/config/config.yaml"

var (
	ConfRabbitMQ Rabbitmq
	ConfMySQL    Mysql
	ConfRedis    Redis
	ConfEtcd     Etcd

	ConfServer Server
	ConfJudge  Judge
)

// InitConfig 初始化配置
func InitConfig() error {
	conf := new(Default)

	// 读取YAML
	vp := viper.New()
	vp.SetConfigType("yaml")
	vp.SetConfigFile(configPath)
	if err := vp.ReadInConfig(); err != nil {
		return err
	}
	if err := vp.Unmarshal(conf); err != nil {
		return err
	}

	// config
	{
		ConfRabbitMQ = conf.Rabbitmq
		ConfMySQL = conf.Mysql
		ConfRedis = conf.Redis
		ConfEtcd = conf.Etcd

		ConfServer = conf.Server
		ConfJudge = conf.Judge
	}

	// consts
	return initConsts()
}
