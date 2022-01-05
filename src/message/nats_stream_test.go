package test

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

func TestNatsJsPub1(t *testing.T) {
	js, _ := GetStreamOfSend("stream-name", []string{"tag.*"})

	// 发布信息
	_, err := js.Publish("tag.key1", []byte("Hello World"))
	if err != nil {
		return
	}
}

func TestNatsJsSub1(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream()

	// 创建消费者
	js.AddConsumer("stream-name", &nats.ConsumerConfig{
		// 所有消息都确认
		AckPolicy: nats.AckAllPolicy,
		// 只消费最后一次
		DeliverPolicy: nats.DeliverLastPolicy,
	})

	// Simple Async Subscriber
	_, err := js.Subscribe("tag.key1", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Printf("error, %v", err.Error())
		return
	}

	time.Sleep(100000 * time.Second)
	nc.Close()
}

func GetStreamOfSend(streamName string, subjects []string) (nats.JetStreamContext, error) {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream()

	// 创建流通道
	info, _ := js.StreamInfo("stream-test")
	if nil == info {
		_, err := js.AddStream(&nats.StreamConfig{
			Name:     "stream-test",
			Subjects: []string{"test.*"},
		})
		if err != nil {
			return nil, err
		}
	}
	return js, nil
}

func TestNatsJsPub2(t *testing.T) {
	js, _ := GetStreamOfSend("stream-name2", []string{"tag.*"})

	// 发布信息
	_, err := js.Publish("test.tag1", []byte("Hello World1"))
	if err != nil {
		return
	}
}

func TestNatsJsSub2(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream()

	// 创建消费者
	js.AddConsumer("stream-test", &nats.ConsumerConfig{
		// 所有消息都确认
		AckPolicy: nats.AckAllPolicy,
		// 只消费最后一次
		DeliverPolicy: nats.DeliverLastPolicy,
	})

	// Simple Async Subscriber
	js.Subscribe("test.tag1", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	time.Sleep(100000 * time.Second)
	nc.Close()
}

func TestNatsJsSub2and2(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream()

	// 创建消费者
	js.AddConsumer("stream-test", &nats.ConsumerConfig{
		// 所有消息都确认
		AckPolicy: nats.AckAllPolicy,
		// 只消费最后一次
		DeliverPolicy: nats.DeliverLastPolicy,
	})

	// Simple Async Subscriber
	js.Subscribe("test.tag2", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})

	time.Sleep(100000 * time.Second)
	nc.Close()
}
