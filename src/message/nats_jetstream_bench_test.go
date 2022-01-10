package test

import (
	"context"
	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
	"github.com/simonalong/gole/util"
	"log"
	"testing"
	"time"
)

var benchStreamName = "streambench"
var subjectAll = "subjectAll"

func BenchmarkTest1(b *testing.B) {
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	info, err := js.StreamInfo(benchStreamName)
	if nil == info {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:       benchStreamName,
			Subjects:   []string{subjectAll},
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
			for n := 0; n < b.N; n++ {
				js.Publish(benchStreamName, []byte("message - "+util.ToString(n)), nats.Context(ctx))
			}

			time.Sleep(1 * time.Second)
			i++
		}
	}()

	for {
		select {
		case <-ctx.Done():
			cancel()
			//log.Printf("sent %d messages with average time of %f", totalMessages, math.Round(float64(totalTime/totalMessages)))
			js.DeleteStream(benchStreamName)
			return
		case usec := <-results:
			totalTime += usec
			totalMessages++
		}
	}
}

func TestForBench(t *testing.T) {
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
	js, _ := nc.JetStream()
	sub, _ := js.PullSubscribe(benchStreamName, "group")

	for {
		msgs, err := sub.Fetch(1, nats.Context(ctx))
		if nil != err {
			log.Printf("err %v", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		msg := msgs[0]
		msg.Ack(nats.Context(ctx))
	}
}
