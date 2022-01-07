package test

import (
	"context"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"github.com/simonalong/gole/util"
	"log"
	"testing"
	"time"

	"github.com/nats-io/nats.go"
)

// TestMessage is a message that can help test timings on jetstream
type TestMessage struct {
	ID          int       `json:"id"`
	PublishTime time.Time `json:"publish_time"`
}

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
}

var stream = streamName

func TestPubSub(t *testing.T) {
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()

	nctx := nats.Context(context.Background())

	info, _ := js.StreamInfo(stream)
	if nil == info {
		js.AddStream(&nats.StreamConfig{
			Name:     stream,
			Subjects: []string{subject},
		}, nctx)
	}

	// Custom context with timeout
	tctx, cancel := context.WithTimeout(nctx, 10*time.Second)
	// Set a timeout for publishing using context.
	deadlineCtx := nats.Context(tctx)
	// our resulting usec measurements
	results := make(chan int64)

	var totalTime int64
	var totalMessages int64

	go func() {
		i := 0
		for {
			js.Publish(subject, []byte("message---"+util.ToString(i)), deadlineCtx)
			time.Sleep(1 * time.Second)
			fmt.Println("[publisher] sent ", i)
			i++
		}
	}()

	for {
		select {
		case <-deadlineCtx.Done():
			cancel()
			js.DeleteStream(stream)
			return
		case usec := <-results:
			totalTime += usec
			totalMessages++
		}
	}
}

func TestNatsJsBaseSub(t *testing.T) {
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	sub, _ := js.QueueSubscribeSync(subject, "myqueuegroup", nats.Durable(id), nats.DeliverNew())

	// Custom context with timeout
	tctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	// Set a timeout for publishing using context.
	deadlineCtx := nats.Context(tctx)

	for {
		msg, err := sub.NextMsgWithContext(deadlineCtx)
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				break
			}
			time.Sleep(1 * time.Second)
			continue
		}

		fmt.Println("message == ", msg.Data)

		err = msg.Ack(deadlineCtx)
		if err != nil {
			fmt.Printf("[consumer: %s] error acking message: %v", id, err.Error())
		}
	}
}

func TestNatsJsBaseSub2(t *testing.T) {
	//func sub1(ctx nats.ContextOpt, subject string, results chan int64) error {
	id := uuid.NewV4().String()
	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		log.Fatalf("[%s] unable to connect to nats: %v", id, err)
	}

	var js nats.JetStream
	// Custom context with timeout
	tctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	deadlineCtx := nats.Context(tctx)

	js, _ = nc.JetStream()

	sub, _ := js.PullSubscribe(subject, "group")

	for {
		msgs, err := sub.Fetch(1, deadlineCtx)
		if len(msgs) > 0 {
			msg := msgs[0]

			fmt.Println("[consumer: ] received msg () after waiting", id, msgs)

			err = msg.Ack(deadlineCtx)
			if err != nil {
			}
		}
	}

	//for {
	//	msgs, err := sub.Fetch(1, nats.Context(ctx))
	//	if err != nil {
	//		if errors.Is(err, context.DeadlineExceeded) {
	//			break
	//		}
	//
	//		log.Printf("[consumer: %s] error consuming, sleeping for a second: %v", id, err)
	//		time.Sleep(1 * time.Second)
	//
	//		continue
	//	}
	//	msg := msgs[0]
	//
	//	var tMsg *TestMessage
	//
	//	err = json.Unmarshal(msg.Data, &tMsg)
	//	if err != nil {
	//		log.Printf("[consumer: %s] error consuming, sleeping for a second: %v", id, err)
	//		time.Sleep(1 * time.Second)
	//
	//		continue
	//	}
	//
	//	tm := time.Since(tMsg.PublishTime).Microseconds()
	//	results <- tm
	//
	//	log.Printf("[consumer: %s] received msg (%d) after waiting usec: %d", id, tMsg.ID, tm)
	//
	//	err = msg.Ack(nats.Context(ctx))
	//	if err != nil {
	//		log.Printf("[consumer: %s] error acking message: %v", id, err)
	//	}
	//
	//}

}
