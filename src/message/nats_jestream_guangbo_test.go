package test

//
//import (
//	"context"
//	"github.com/lunny/log"
//	"github.com/nats-io/nats.go"
//	uuid "github.com/satori/go.uuid"
//	"github.com/simonalong/gole/util"
//	"math"
//	"testing"
//	"time"
//)
//
//var postfix = "-bc"
//var streamSub = "guangboStream" + postfix
//var subAll = "guangbo.subject" + postfix + ".*"
//var sub = "guangbo.subject" + postfix + ".key"
//
//func TestSend1(t *testing.T) {
//	nc, _ := nats.Connect("localhost:4222")
//	js, _ := nc.JetStream()
//	nctx := nats.Context(context.Background())
//	info, _ := js.StreamInfo(broadcastStreamName)
//	if nil == info {
//		_, _ = js.AddStream(&nats.StreamConfig{
//			Name:     broadcastStreamName,
//			Subjects: []string{broadcastSubjectAll},
//			Retention:  nats.WorkQueuePolicy,
//			Replicas:   1,
//			Discard:    nats.DiscardOld,
//			Duplicates: 30 * time.Second,
//		}, nctx)
//	}
//
//	tctx, cancel := context.WithTimeout(nctx, 10000*time.Second)
//	deadlineCtx := nats.Context(tctx)
//
//	results := make(chan int64)
//	var totalTime int64
//	var totalMessages int64
//
//	go func() {
//		i := 0
//		for {
//			js.Publish(broadcastSubject, []byte("data " + util.ToString(i)), deadlineCtx)
//			time.Sleep(1 * time.Second)
//			i++
//		}
//	}()
//
//	for {
//		select {
//		case <-deadlineCtx.Done():
//			cancel()
//			log.Infof("sent %d messages with average time of %f", totalMessages, math.Round(float64(totalTime/totalMessages)))
//			js.DeleteStream(broadcastStreamName)
//			return
//		case usec := <-results:
//			totalTime += usec
//			totalMessages++
//		}
//	}
//}
//
//func TestConsumer11(t *testing.T) {
//	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
//	id := uuid.NewV4().String()
//	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
//	js, _ := nc.JetStream()
//	sub, _ := js.QueueSubscribeSync(broadcastSubject, "myqueuegroup", nats.Durable(id), nats.DeliverNew())
//
//	for {
//		msg, err := sub.NextMsgWithContext(ctx)
//		if nil != err {
//			log.Printf("err  sub4 %v", err.Error())
//			time.Sleep(1 * time.Second)
//			continue
//		}
//		log.Printf("[consumer: %s] received msg (%v)", id, string(msg.Data))
//		msg.Ack(nats.Context(ctx))
//	}
//}
//
//func TestConsumer12(t *testing.T) {
//	ctx, _ := context.WithTimeout(context.Background(), 1000*time.Second)
//	id := uuid.NewV4().String()
//	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
//	js, _ := nc.JetStream()
//	sub, _ := js.QueueSubscribeSync(broadcastSubject, "myqueuegroup", nats.Durable(id), nats.DeliverNew())
//
//	for {
//		msg, err := sub.NextMsgWithContext(ctx)
//		if nil != err {
//			log.Printf("err  sub4 %v", err.Error())
//			time.Sleep(1 * time.Second)
//			continue
//		}
//		log.Printf("[consumer: %s] received msg (%v)", id, string(msg.Data))
//		msg.Ack(nats.Context(ctx))
//	}
//}
