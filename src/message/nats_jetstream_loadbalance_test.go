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

var fetchStreamName = "fetchStreamName"
var fetchSubjectAll = "fetch.subject.*"
var fetchSubject = "fetch.subject.key1"

func TestProducer1(t *testing.T) {
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	info, err := js.StreamInfo(fetchStreamName)
	if nil == info {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:       fetchStreamName,
			Subjects:   []string{fetchSubjectAll},
			Retention:  nats.WorkQueuePolicy,
			Replicas:   1,
			Discard:    nats.DiscardOld,
			Duplicates: 30 * time.Second,
		})
		if err != nil {
			log.Fatalf("can't add: %v", err)
		}
	}

	results := make(chan int64)
	var totalTime int64
	var totalMessages int64

	go func() {
		num := 100
		data := 10000
		for i := 0; i < num*data; i++ {
			js.PublishAsync(fetchSubject, []byte("message=="+util.ToString(i)), nats.Context(ctx))
			//log.Printf("[publisher] sent %d", i)
			//time.Sleep(1 * time.Second)
			i++
		}
		log.Printf("send finish")
	}()

	for {
		select {
		case <-ctx.Done():
			cancel()
			//	//log.Printf("sent %d messages with average time of %f", totalMessages, math.Round(float64(totalTime/totalMessages)))
			//	//js.DeleteStream(fetchStreamName)
			return
		case usec := <-results:
			totalTime += usec
			totalMessages++
		}
	}
}

var count1 = 0
var count2 = 0

func TestConsumer1(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
	js, _ := nc.JetStream()
	sub, _ := js.PullSubscribe(fetchSubject, "group")

	// todo，假设这个流没有关闭，感觉会有问题
	// todo，流如果关闭，则会有响应的问题

	for {
		// 默认6秒同步拉取，拉取不到，则会上报timeout
		msgs, err := sub.Fetch(1, nats.Context(ctx))
		if nil != err {
			log.Printf("err %v", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		msg := msgs[0]
		count1++
		if count1%1 == 0 {
			log.Printf("[consumer: %s] received msg (%v) ratio: %s", id, string(msg.Data), util.ToString((count1*100)/(100*10000)))
		}
		msg.Ack(nats.Context(ctx))
	}
}

func TestConsumer2(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
	js, _ := nc.JetStream()
	sub, _ := js.PullSubscribe(fetchSubject, "group")

	for {
		msgs, err := sub.Fetch(1, nats.Context(ctx))
		if nil != err {
			log.Printf("err %v", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		msg := msgs[0]
		count2++
		if count2%1 == 0 {
			log.Printf("[consumer: %s] received msg (%v) ratio: %s", id, string(msg.Data), util.ToString((count2*100)/(100*10000)))
		}

		msg.Ack(nats.Context(ctx))
	}
}
