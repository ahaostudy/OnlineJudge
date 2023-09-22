package mq

var (
	RMQJudge         *RabbitMQ
	RMQPrivate       *RabbitMQ
	RMQContestSubmit *RabbitMQ
	RMQSubmit        *RabbitMQ
)

// RunJudgeMQ 启动JudgeMQ
func RunJudgeMQ() *RabbitMQ {
	RMQJudge = NewWorkRabbitMQ("judge")
	go RMQJudge.Consume(Judge)
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

func RunSubmitMQ() *RabbitMQ {
	RMQSubmit = NewWorkRabbitMQ("contest_submit")
	go RMQSubmit.Consume(Submit)
	return RMQSubmit
}
