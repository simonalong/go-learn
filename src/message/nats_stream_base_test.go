package test

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"os"
	"testing"
	"time"
)

const (
	streamName   = "stream-name5"
	consumerName = "consumer-name5"
	groupName    = "group-name5"
	tag          = "tag5"
)

func TestNatsJsBasePub1(t *testing.T) {
	js, _ := GetStreamOfSend(streamName, []string{tag + ".*"})

	// 发布信息
	_, err := js.Publish(tag+".key1", []byte("Hello World11"))
	if err != nil {
		return
	}
}

func TestNatsJsBaseSub1(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream()

	UseConsumer(js, streamName, consumerName, groupName)
	//
	//
	//sub, _ := js.SubscribeSync(tag+".key1")
	//msg, _ := sub.NextMsg(2 * time.Second)
	//fmt.Println("---->>>m3", string(msg.Data))
	//msg.Ack()

	sub, err := js.QueueSubscribeSync(tag+".key1", groupName, nats.Durable(consumerName), nats.MaxDeliver(3), nats.AckExplicit())
	m, err := sub.NextMsg(2 * time.Second)
	if err != nil {
		fmt.Println("--->>>3error", err)
		os.Exit(-1)
	}
	fmt.Println("---->>>m3", string(m.Data))
	m.Ack()
	sub.Unsubscribe()

	//// Simple Async Subscriber
	//_, err := js.Subscribe(tag+".key1", func(m *nats.Msg) {
	//	fmt.Printf("key1: Received a message: %s\n", string(m.Data))
	//	m.Ack()
	//}, nats.ManualAck())
	//
	//if err != nil {
	//	fmt.Printf("error, %v", err.Error())
	//	return
	//}
	//
	//// Create a druable consumer
	//js.DeleteConsumer("stream-name4", "consumer4")
	//
	////js.DeleteStream("stream-name1")

	time.Sleep(120000 * time.Second)

	nc.Close()
}

func UseConsumer(jetStreamContext nats.JetStreamContext, streamName, consumerName, group string) {
	if cinfo, _ := jetStreamContext.ConsumerInfo(streamName, consumerName); cinfo == nil {
		_, err := jetStreamContext.AddConsumer(streamName, &nats.ConsumerConfig{
			Durable:      consumerName,
			DeliverGroup: group,
			AckPolicy:    nats.AckExplicitPolicy,
		})
		if err != nil {
			fmt.Println("--->>>11baerror", err)
			os.Exit(-1)
		}
	}
}

func TestNatsJsBase0(t *testing.T) {
	timeout := 5 * time.Second
	// Connect to NATS
	nc, err := nats.Connect("nats://localhost:4222")
	if err != nil {
		fmt.Println("--->>>error", err)
		os.Exit(-1)
	}

	// Create JetStream Context
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		fmt.Println("--->>>1error", err)
		os.Exit(-1)
	}

	err = js.DeleteStream("ORDERS")
	if err != nil {
		fmt.Println("--->>>10aerror", err)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
	})
	if err != nil {
		fmt.Println("--->>>11aerror", err)
		os.Exit(-1)
	}

	sinfo, err := js.StreamInfo("ORDERS")
	if err != nil {
		fmt.Println("--->>>11aerror", err)
		os.Exit(-1)
	}
	fmt.Println("--->>stream info", sinfo.Config.Name, sinfo.Config.Subjects)

	if cinfo, _ := js.ConsumerInfo("ORDERS", "MONITOR"); cinfo == nil {
		_, err := js.AddConsumer("ORDERS", &nats.ConsumerConfig{
			Durable:      "MONITOR",
			DeliverGroup: "group",
			AckPolicy:    nats.AckExplicitPolicy,
		})
		if err != nil {
			fmt.Println("--->>>11baerror", err)
			os.Exit(-1)
		}
	}

	// Simple Stream Publisher
	_, err = js.Publish("ORDERS.bar", []byte("hellobar"))
	if err != nil {
		fmt.Println("--->>>1aerror", err)
		os.Exit(-1)
	}

	// Simple Async Stream Publisher
	for i := 0; i < 500; i++ {
		_, err := js.Publish("ORDERS.scratch", []byte(fmt.Sprintf("%s-%d", "hello", i)))
		if err != nil {
			fmt.Println("--->>>1berror", err)
			os.Exit(-1)
		}
	}
	select {
	case <-js.PublishAsyncComplete():
		fmt.Println("Publish complete")
	case <-time.After(5 * time.Second):
		fmt.Println("Did not resolve in time")
	}

	// Simple Sync Durable Consumer (optional SubOpts at the end)
	sub, err := js.QueueSubscribeSync("ORDERS.scratch", "group", nats.Durable("MONITOR"), nats.MaxDeliver(3), nats.AckExplicit())
	m, err := sub.NextMsg(timeout)
	if err != nil {
		fmt.Println("--->>>3error", err)
		os.Exit(-1)
	}
	fmt.Println("---->>>m3", string(m.Data))
	m.Ack()
	sub.Unsubscribe()

	// Simple Sync Durable Consumer (optional SubOpts at the end)
	sub, err = js.QueueSubscribeSync("ORDERS.scratch", "group", nats.Durable("MONITOR"), nats.MaxDeliver(3), nats.AckExplicit())
	m, err = sub.NextMsg(timeout)
	if err != nil {
		fmt.Println("--->>>3aerror", err)
		os.Exit(-1)
	}
	fmt.Println("---->>>m4", string(m.Data))
	m.Ack()
	sub.Unsubscribe()

	// Drain
	sub.Drain()
}
