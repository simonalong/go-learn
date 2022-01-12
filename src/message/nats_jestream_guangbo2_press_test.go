package test

import (
	"context"
	"github.com/lunny/log"
	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
	"github.com/simonalong/gole/util"
	"testing"
	"time"
)

var indexPress = "-index1"
var broadcastPressStreamName = "broadcastStreamName" + indexPress
var broadcastPressSubjectAll = "broadcast.subject" + indexPress + ".*"
var broadcastPressSubject = "broadcast.subject" + indexPress + ".key"

var totalBNum = 100
var totalBSize = 10000

func TestBroadcastPressProducer1(t *testing.T) {
	for i := 0; i < totalBNum; i++ {
		sendBMsg()
	}

	time.Sleep(1000000 * time.Hour)
}

func sendBMsg() {
	if nil == natsConnect {
		nc, _ := nats.Connect("localhost:4222")
		natsConnect = nc
	}

	js, _ := natsConnect.JetStream()
	ctx := context.Background()

	info, err := js.StreamInfo(broadcastPressStreamName)
	if nil == info {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     broadcastPressStreamName,
			Subjects: []string{broadcastPressSubjectAll},
		}, nats.Context(ctx))
		if err != nil {
			log.Fatalf("can't add: %v", err)
		}
	}

	go func() {
		for i := 0; i < totalBSize; i++ {
			// 同步：
			//js.Publish(broadcastPressSubject, []byte("message=="+util.ToString(i)), nats.Context(ctx))
			js.PublishAsync(broadcastPressSubject, []byte("message=="+util.ToString(i)))
		}
		log.Printf("send finish")
	}()
}

var pressBCount1 = 0
var pressBCount2 = 0

func TestDemoPressConsumer12(t *testing.T) {
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
	js, _ := nc.JetStream()
	sub, _ := js.QueueSubscribeSync(broadcastPressSubject, "myqueuegroup", nats.Durable(id), nats.DeliverNew())

	for {
		msg, err := sub.NextMsgWithContext(nats.Context(context.Background()))
		if nil != err {
			log.Printf("err  sub4 %v", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		pressBCount1++
		if pressBCount1%((totalBNum*totalBSize)/100) == 0 {
			log.Printf("[consumer: %s] received msg (%v) ratio: %s", id, string(msg.Data), util.ToString((pressBCount1*100)/(totalBNum*totalBSize)))
		}
		msg.Ack()
	}
}

func TestDemoPressConsumer13(t *testing.T) {
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
	js, _ := nc.JetStream()
	sub, _ := js.QueueSubscribeSync(broadcastPressSubject, "myqueuegroup", nats.Durable(id), nats.DeliverNew())

	for {
		msg, err := sub.NextMsgWithContext(nats.Context(context.Background()))
		if nil != err {
			log.Printf("err  sub4 %v", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		pressBCount2++
		if pressBCount2%((totalBNum*totalBSize)/100) == 0 {
			log.Printf("[consumer: %s] received msg (%v) ratio: %s", id, string(msg.Data), util.ToString((pressBCount2*100)/(totalBNum*totalBSize)))
		}
		msg.Ack()
	}
}
