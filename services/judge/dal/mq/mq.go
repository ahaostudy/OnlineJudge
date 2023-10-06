package mq

import (
	"main/common/rabbitmq"
	"main/services/judge/config"
)

var (
	RMQJudge *rabbitmq.RabbitMQ
)

func Run() *rabbitmq.RabbitMQ {
	RMQJudge = rabbitmq.NewWorkRabbitMQ(&rabbitmq.Config{
		Username: config.Config.Rabbitmq.Username,
		Password: config.Config.Rabbitmq.Password,
		Host:     config.Config.Rabbitmq.Host,
		Port:     config.Config.Rabbitmq.Port,
		Vhost:    config.Config.Rabbitmq.Vhost,
	}, "judge")
	go RMQJudge.Consume(Judge)
	return RMQJudge
}
