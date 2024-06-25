package nats

import (
	"fmt"
	"github.com/nsqio/go-nsq"
	"github.com/simonalong/gole/util"
	"log"
	"testing"
)

type myMessageHandler struct{}

// HandleMessage implements the Handler interface.
func (h *myMessageHandler) HandleMessage(m *nsq.Message) error {
	if len(m.Body) == 0 {
		// Returning nil will automatically send a FIN command to NSQ to mark the message as processed.
		// In this case, a message with an empty body is simply ignored/discarded.
		return nil
	}

	// do whatever actual message processing is desired
	//err := processMessage(m.Body)
	fmt.Println(m.Body)

	// Returning a non-nil error will automatically send a REQ command to NSQ to re-queue the message.
	return nil
}

var totalNsqNum = 100
var totalNsqSize = 10000

func TestPub(t *testing.T) {
	cfg := nsq.NewConfig()
	// 连接 nsqd 的 tcp 连接
	//producer, err := nsq.NewProducer("127.0.0.1:4150", cfg)
	producer, err := nsq.NewProducer("127.0.0.1:24150", cfg)
	if err != nil {
		log.Fatal(err)
	}

	doneChan := make(chan *nsq.ProducerTransaction, totalNsqNum*totalNsqSize)

	// 发布消息
	for i := 0; i < totalNsqNum*totalNsqSize; i++ {
		// 同步：
		//js.Publish(broadcastPressSubject, []byte("message=="+util.ToString(i)), nats.Context(ctx))
		if err := producer.PublishAsync("test-topic", []byte(fmt.Sprintf("test %d", i)), doneChan, "test"); err != nil {
			log.Fatal("publish error: " + err.Error())
		}
		//time.Sleep(1 * time.Second)
	}

	for i := 0; i < totalNsqNum*totalNsqSize; i++ {
		trans := <-doneChan
		if trans.Error != nil {
			t.Fatalf(trans.Error.Error())
		}
		if trans.Args[0].(string) != "test" {
			t.Fatalf(`proxied arg "%s" != "test"`, trans.Args[0].(string))
		}
	}

	log.Printf("send finish")
}

var pressNsqCount1 = 0
var pressNsqCount2 = 0

func TestSub1(t *testing.T) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test-topic", "channel0", cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 处理信息
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		pressNsqCount1++
		if pressNsqCount1%((totalNsqNum*totalNsqSize)/100) == 0 {
			log.Printf("[consumer] received msg (%v) ratio: %s", string(message.Body), util.ToString((pressNsqCount1*100)/(totalNsqNum*totalNsqSize)))
		}
		//log.Printf("[consumer] received msg (%v) ratio: %s", string(message.Body), util.ToString((pressNsqCount1*100)/(totalNsqNum*totalNsqSize)))
		return nil
	}))

	// 连接 nsqd 的 tcp 连接
	//if err := consumer.ConnectToNSQD("127.0.0.1:4150"); err != nil {
	if err := consumer.ConnectToNSQD("127.0.0.1:24150"); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
}

func TestSub2(t *testing.T) {
	cfg := nsq.NewConfig()
	consumer, err := nsq.NewConsumer("test-topic", "channel1", cfg)
	if err != nil {
		log.Fatal(err)
	}

	// 处理信息
	consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		pressNsqCount2++
		if pressNsqCount2%((totalNsqNum*totalNsqSize)/100) == 0 {
			log.Printf("[consumer] received msg (%v) ratio: %s", string(message.Body), util.ToString((pressNsqCount2*100)/(totalNsqNum*totalNsqSize)))
		}
		return nil
	}))

	// 连接 nsqd 的 tcp 连接
	//if err := consumer.ConnectToNSQD("127.0.0.1:4150"); err != nil {
	if err := consumer.ConnectToNSQD("127.0.0.1:24150"); err != nil {
		log.Fatal(err)
	}
	<-consumer.StopChan
}
