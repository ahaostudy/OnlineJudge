package mq

var (
	RMQJudge   *RabbitMQ
	RMQPrivate *RabbitMQ
)

// RunJudgeMQ 启动JudgeMQ
func RunJudgeMQ() *RabbitMQ {
	RMQJudge = NewWorkRabbitMQ("judge")
	go RMQJudge.Consume(Judge)
	go RMQJudge.Consume(Judge)
	return RMQJudge
}

// RunPrivateMQ 启动PrivateMQ
func RunPrivateMQ() *RabbitMQ {
	RMQPrivate = NewWorkRabbitMQ("private")
	go RMQPrivate.Consume(Judge)
	return RMQPrivate
}
