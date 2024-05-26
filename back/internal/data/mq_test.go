/*
@Time : 2024/5/25 下午8:24
@Author : ljn
@File : mq_test
@Software: GoLand
*/

package data

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
	"sync"
	"testing"
)

func TestQueryMsgCount(t *testing.T) {
	c := ConnMq("amqp://root:123456@localhost:5672/heritage")
	q, err := c.Channel.QueueDeclarePassive(
		"apply_inheritor_queue", // 队列名称
		false,                   // 不持久化
		false,                   // 不自动删除
		false,                   // 不等待消息被发送
		false,                   // 不设置额外的参数
		nil,                     // 设置额外的参数
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Queue messages: %d", q.Messages)
}

func TestConsumeMsg(t *testing.T) {
	c := ConnMq("amqp://root:123456@localhost:5672/heritage")
	// 定义一个交付通道，用于接收消息
	deliveries, err := c.Channel.Consume(
		"apply_inheritor_queue", // 队列名称
		"",                      // 消费者名称，为空则由RabbitMQ自动生成
		false,                   // 自动确认消息
		false,                   // 不从服务器获取（不推送）未被ack的消息
		false,                   // 不等待服务器响应，立即开始接收消息
		false,                   // 不设置额外的参数
		nil,                     // 设置额外的参数
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	// 用于同步等待的 WaitGroup
	var wg sync.WaitGroup
	// 用于存储接收到的消息
	messages := make([]*amqp.Delivery, 0, 6)
	// 用于同步访问消息切片的互斥锁
	var mutex sync.Mutex

	// 启动一个goroutine来接收消息
	go func() {
		for d := range deliveries {
			mutex.Lock()
			messages = append(messages, &d)
			mutex.Unlock()
			wg.Done()
		}
	}()

	// 等待接收到10条消息
	wg.Add(6)
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-done
	for _, msg := range messages {
		fmt.Println(string(msg.Body))
		// 手动确认消息
		//if err = msg.Ack(false); err != nil {
		//	log.Fatalf("Failed to ack message: %v", err)
		//}
	}

	// 关闭通道和连接
	if err = c.Channel.Cancel("", false); err != nil {
		log.Fatalf("Failed to cancel channel: %v", err)
	}
	select {
	case <-make(chan interface{}):

	}
}
