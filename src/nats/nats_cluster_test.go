package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

func TestNatsClusterSub(t *testing.T) {

	// Do something with the connection
	//nc, err := nats.Connect("nats://127.0.0.1:4222,nats://127.0.0.1:4223,nats://127.0.0.1:4224")
	nc, err := nats.Connect("nats://127.0.0.1:4222", nats.UserInfo("admin", "admin-demo123@bbieat.com"))
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

func TestNatsClusterPub(t *testing.T) {
	//nc, err := nats.Connect("nats://127.0.0.1:4222,nats://127.0.0.1:4223,nats://127.0.0.1:4224")
	nc, err := nats.Connect("nats://127.0.0.1:4224", nats.UserInfo("admin", "admin-demo123@bbieat.com"))
	if err != nil {
		// 使用token后这个报错：nats: Authorization Violation
		fmt.Println(err.Error())
	}

	nc.Publish("foo", []byte("Hello World"))
	nc.Close()
}
