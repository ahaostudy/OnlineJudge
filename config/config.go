package config

import "github.com/spf13/viper"

const configPath = "config/config.yaml"

var (
	ConfRabbitMQ Rabbitmq
	ConfMySQL    Mysql
	ConfRedis    Redis
	ConfMongodb  Mongodb
	ConfEtcd     Etcd
	ConfJaeger   Jaeger
	ConfAuth     Auth

	ConfPrivate Private
	ConfJudge   Judge
	ConfProblem Problem
	ConfSubmit  Submit
	ConfChatGPT Chatgpt
	ConfContest Contest
	ConfUser    User
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
		ConfMongodb = conf.Mongodb
		ConfEtcd = conf.Etcd
		ConfJaeger = conf.Jaeger
		ConfAuth = conf.Auth

		ConfPrivate = conf.Private
		ConfJudge = conf.Judge
		ConfProblem = conf.Problem
		ConfSubmit = conf.Submit
		ConfChatGPT = conf.Chatgpt
		ConfContest = conf.Contest
		ConfUser = conf.User
	}

	// consts
	return initConsts()
}
