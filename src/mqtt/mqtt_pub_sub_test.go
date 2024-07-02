package main

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

// 数据订阅
func TestMqttSub(t *testing.T) {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("demo1-sub-service")
	//opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:11883").SetClientID("demo-sub-service")
	opts.SetUsername("admin")
	opts.SetPassword("admin123@bbieat.com")
	opts.SetPingTimeout(1 * time.Second)
	opts.SetKeepAlive(60 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("connect fail", token.Error())
		return
	}

	// 订阅主题
	if token := c.Subscribe("nup/tenant/status", 0, func(client mqtt.Client, msg mqtt.Message) {
		var messageData = string(msg.Payload())
		fmt.Println("收到消息", messageData)
	}); token.Wait() && token.Error() != nil {
		fmt.Println("create topic fail")
		return
	}

	time.Sleep(100000 * time.Hour)
}

// 数据发布
func TestMqttPub(t *testing.T) {
	opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:1883").SetClientID("demo-pub-service")
	//opts := mqtt.NewClientOptions().AddBroker("tcp://localhost:11883").SetClientID("demo-pub-service")
	opts.SetUsername("admin")
	opts.SetPassword("admin123@bbieat.com")
	opts.SetPingTimeout(1 * time.Second)
	opts.SetKeepAlive(60 * time.Second)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("connect fail")
		return
	}

	num := 10
	for i := 0; i < num; i++ {
		text := fmt.Sprintf("Message %d", i)
		token := c.Publish("nup/tenant/status", 0, false, text)
		token.Wait()
		fmt.Println("发送数据 ", text)
		time.Sleep(time.Second)
	}
}

// nats：订阅
func TestNatsSubForMqtt(t *testing.T) {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Async Subscriber
	_, err := nc.Subscribe("nup.tenant.status", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		return
	}

	time.Sleep(100000 * time.Second)
	nc.Close()
}

// nats：发布
func TestNatsPub(t *testing.T) {
	// Connect to a server
	nc, _ := nats.Connect(nats.DefaultURL)

	// Simple Publisher
	err := nc.Publish("nup.tenant.status", []byte("Hello World"))
	if err != nil {
		return
	}

	for i := 0; i < 10; i++ {
		text := fmt.Sprintf("Hello World %d", i)
		nc.Publish("nup.tenant.status", []byte(text))
		fmt.Println("发送数据 ", text)
		time.Sleep(time.Second)
	}

	nc.Close()
}
