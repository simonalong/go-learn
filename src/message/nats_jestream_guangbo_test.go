package test

import (
	"context"
	"github.com/lunny/log"
	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
	"github.com/simonalong/gole/util"
	"math"
	"testing"
	"time"
)

var broadcastStreamSub = "stream"
var broadcastSubAll = "broadcast.subject.*"
var broadcastSub = "broadcast.subject.key1"

func TestSend1(t *testing.T) {
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	info, err := js.StreamInfo(broadcastStreamSub)
	if nil == info {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:       broadcastStreamSub,
			Subjects:   []string{broadcastSubAll},
			Retention:  nats.WorkQueuePolicy,
			Replicas:   1,
			Discard:    nats.DiscardOld,
			Duplicates: 30 * time.Second,
		}, nats.Context(ctx))
		if err != nil {
			log.Fatalf("can't add: %v", err)
		}
	}

	results := make(chan int64)
	var totalTime int64
	var totalMessages int64

	go func() {
		i := 0
		for {
			js.Publish(broadcastStreamSub, []byte("message=="+util.ToString(i)), nats.Context(ctx))
			log.Printf("[publisher] sent %d", i)
			time.Sleep(1 * time.Second)
			i++
		}
	}()

	for {
		select {
		case <-ctx.Done():
			cancel()
			log.Printf("sent %d messages with average time of %f", totalMessages, math.Round(float64(totalTime/totalMessages)))
			js.DeleteStream(broadcastStreamSub)
			return
		case usec := <-results:
			totalTime += usec
			totalMessages++
		}
	}
}

func TestConsumer11(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
	js, _ := nc.JetStream()
	sub, _ := js.QueueSubscribeSync(broadcastStreamSub, "myqueuegroup", nats.Durable(id), nats.DeliverNew())

	for {
		msg, err := sub.NextMsgWithContext(ctx)
		if err != nil {
			log.Printf("err %v", err.Error())
			continue
		}
		log.Printf("[consumer: %s] received msg (%v)", id, string(msg.Data))
		msg.Ack(nats.Context(ctx))
	}
}

func TestConsumer12(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
	js, _ := nc.JetStream()
	sub, _ := js.QueueSubscribeSync(broadcastStreamSub, "myqueuegroup", nats.Durable(id), nats.DeliverNew())

	for {
		msg, _ := sub.NextMsgWithContext(ctx)
		log.Printf("[consumer: %s] received msg (%v)", id, string(msg.Data))
		msg.Ack(nats.Context(ctx))
	}
}
