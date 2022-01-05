package test

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

func TestNatsSub11(t *testing.T) {
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

func TestNatsSub12(t *testing.T) {
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

// 发布订阅模型
func TestNatsPub1(t *testing.T) {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Publisher
	err := nc.Publish("foo", []byte("Hello World"))
	if err != nil {
		return
	}

	nc.Close()
}

// 请求响应模型：请求
func TestNatsPub2(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	// 发出请求，并获取响应
	msg, err := nc.Request("request", []byte("help me"), 10*time.Millisecond)
	if err != nil {
		// 如果超时，则这里返回
		return
	}

	fmt.Println(string(msg.Data))

	nc.Close()
}

// 请求响应模型：响应1
func TestNatsSub21(t *testing.T) {
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
func TestNatsSub22(t *testing.T) {
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

func TestNatsPub3(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	// 发出请求，并获取响应
	err := nc.Publish("send3", []byte("help me"))
	if err != nil {
		return
	}

	nc.Close()
}

//
func TestNatsSub31(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	// 通道订阅
	ch := make(chan *nats.Msg, 64)
	nc.ChanSubscribe("send3", ch)
	msg := <-ch

	fmt.Println(string(msg.Data))

	// Close connection
	nc.Close()
}

// 消息的层级结构，通过subj进行处理，支持*和>
// * 匹配单个
// type.*.tag 匹配 type.key1.tag、type.key2.tag等
// type.*     匹配 type.key1、type.key2等，但是不匹配 type.key1.tag
// > 匹配多个
// type.>     匹配 type.key1.tag、type.key1.tag、type.key2.tag也匹配type.key1、type.key2等
func TestNatsPub4(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	// 发出请求，并获取响应
	err := nc.Publish("type.key.tag", []byte("help me"))
	if err != nil {
		return
	}

	nc.Close()
}

func TestNatsSub41(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	// 通道订阅
	ch := make(chan *nats.Msg, 64)
	nc.ChanSubscribe("type.key.tag", ch)
	msg := <-ch

	fmt.Println(string(msg.Data))

	// Close connection
	nc.Close()
}

func TestNatsSub42(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	// 通道订阅
	ch := make(chan *nats.Msg, 64)
	nc.ChanSubscribe("type.*.tag", ch)
	msg := <-ch

	fmt.Println(string(msg.Data))

	// Close connection
	nc.Close()
}
