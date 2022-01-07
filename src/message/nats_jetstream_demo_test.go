package test

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/simonalong/gole/util"
	"log"
	"math"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
)

// TestMessage is a message that can help test timings on jetstream

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

func TestDemo(t *testing.T) {
	//stream := uuid.NewV4().String()
	// subject := fmt.Sprintf("%s-bar", id)
	subject := stream

	nc, _ := nats.Connect("localhost:4222")

	js, _ := nc.JetStream()
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	info, err := js.StreamInfo(stream)
	if nil == info {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:      stream,
			Subjects:  []string{subject},
			Retention: nats.WorkQueuePolicy,
		}, nats.Context(ctx))
		if err != nil {
			log.Fatalf("can't add: %v", err)
		}
	}

	// Our resulting use measurements
	results := make(chan int64)
	var totalTime int64
	var totalMessages int64

	go func() {
		err := sub(ctx, subject, results)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}()

	// our publisher thread
	go func() {
		i := 0
		for {
			js.Publish(subject, []byte("message=="+util.ToString(i)), nats.Context(ctx))
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
			js.DeleteStream(stream)
			return
		case usec := <-results:
			totalTime += usec
			totalMessages++
		}
	}
}

func sub(ctx context.Context, subject string, results chan int64) error {
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
	js, _ := nc.JetStream()
	sub, _ := js.PullSubscribe(subject, "group")

	for {
		msgs, err := sub.Fetch(1, nats.Context(ctx))
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				break
			}

			log.Printf("[consumer: %s] error consuming, sleeping for a second: %v", id, err)
			time.Sleep(1 * time.Second)

			continue
		}
		msg := msgs[0]

		var tMsg *TestMessage

		err = json.Unmarshal(msg.Data, &tMsg)
		if err != nil {
			log.Printf("[consumer: %s] error consuming, sleeping for a second: %v", id, err)
			time.Sleep(1 * time.Second)

			continue
		}

		tm := time.Since(tMsg.PublishTime).Microseconds()
		results <- tm

		log.Printf("[consumer: %s] received msg (%d) after waiting usec: %d", id, tMsg.ID, tm)

		err = msg.Ack(nats.Context(ctx))
		//if err != nil {
		//	log.Printf("[consumer: %s] error acking message: %v", id, err)
		//}

	}

	return nil
}
