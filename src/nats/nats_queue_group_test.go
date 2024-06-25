package nats

import (
	"fmt"
	"github.com/isyscore/isc-gobase/logger"
	"github.com/nats-io/nats.go"
	"strconv"
	"testing"
	"time"
)

// 队列分组模型：请求
func TestNatsQueueGroupSub(t *testing.T) {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	for i := 0; i < 20; i++ {
		// Simple Publisher
		err := nc.Publish("queue", []byte("Hello World "+strconv.Itoa(i)))
		if err != nil {
			return
		}
	}

	nc.Close()
}

// 队列分组模型：消费者1
func TestQueueGroup1(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer nc.Close()

	if _, err := nc.QueueSubscribe("queue", "workers", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	}); err != nil {
		logger.Fatal(err.Error())
	}
	time.Sleep(100000 * time.Second)
}

// 队列分组模型：消费者2
func TestQueueGroup2(t *testing.T) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer nc.Close()

	if _, err := nc.QueueSubscribe("queue", "workers", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	}); err != nil {
		logger.Fatal(err.Error())
	}
	time.Sleep(100000 * time.Second)
}
