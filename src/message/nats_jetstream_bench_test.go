package test

import (
	"context"
	"fmt"
	"github.com/lunny/log"
	"github.com/nats-io/nats.go"
	uuid "github.com/satori/go.uuid"
	"github.com/simonalong/gole/util"
	"testing"
	"time"
)

var benchindex = "-index1"
var benchbroadcastStreamName = "benchbroadcastStreamName" + benchindex
var benchbroadcastSubjectAll = "broadcast.subject" + benchindex + ".*"
var benchbroadcastSubject = "broadcast.subject" + benchindex + ".key"

func Benchmark_Testfsd(b *testing.B) {
	//func TestDemo(t *testing.T) {
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	nctx := nats.Context(context.Background())
	info, _ := js.StreamInfo(benchbroadcastStreamName)
	if nil == info {
		_, _ = js.AddStream(&nats.StreamConfig{
			Name:     benchbroadcastStreamName,
			Subjects: []string{benchbroadcastSubjectAll},
		}, nctx)
	}

	tctx, cancel := context.WithTimeout(nctx, 10000*time.Second)
	defer cancel()
	deadlineCtx := nats.Context(tctx)

	//results := make(chan int64)
	//var totalTime int64
	//var totalMessages int64

	fmt.Println(b.N)
	for i := 0; i < b.N; i++ {
		js.Publish(benchbroadcastSubject, []byte("data "+util.ToString(i)), deadlineCtx)
	}
}

var num = 0

func TestForBench(t *testing.T) {
	ctx2, _ := context.WithTimeout(context.Background(), 1000*time.Second)
	ctx := nats.Context(ctx2)
	id := uuid.NewV4().String()
	nc, _ := nats.Connect("localhost:4222", nats.Name(id))
	js, _ := nc.JetStream()
	sub, _ := js.QueueSubscribeSync(benchbroadcastSubject, "myqueuegroup", nats.Durable(id), nats.DeliverNew())

	for {
		msg, err := sub.NextMsgWithContext(ctx)
		if nil != err {
			log.Printf("err  sub4 %v", err.Error())
			time.Sleep(1 * time.Second)
			continue
		}
		if num%100 == 0 {
			log.Printf("[consumer sub4: %s] received msg (%v)", id, string(msg.Data))
		}
		num++
		msg.Ack(ctx)
	}
}
