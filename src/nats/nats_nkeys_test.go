package nats

import (
	"fmt"
	"github.com/isyscore/isc-gobase/logger"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

func TestNKeysSub(t *testing.T) {
	opt, err := nats.NkeyOptionFromSeed("./nkeys/seed.txt")
	if err != nil {
		logger.Info(err.Error())
		return
	}

	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL, opt)
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
func TestNKeysPub(t *testing.T) {
	// Connect to a server
	opt, err := nats.NkeyOptionFromSeed("./nkeys/seed.txt")
	if err != nil {
		logger.Info(err.Error())
		return
	}

	// Connect to a server
	nc, err := nats.Connect(nats.DefaultURL, opt)
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	// Simple Publisher
	err = nc.Publish("foo", []byte("Hello World"))
	if err != nil {
		return
	}

	nc.Close()
}
