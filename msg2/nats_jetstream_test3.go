package main

import (
	"context"
	"errors"
	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
	"github.com/simonalong/gole/util"
	"log"
	"math"
	"time"
)

// TestMessage is a message that can help test timings on jetstream

//func init() {
//	log.SetFlags(log.LstdFlags | log.Lshortfile)
//}
//
const (
	index      = "a3"
	streamName = "stream-name" + index
	subjectAll = "subject.*"
	subject    = "subject." + index
	consumer   = "consumer1"
	group      = "groupname"
)

func main() {
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	defer cancel()

	info, err := js.StreamInfo(streamName)
	if nil == info {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:      streamName,
			Subjects:  []string{subject},
			Retention: nats.WorkQueuePolicy,
		}, nats.Context(ctx))
		if err != nil {
			log.Fatalf("can't add: %v", err)
		}
	}

	results := make(chan int64)
	var totalTime int64
	var totalMessages int64

	go func() {
		err := sub()
		if err != nil {
			log.Fatalf("%v", err)
		}
	}()

	// our publisher thread
	go func() {
		i := 0
		for {
			_, err = js.Publish(subject, []byte("message="+util.ToString(i)), nats.Context(ctx))
			if err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					return
				}
				log.Printf("error publishing: %v", err)
			}

			log.Printf("[publisher] sent %d,", i)
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

func sub() error {
	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
	var js nats.JetStream
	js, _ = nc.JetStream()
	sub, err := js.PullSubscribe(subject, "group")
	if err != nil {
		return err
	}

	for {
		msgs, _ := sub.Fetch(1, nats.Context(ctx))
		msg := msgs[0]
		log.Printf("[consumer: %s] received msg (%v)", id, string(msg.Data))
		err = msg.Ack(nats.Context(ctx))
	}

	return nil
}

//func main() {
//	nc, _ := nats.Connect("localhost:4222")
//
//	js, _ := nc.JetStream()
//	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
//	defer cancel()
//
//	info, err := js.StreamInfo(streamName)
//	if nil == info {
//		_, err = js.AddStream(&nats.StreamConfig{
//			Name:      streamName,
//			Subjects:  []string{subject},
//			Retention: nats.WorkQueuePolicy,
//		}, nats.Context(ctx))
//		if err != nil {
//			log.Fatalf("can't add: %v", err)
//		}
//	}
//
//	// Our resulting use measurements
//	results := make(chan int64)
//	var totalTime int64
//	var totalMessages int64
//
//	go func() {
//		err := sub()
//		if err != nil {
//			log.Fatalf("%v", err)
//		}
//	}()
//
//	go func() {
//		i := 0
//		for {
//			js.Publish(subject, []byte("message=="+util.ToString(i)), nats.Context(ctx))
//			log.Printf("[publisher] sent %d", i)
//			time.Sleep(1 * time.Second)
//			i++
//		}
//	}()
//
//	for {
//		select {
//		case <-ctx.Done():
//			cancel()
//			log.Printf("sent %d messages with average time of %f", totalMessages, math.Round(float64(totalTime/totalMessages)))
//			js.DeleteStream(streamName)
//			return
//		case usec := <-results:
//			totalTime += usec
//			totalMessages++
//		}
//	}
//}
//
//func sub() error {
//	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
//	id := uuid.NewV4().String()
//	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
//	var js nats.JetStream
//	js, _ = nc.JetStream()
//	sub, err := js.PullSubscribe(subject, "group")
//	if err != nil {
//		return err
//	}
//
//	for {
//		msgs, _ := sub.Fetch(1, nats.Context(ctx))
//		msg := msgs[0]
//		log.Printf("[consumer: %s] received msg (%v)", id, string(msg.Data))
//		err = msg.Ack(nats.Context(ctx))
//	}
//
//	return nil
//}
