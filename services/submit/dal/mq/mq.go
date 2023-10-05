package mq

import (
	"main/common/rabbitmq"
	"main/services/submit/config"
)

var (
	RMQContestSubmit *rabbitmq.RabbitMQ
)

func Run() *rabbitmq.RabbitMQ {
	RMQContestSubmit = rabbitmq.NewWorkRabbitMQ(&rabbitmq.Config{
		Username: config.Config.Rabbitmq.Username,
		Password: config.Config.Rabbitmq.Password,
		Host:     config.Config.Rabbitmq.Host,
		Port:     config.Config.Rabbitmq.Port,
		Vhost:    config.Config.Rabbitmq.Vhost,
	}, "contest_submit")
	go RMQContestSubmit.Consume(ContestSubmit)
	return RMQContestSubmit
}
