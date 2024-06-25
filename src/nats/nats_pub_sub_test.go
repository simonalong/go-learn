package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

// 发布订阅模型：订阅1
func TestNatsSub1(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	// Simple Async Subscriber
	_, err = nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		return
	}

	time.Sleep(100000 * time.Second)
	nc.Close()
}

// 发布订阅模型：订阅2
func TestNatsSub2(t *testing.T) {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Async Subscriber
	_, err := nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		return
	}

	time.Sleep(100000 * time.Second)
	nc.Close()
}

// 发布订阅模型：发布
func TestNatsPub(t *testing.T) {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Publisher
	err := nc.Publish("foo", []byte("Hello World"))
	if err != nil {
		return
	}

	nc.Close()
}
