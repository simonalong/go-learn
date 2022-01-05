package test

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/simonalong/gole/util"
	"testing"
	"time"
)

func TestNatsJsBasePub1(t *testing.T) {
	js, _ := GetStreamOfSend("stream-name1", []string{"tag1.*"})

	// 发布信息
	_, err := js.Publish("tag1.key1", []byte("Hello World11"))
	if err != nil {
		return
	}
}

func TestNatsJsBaseSub1(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream()

	consumerInfo, _ := js.AddConsumer("stream-name1", &nats.ConsumerConfig{
		AckPolicy: nats.AckAllPolicy,
	})

	fmt.Println(util.ToJsonString(consumerInfo))

	// Simple Async Subscriber
	_, err := js.Subscribe("tag1.key1", func(m *nats.Msg) {
		fmt.Printf("key1: Received a message: %s\n", string(m.Data))
		m.Ack()
	})
	if err != nil {
		fmt.Printf("error, %v", err.Error())
		return
	}

	time.Sleep(100000 * time.Second)
	nc.Close()
}
