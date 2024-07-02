package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

func TestCertSub(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL, nats.UserCredentials("/Users/zhouzhenyong/.local/share/nats/nsc/keys/creds/OperatorTest/AccountTest/UserTest.creds"))
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	// Simple Async Subscriber
	_, err = nc.Subscribe(">", func(m *nats.Msg) {

		fmt.Printf("Received a message: %s - %s\n", m.Subject, string(m.Data))
	})
	if err != nil {
		return
	}

	time.Sleep(100000 * time.Second)
	nc.Close()
}

// 发布订阅模型：订阅2
func TestCertPub(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL, nats.UserCredentials("/Users/zhouzhenyong/.local/share/nats/nsc/keys/creds/OperatorTest/AccountTest/UserTest.creds"))
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	// Simple Publisher
	err = nc.Publish("foo.sdf.sdf", []byte("Hello World"))
	if err != nil {
		return
	}

	nc.Close()
}
