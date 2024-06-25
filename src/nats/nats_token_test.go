package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

func TestNatsTokenSub(t *testing.T) {
	// Connect to a server
	nc, err := nats.Connect("nats://127.0.0.1:4222", nats.Token("seatak-demo@bbieat.com"))
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	_, err = nc.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Println(err.Error())
	}

	time.Sleep(100000 * time.Second)
}

func TestTokenPub(t *testing.T) {
	nc, _ := nats.Connect("nats://127.0.0.1:4222", nats.Token("seatak-demo@bbieat.com"))

	nc.Publish("foo", []byte("Hello World"))
	nc.Close()
}
