package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
)

func TestNatsPub3(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	// 发出请求，并获取响应
	err := nc.Publish("send3", []byte("help me"))
	if err != nil {
		return
	}

	nc.Close()
}

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
