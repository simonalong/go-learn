package nats

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/simonalong/gole/util"
	"testing"
	"time"
)

// 发布消息
func TestPublish(t *testing.T) {
	subjectName := "topic.test4"
	js := getJs(subjectName)

	js.Publish(subjectName, []byte("你好"))

	for i := 0; i < 50; i++ {
		js.PublishAsync(subjectName, []byte("你好"+util.ToString(i)))
	}
	select {
	case <-js.PublishAsyncComplete():
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}
}

// pull消费
func TestConsumerPull(t *testing.T) {
	subjectName := "topic.test4"
	js := getJs(subjectName)

	sub, _ := js.PullSubscribe(subjectName, "wq", nats.PullMaxWaiting(128))

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}

		msgs, _ := sub.Fetch(10, nats.Context(ctx))
		for _, msg := range msgs {
			fmt.Println(string(msg.Data))
			msg.Ack()
		}
	}
}

// 推送消费：示例
func TestConsumerPush(t *testing.T) {
	subjectName := "topic.test4"
	js := getJs(subjectName)

	//js.Subscribe(subjectName, func(msg *nats.Msg) {
	//	//meta, _ := msg.Metadata()
	//	//fmt.Printf("Stream Sequence  : %v\n", meta.Sequence.Stream)
	//	//fmt.Printf("Consumer Sequence: %v\n", meta.Sequence.Consumer)
	//	fmt.Println(string(msg.Data))
	//})

	// Async subscriber with manual acks.
	js.Subscribe(subjectName, func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
		msg.Ack()
	}, nats.ManualAck())

	time.Sleep(5 * time.Millisecond)
}
