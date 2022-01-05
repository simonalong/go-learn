package test

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"log"
	"testing"
	"time"
)

type myMessageHandler struct{}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		// In this case, a message with an empty body is simply ignored/discarded.
		return nil
	}

	// do whatever actual message processing is desired
	//err := processMessage(m.Body)
	fmt.Println(m.Body)

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return nil
}

func TestPub(t *testing.T) {
	cfg := nsq.NewConfig()
	// 连接 nsqd 的 tcp 连接
	//producer, err := nsq.NewProducer("127.0.0.1:4150", cfg)
	producer, err := nsq.NewProducer("127.0.0.1:32811", cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 发布消息
	var count int
	for {
		count++
		body := fmt.Sprintf("test %d", count)
		fmt.Println("发布消息：" + body)
		if err := producer.Publish("test", []byte(body)); err != nil {
			log.Fatal("publish error: " + err.Error())
		}
		time.Sleep(1 * time.Second)
	}
}

func TestSub(t *testing.T) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test", "levonfly", cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 处理信息
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Println(string(message.Body))
		return nil
	}))

	// 连接 nsqd 的 tcp 连接
	//if err := consumer.ConnectToNSQD("127.0.0.1:4150"); err != nil {
	if err := consumer.ConnectToNSQD("127.0.0.1:32811"); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
}
