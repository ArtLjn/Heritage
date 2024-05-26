/*
@Time : 2024/5/25 上午11:39
@Author : ljn
@File : mq
@Software: GoLand
*/

package data

import (
	"github.com/streadway/amqp"
	"log"
)

type MQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func NewMQ(con *amqp.Connection, ch *amqp.Channel) *MQ {
	return &MQ{
		Conn:    con,
		Channel: ch,
	}
}

func ConnMq(url string) *MQ { // 修改返回值类型以包含错误处理
	conn, err := amqp.Dial(url)
	if err != nil {
		panic(err)
	}

	channel, err := conn.Channel()
	if err != nil {
		err = conn.Close()
		if err != nil {
			panic(err)
		}
		panic(err)
	}

	return NewMQ(conn, channel)
}

// CreateExchange 创建交换机
func (m *MQ) CreateExchange(exchangeName string, exchangeType string) error {
	err := m.Channel.ExchangeDeclare(
		exchangeName, // name
		exchangeType, // type
		true,         // durable
		false,        // auto-deleted
		false,        // internal
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		return err
	}

	log.Printf("Exchange %s declared", exchangeName)
	return nil
}

func (m *MQ) CreateQueueAndBind(queueName, routingKey string, exchangeName string) error {
	_, err := m.Channel.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return err
	}

	err = m.Channel.QueueBind(
		queueName,    // queue name
		routingKey,   // routing key
		exchangeName, // exchange
		false,
		nil,
	)
	if err != nil {
		return err
	}

	log.Printf("Queue %s bound to exchange %s with routing key %s", queueName, exchangeName, routingKey)
	return nil
}

func (m *MQ) PublishMsg(exchange, queue string, body []byte) error {
	// 发送消息到RabbitMQ
	if err := m.Channel.Publish(exchange, queue, false, false, amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	}); err != nil {
		return err
	}
	return nil
}

// StartMq 处理评论信息
func (m *MQ) StartMq(exchange string, queue ...string) {
	// 声明交换机
	err := m.CreateExchange(exchange, "direct")
	if err != nil {
		err = m.Conn.Close()
		if err != nil {
			panic(err)
		}
	}
	for _, q := range queue {
		err = m.CreateQueueAndBind(q, q, exchange)
		if err != nil {
			err = m.Conn.Close()
			if err != nil {
				panic(err)
			}
		}
	}
}
