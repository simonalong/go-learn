package test

import (
	"context"
	"github.com/simonalong/gole/util"
	"log"
	"math"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
)

var fetchStreamName = "fetchStream"
var fetchSubjectAll = "fetch.subject.*"
var fetchSubject = "fetch.subject.key1"

func TestProducer1(t *testing.T) {
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
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
			js.Publish(fetchSubject, []byte("message=="+util.ToString(i)), nats.Context(ctx))
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
			js.DeleteStream(fetchStreamName)
			return
		case usec := <-results:
			totalTime += usec
			totalMessages++
		}
	}
}

func TestConsumer1(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
	js, _ := nc.JetStream()
	sub, _ := js.PullSubscribe(fetchSubject, "group")

	for {
		msgs, _ := sub.Fetch(1, nats.Context(ctx))
		msg := msgs[0]
		log.Printf("[consumer: %s] received msg (%v)", id, string(msg.Data))
		msg.Ack(nats.Context(ctx))
	}
}

func TestConsumer2(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
	js, _ := nc.JetStream()
	sub, _ := js.PullSubscribe(fetchSubject, "group")

	for {
		msgs, _ := sub.Fetch(1, nats.Context(ctx))
		msg := msgs[0]
		log.Printf("[consumer: %s] received msg (%v)", id, string(msg.Data))
		msg.Ack(nats.Context(ctx))
	}
}
