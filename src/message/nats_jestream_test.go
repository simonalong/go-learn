package test

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/lunny/log"
	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
	"github.com/simonalong/gole/util"
	"math"
	"testing"
	"time"
)

type TestMessage1 struct {
	ID          int       `json:"id"`
	PublishTime time.Time `json:"publish_time"`
}

func TestSend(t *testing.T) {
	//stream := uuid.NewV4().String()
	stream := "asdfasdfasdfasdf"
	// subject := fmt.Sprintf("%s-bar", id)
	subject := stream

	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		log.Fatalf("unable to connect to nats: %v", err)
	}

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalf("error getting jetstream: %v", err)
	}

	nctx := nats.Context(context.Background())

	info, err := js.StreamInfo(stream)
	if err == nil {
		log.Fatalf("Stream already exists: %v", info)
	}

	_, err = js.AddStream(&nats.StreamConfig{
		Name:     stream,
		Subjects: []string{subject},
	}, nctx)
	if err != nil {
		log.Fatalf("can't add: %v", err)
	}

	// Custom context with timeout
	tctx, cancel := context.WithTimeout(nctx, 10000*time.Second)
	// Set a timeout for publishing using context.
	deadlineCtx := nats.Context(tctx)
	// our resulting usec measurements
	results := make(chan int64)

	var totalTime int64

	var totalMessages int64

	// our publisher thread
	go func() {
		i := 0

		for {
			start := time.Now()

			bytes, err := json.Marshal(&TestMessage1{
				ID:          i,
				PublishTime: start,
			})
			if err != nil {
				log.Fatalf("could not get bytes from literal TestMessage... %v", err)
			}

			_, err = js.Publish(subject, bytes, deadlineCtx)
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					return
				}

				log.Errorf("error publishing: %v", err)
			}

			log.Infof("[publisher] sent %d, publish time usec: %d", i, time.Since(start).Microseconds())
			time.Sleep(1 * time.Second)

			i++
		}
	}()

	for {
		select {
		case <-deadlineCtx.Done():
			cancel()
			log.Infof("sent %d messages with average time of %f", totalMessages, math.Round(float64(totalTime/totalMessages)))
			js.DeleteStream(stream)
			return
		case usec := <-results:
			totalTime += usec
			totalMessages++
		}
	}
}

func TestSub1(t *testing.T) {
	subject := "asdfasdfasdfasdf"
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	id := uuid.NewV4().String()

	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		log.Fatalf("[%s] unable to connect to nats: %v", id, err)
	}

	sub, err := js.QueueSubscribeSync(subject, "myqueuegroup", nats.Durable(id), nats.DeliverNew())
	if err != nil {
		return
	}

	for {
		msg, err := sub.NextMsgWithContext(nats.Context(ctx))
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				break
			}

			log.Errorf("[consumer: %s] error consuming, sleeping for a second: %v", id, err)
			time.Sleep(1 * time.Second)

			continue
		}

		var tMsg *TestMessage1

		err = json.Unmarshal(msg.Data, &tMsg)
		if err != nil {
			log.Errorf("[consumer: %s] error consuming, sleeping for a second: %v", id, err)
			time.Sleep(1 * time.Second)

			continue
		}

		tm := time.Since(tMsg.PublishTime).Microseconds()

		log.Infof("[consumer: %s] received msg (%d) after waiting usec: %d", id, tMsg.ID, tm)

		err = msg.Ack(nats.Context(ctx))
		if err != nil {
			log.Errorf("[consumer: %s] error acking message: %v", id, err)
		}

	}
}

func TestSub2(t *testing.T) {
	subject := "asdfasdfasdfasdf"
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()
	id := uuid.NewV4().String()

	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		log.Fatalf("[%s] unable to connect to nats: %v", id, err)
	}

	sub, err := js.QueueSubscribeSync(subject, "myqueuegroup", nats.Durable(id), nats.DeliverNew())
	if err != nil {
		return
	}

	for {
		msg, err := sub.NextMsgWithContext(nats.Context(ctx))
		if err != nil {
			if errors.Is(err, context.DeadlineExceeded) {
				break
			}

			log.Errorf("[consumer: %s] error consuming, sleeping for a second: %v", id, err)
			time.Sleep(1 * time.Second)

			continue
		}

		var tMsg *TestMessage1

		err = json.Unmarshal(msg.Data, &tMsg)
		if err != nil {
			log.Errorf("[consumer: %s] error consuming, sleeping for a second: %v", id, err)
			time.Sleep(1 * time.Second)

			continue
		}

		tm := time.Since(tMsg.PublishTime).Microseconds()

		log.Infof("[consumer: %s] received msg (%d) after waiting usec: %d", id, tMsg.ID, tm)

		err = msg.Ack(nats.Context(ctx))
		if err != nil {
			log.Errorf("[consumer: %s] error acking message: %v", id, err)
		}

	}
}

func TestSend1(t *testing.T) {
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	info, err := js.StreamInfo(streamName)
	if nil == info {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:       streamName,
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
			js.DeleteStream(streamName)
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
	sub, _ := js.QueueSubscribeSync(subject, "myqueuegroup", nats.Durable(id), nats.DeliverNew())

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
	sub, _ := js.QueueSubscribeSync(subject, "myqueuegroup", nats.Durable(id), nats.DeliverNew())

	for {
		msg, _ := sub.NextMsgWithContext(ctx)
		log.Printf("[consumer: %s] received msg (%v)", id, string(msg.Data))
		msg.Ack(nats.Context(ctx))
	}
}
