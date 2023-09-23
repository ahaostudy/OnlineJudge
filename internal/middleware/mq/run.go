package mq

import "main/config"

var (
	RMQJudge         *RabbitMQ
	RMQPrivate       *RabbitMQ
	RMQContestSubmit *RabbitMQ
)

// RunJudgeMQ 启动JudgeMQ
func RunJudgeMQ() *RabbitMQ {
	RMQJudge = NewWorkRabbitMQ("judge")
	for i := 0; i < config.ConfJudge.MaxJudgerCount; i++ {
		go RMQJudge.Consume(Judge)
	}
	return RMQJudge
}

// RunPrivateMQ 启动PrivateMQ
func RunPrivateMQ() *RabbitMQ {
	RMQPrivate = NewWorkRabbitMQ("private")
	go RMQPrivate.Consume(PrivateJudge)
	return RMQPrivate
}

// RunContestSubmitMQ 启动ContestSubmitMQ
func RunContestSubmitMQ() *RabbitMQ {
	RMQContestSubmit = NewWorkRabbitMQ("contest_submit")
	go RMQContestSubmit.Consume(ContestSubmit)
	return RMQContestSubmit
}
