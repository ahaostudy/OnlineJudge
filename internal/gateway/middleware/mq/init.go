package mq

var (
	RMQJudge *RabbitMQ
)

// InitRabbitMQ 初始化RabbitMQ
func InitRabbitMQ() {
	// 创建MQ并启动消费者
	// 所有 RabbitMQ 复用同一个连接

	RMQJudge = NewWorkRabbitMQ("judge")
	go RMQJudge.Consume(Judge)
}

// DestroyRabbitMQ 销毁RabbitMQ
func DestroyRabbitMQ() {
	RMQJudge.Destroy()
}
