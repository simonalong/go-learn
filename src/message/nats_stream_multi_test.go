package test

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

func TestNatsJsMultiPub1(t *testing.T) {
	js, _ := GetStreamOfSend("stream-name", []string{"tag.*"})

	// 发布信息
	_, err := js.Publish("tag.key1", []byte("Hello World11"))
	if err != nil {
		return
	}
}

func TestNatsJsMultiPub2(t *testing.T) {
	js, _ := GetStreamOfSend("stream-name", []string{"tag.*"})

	// 发布信息
	_, err := js.Publish("tag.key1", []byte("Hello World21"))
	if err != nil {
		return
	}

	// 发布信息
	_, err = js.Publish("tag.key2", []byte("Hello World22"))
	if err != nil {
		return
	}
}

func TestNatsJsMultiSub1(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream()

	// Simple Async Subscriber
	_, err := js.Subscribe("tag.key1", func(m *nats.Msg) {
		fmt.Printf("key1: Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Printf("error, %v", err.Error())
		return
	}

	time.Sleep(100000 * time.Second)
	nc.Close()
}

func TestNatsJsMultiSub2(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream()

	// Simple Async Subscriber
	_, err := js.Subscribe("tag.key2", func(m *nats.Msg) {
		fmt.Printf("key2: Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Printf("error, %v", err.Error())
		return
	}

	time.Sleep(100000 * time.Second)
	nc.Close()
}

func GetStreamOfSend(streamName string, subjects []string) (nats.JetStreamContext, error) {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream()

	// 创建流通道
	info, _ := js.StreamInfo(streamName)
	if nil == info {
		_, err := js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: subjects,
		})
		if err != nil {
			return nil, err
		}
	}
	return js, nil
}
