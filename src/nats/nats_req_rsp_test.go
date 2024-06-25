package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

// 请求响应模型：响应1
func TestNatsRsp1(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	// 接收到请求，并返回响应
	_, err := nc.Subscribe("request", func(m *nats.Msg) {
		m.Respond([]byte("answer is 111"))
	})
	if err != nil {
		return
	}

	time.Sleep(100000 * time.Second)
	nc.Close()
}

// 请求响应模型：响应2
func TestNatsRsp2(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	// 接收到请求，并返回响应
	_, err := nc.Subscribe("request", func(m *nats.Msg) {
		m.Respond([]byte("answer is 000"))
	})
	if err != nil {
		return
	}

	time.Sleep(100000 * time.Second)
	nc.Close()
}

// 请求响应模型：请求
func TestNatsReq(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	for i := 0; i < 5; i++ {
		// 发出请求，并获取响应
		msg, err := nc.Request("request", []byte("help me"), 10*time.Millisecond)
		if err != nil {
			// 如果超时，则这里返回
			return
		}
		fmt.Println(string(msg.Data))
	}
	nc.Close()
}
