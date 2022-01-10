package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"time"

	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
	log "github.com/sirupsen/logrus"
)

type TestMessage1 struct {
	ID          int       `json:"id"`
	PublishTime time.Time `json:"publish_time"`
}

func main() {

	stream := uuid.NewV4().String()
	// subject := fmt.Sprintf("%s-bar", id)
	subject := stream

	fmt.Println(stream)

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

	go func() {
		err := sub1(deadlineCtx, subject, results)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}()

	go func() {
		err := sub1(deadlineCtx, subject, results)
		if err != nil {
			log.Fatalf("%v", err)
		}
	}()

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

func sub1(ctx nats.ContextOpt, subject string, results chan int64) error {
	id := uuid.NewV4().String()

	nc, err := nats.Connect("localhost:4222")
	if err != nil {
		log.Fatalf("[%s] unable to connect to nats: %v", id, err)
	}

	var js nats.JetStream

	js, err = nc.JetStream()
	if err != nil {
		return err
	}

	sub, err := js.QueueSubscribeSync(subject, "myqueuegroup", nats.Durable(id), nats.DeliverNew())
	if err != nil {
		return err
	}

	for {
		msg, err := sub.NextMsgWithContext(ctx)
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
		results <- tm

		log.Infof("[consumer: %s] received msg (%d) after waiting usec: %d", id, tMsg.ID, tm)

		err = msg.Ack(ctx)
		if err != nil {
			log.Errorf("[consumer: %s] error acking message: %v", id, err)
		}

	}

	return nil
}
