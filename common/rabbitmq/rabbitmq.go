package rabbitmq

import (
	"fmt"

	"github.com/streadway/amqp"
)

// 全局connection对象
// 所有RabbitMQ都会复用该对象
var conn *amqp.Connection

type Config struct {
	Username string
	Password string
	Host     string
	Port     int
	Vhost    string
}

// 初始化connection
func initConn(c *Config) {
	mqUrl := fmt.Sprintf(
		"amqp://%s:%s@%s:%d/%s",
		c.Username, c.Password, c.Host, c.Port, c.Vhost,
	)

	var err error
	conn, err = amqp.Dial(mqUrl)
	if err != nil {
		panic(err.Error())
	}
}

// RabbitMQ RabbitMQ结构体
type RabbitMQ struct {
	conn     *amqp.Connection
	channel  *amqp.Channel
	Exchange string
	Key      string
}

// NewRabbitMQ 创建RabbitMQ对象
func NewRabbitMQ(exchange string, key string) *RabbitMQ {
	return &RabbitMQ{Exchange: exchange, Key: key}
}

// Destroy 断开 channel 和 connection
func (r *RabbitMQ) Destroy() {
	_ = r.channel.Close()
	_ = r.conn.Close()
}

// NewWorkRabbitMQ 创建Work模式的RabbitMQ实例
func NewWorkRabbitMQ(c *Config, queue string) *RabbitMQ {
	// new rabbitmq
	rabbitmq := NewRabbitMQ("", queue)

	// get connection
	if conn == nil {
		initConn(c)
	}
	rabbitmq.conn = conn

	// get channel
	var err error
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	if err != nil {
		panic(err.Error())
	}

	return rabbitmq
}

// Publish 发送消息
func (r *RabbitMQ) Publish(message []byte) error {
	// 创建队列（不存在时）
	// 使用默认交换机的情况下，queue即为key
	_, err := r.channel.QueueDeclare(r.Key, false, false, false, false, nil)
	if err != nil {
		return err
	}

	// 调用 channel 发送消息到队列
	return r.channel.Publish(r.Exchange, r.Key, false, false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		},
	)
}

// Consume 消费者
// handle: 消息的消费业务函数，用于消费消息
func (r *RabbitMQ) Consume(handle func(msg *amqp.Delivery) error) {
	// 创建队列
	q, err := r.channel.QueueDeclare(r.Key, false, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	// 接收消息
	msgs, err := r.channel.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		panic(err)
	}

	// 处理消息
	for msg := range msgs {
		if err := handle(&msg); err != nil {
			fmt.Println(err.Error())
		}
	}
}
