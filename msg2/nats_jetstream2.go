package main

import (
	"context"
	"fmt"
	"github.com/simonalong/gole/util"
	"math"
	"time"

	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

var streamSub = "streamDffffdsfdfdssffemo11"
var subAll = streamSub

func main() {

	//stream := uuid.NewV4().String()
	stream := streamSub
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
	if nil == info {
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     stream,
			Subjects: []string{subject},
		}, nctx)
		if err != nil {
			log.Fatalf("can't add: %v", err)
		}
	}

	// Custom context with timeout
	tctx, cancel := context.WithTimeout(nctx, 10000*time.Second)
	// Set a timeout for publishing using context.
	deadlineCtx := nats.Context(tctx)
	// our resulting usec measurements
	results := make(chan int64)

	var totalTime int64

	var totalMessages int64

	//go func() {
	//	err := sub4()
	//	if err != nil {
	//		log.Fatalf("%v", err)
	//	}
	//}()

	// our publisher thread
	go func() {
		i := 0
		for {
			js.Publish(subject, []byte("data "+util.ToString(i)), deadlineCtx)
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

	//nc, _ := nats.Connect("localhost:4222")
	//js, _ := nc.JetStream()
	//ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
	//defer cancel()
	//
	//info, err := js.StreamInfo(streamSub)
	//if nil == info {
	//	_, err = js.AddStream(&nats.StreamConfig{
	//		Name:       streamSub,
	//		Subjects:   []string{subAll},
	//		Retention:  nats.WorkQueuePolicy,
	//		Replicas:   1,
	//		Discard:    nats.DiscardOld,
	//		Duplicates: 30 * time.Second,
	//	}, nats.Context(ctx))
	//	if err != nil {
	//		log.Fatalf("can't add: %v", err)
	//	}
	//}
	//
	//results := make(chan int64)
	//var totalTime int64
	//var totalMessages int64
	//
	//go func() {
	//	err := sub4()
	//	if err != nil {
	//		log.Fatalf("%v", err)
	//	}
	//}()
	//
	//go func() {
	//	i := 0
	//	for {
	//		js.Publish(streamSub, []byte("message=="+util.ToString(i)), nats.Context(ctx))
	//		log.Printf("[publisher] sent %d", i)
	//		time.Sleep(1 * time.Second)
	//		i++
	//	}
	//}()
	//
	//for {
	//	select {
	//	case <-ctx.Done():
	//		cancel()
	//		log.Printf("sent %d messages with average time of %f", totalMessages, math.Round(float64(totalTime/totalMessages)))
	//		js.DeleteStream(streamSub)
	//		return
	//	case usec := <-results:
	//		totalTime += usec
	//		totalMessages++
	//	}
	//}
}

//
//func sub4() error {
//	ctx2, _ := context.WithTimeout(context.Background(), 1000*time.Second)
//	ctx := nats.Context(ctx2)
//	id := uuid.NewV4().String()
//	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
//	js, _ := nc.JetStream()
//	sub, _ := js.QueueSubscribeSync("streamDffffdfdssffemo11", "myqueuegroup", nats.Durable(id), nats.DeliverNew())
//
//	for {
//		msg, err := sub.NextMsgWithContext(ctx)
//		if nil != err {
//			log.Printf("err  sub4 %v", err.Error())
//			time.Sleep(1 * time.Second)
//			continue
//		}
//		log.Printf("[consumer sub4: %s] received msg (%v)", id, string(msg.Data))
//		msg.Ack(ctx)
//	}
//	return nil
//}
