package test

import (
	"context"
	"github.com/simonalong/gole/util"
	"log"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
)

var fetchPressStreamName = "fetchPressStreamName"
var fetchPressSubjectAll = "fetch.press.subject.*"
var fetchPressSubject = "fetch.press.subject.key1"

var natsConnect *nats.Conn

var totalNum = 100
var totalSize = 10000

func TestPressProducer1(t *testing.T) {
	for i := 0; i < totalNum; i++ {
		sendMsg()
	}

	time.Sleep(1000000 * time.Hour)
}

func sendMsg() {
	if nil == natsConnect {
		nc, _ := nats.Connect("localhost:4222")
		natsConnect = nc
	}

	js, _ := natsConnect.JetStream()
	ctx := context.Background()

	info, err := js.StreamInfo(fetchPressStreamName)
	if nil == info {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:       fetchPressStreamName,
			Subjects:   []string{fetchPressSubjectAll},
			Retention:  nats.WorkQueuePolicy,
			Replicas:   1,
			Discard:    nats.DiscardOld,
			Duplicates: 30 * time.Second,
		}, nats.Context(ctx))
		if err != nil {
			log.Fatalf("can't add: %v", err)
		}
	}

	go func() {
		for i := 0; i < totalSize; i++ {
			// 同步：js.Publish(fetchPressSubject, []byte("message=="+util.ToString(i)), nats.Context(ctx))
			js.PublishAsync(fetchPressSubject, []byte("message=="+util.ToString(i)))
			//log.Printf("[publisher] sent %d", i)
			//time.Sleep(1 * time.Second)
		}
		log.Printf("send finish")
	}()
}

var pressCount1 = 0
var pressCount2 = 0

func TestPressConsumer1(t *testing.T) {
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
	js, _ := nc.JetStream()
	sub, _ := js.PullSubscribe(fetchPressSubject, "group")

	for {
		msgs, err := sub.Fetch(1)
		if nil != err {
			log.Printf("err %v", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		msg := msgs[0]
		pressCount1++
		if pressCount1%((totalNum*totalSize)/100) == 0 {
			log.Printf("[consumer: %s] received msg (%v) ratio: %s", id, string(msg.Data), util.ToString((pressCount1*100)/(totalNum*totalSize)))
		}
		msg.Ack()
	}
}

func TestPressConsumer2(t *testing.T) {
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
	js, _ := nc.JetStream()
	sub, _ := js.PullSubscribe(fetchPressSubject, "group")

	for {
		msgs, err := sub.Fetch(1)
		if nil != err {
			log.Printf("err %v", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		msg := msgs[0]
		pressCount2++
		if pressCount2%((totalNum*totalSize)/100) == 0 {
			log.Printf("[consumer: %s] received msg (%v) ratio: %s", id, string(msg.Data), util.ToString((pressCount2*100)/(totalNum*totalSize)))
		}
		msg.Ack()
	}
}
