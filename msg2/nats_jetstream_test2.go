package main

//
//import (
//	"context"
//	"encoding/json"
//	"errors"
//	"log"
//	"math"
//	"time"
//
//	"github.com/nats-io/nats.go"
//	uuid "github.com/satori/go.uuid"
//)
//
//// TestMessage is a message that can help test timings on jetstream
//type TestMessage struct {
//	ID          int       `json:"id"`
//	PublishTime time.Time `json:"publish_time"`
//}
//
//func init() {
//	log.SetFlags(log.LstdFlags | log.Lshortfile)
//}
//
//var stream = "streamName222"
//
//func main() {
//	//stream2 := uuid.NewV4().String()
//	// subject := fmt.Sprintf("%s-bar", id)
//	subject := stream
//
//	nc, err := nats.Connect("localhost:4222")
//	if err != nil {
//		log.Fatalf("unable to connect to nats: %v", err)
//	}
//
//	js, err := nc.JetStream()
//	if err != nil {
//		log.Fatalf("error getting jetstream: %v", err)
//	}
//
//	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Second)
//	defer cancel()
//
//	info, err := js.StreamInfo(stream)
//	if nil == info {
//		_, err = js.AddStream(&nats.StreamConfig{
//			Name:      stream,
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
//
//	var totalTime int64
//
//	var totalMessages int64
//
//	go func() {
//		err := sub(ctx, subject, results)
//		if err != nil {
//			log.Fatalf("%v", err)
//		}
//	}()
//
//	go func() {
//		err := sub(ctx, subject, results)
//		if err != nil {
//			log.Fatalf("%v", err)
//		}
//	}()
//
//	// our publisher thread
//	go func() {
//		i := 0
//
//		for {
//			start := time.Now()
//
//			bytes, err := json.Marshal(&TestMessage{
//				ID:          i,
//				PublishTime: start,
//			})
//			if err != nil {
//				log.Fatalf("could not get bytes from literal TestMessage... %v", err)
//			}
//
//			_, err = js.Publish(subject, bytes, nats.Context(ctx))
//			if err != nil {
//				if errors.Is(err, context.DeadlineExceeded) {
//					return
//				}
//
//				log.Printf("error publishing: %v", err)
//			}
//
//			log.Printf("[publisher] sent %d, publish time usec: %d", i, time.Since(start).Microseconds())
//			time.Sleep(1 * time.Second)
//
//			i++
//		}
//	}()
//
//	for {
//		select {
//		case <-ctx.Done():
//			cancel()
//			log.Printf("sent %d messages with average time of %f", totalMessages, math.Round(float64(totalTime/totalMessages)))
//			js.DeleteStream(stream)
//			return
//		case usec := <-results:
//			totalTime += usec
//			totalMessages++
//		}
//	}
//}
//
//func sub(ctx context.Context, subject string, results chan int64) error {
//	id := uuid.NewV4().String()
//
//	nc, err := nats.Connect("localhost:4222", nats.Name(id))
//	if err != nil {
//		log.Fatalf("[%s] unable to connect to nats: %v", id, err)
//	}
//
//	var js nats.JetStream
//
//	js, err = nc.JetStream()
//	if err != nil {
//		return err
//	}
//
//	sub, err := js.PullSubscribe(subject, "group")
//	if err != nil {
//		return err
//	}
//
//	for {
//		msgs, err := sub.Fetch(1, nats.Context(ctx))
//		if err != nil {
//			if errors.Is(err, context.DeadlineExceeded) {
//				break
//			}
//
//			log.Printf("[consumer: %s] error consuming, sleeping for a second: %v", id, err)
//			time.Sleep(1 * time.Second)
//
//			continue
//		}
//		msg := msgs[0]
//
//		var tMsg *TestMessage
//
//		err = json.Unmarshal(msg.Data, &tMsg)
//		if err != nil {
//			log.Printf("[consumer: %s] error consuming, sleeping for a second: %v", id, err)
//			time.Sleep(1 * time.Second)
//
//			continue
//		}
//
//		tm := time.Since(tMsg.PublishTime).Microseconds()
//		results <- tm
//
//		log.Printf("[consumer: %s] received msg (%d) after waiting usec: %d", id, tMsg.ID, tm)
//
//		err = msg.Ack(nats.Context(ctx))
//		//if err != nil {
//		//	log.Printf("[consumer: %s] error acking message: %v", id, err)
//		//}
//
//	}
//
//	return nil
//}
