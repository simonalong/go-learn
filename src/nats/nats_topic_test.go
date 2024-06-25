package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/simonalong/gole/util"
	"testing"
	"time"
)

// 测试topic的接收：subject.*.test
func TestTopicSub(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	_, err := nc.Subscribe("subject.*.test", func(m *nats.Msg) {
		fmt.Printf("%s\n", string(m.Data))
	})
	if err != nil {
		return
	}

	time.Sleep(100000 * time.Second)
	nc.Close()
}

// 测试topic的接收：subject.*.test
func TestTopicSub_2(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	nc.Subscribe("subject.*.demo1", func(m *nats.Msg) {
		fmt.Printf("%s\n", string(m.Data))
	})

	time.Sleep(100000 * time.Second)
	nc.Close()
}

func TestP1(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	pub(nc, "p1", "subject.demo1")
}
func TestP2(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	pub(nc, "p2", "subject.demo2")
}
func TestP3(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	pub(nc, "p3", "subject.demo1.test")
}

func TestP4(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	pub(nc, "p4", "subject")
}

func TestP5(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	pub(nc, "p5", "subject123")
}

// 测试topic的接收：subject.*.test
func TestTopicSub2(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	_, err := nc.Subscribe("subject.>", func(m *nats.Msg) {
		fmt.Printf("%s\n", string(m.Data))
	})
	if err != nil {
		return
	}

	time.Sleep(100000 * time.Second)
	nc.Close()
}

// 测试topic的接收：subject.*.test
func TestTopicSub3(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	_, err := nc.Subscribe("subject>", func(m *nats.Msg) {
		fmt.Printf("%s\n", string(m.Data))
	})
	if err != nil {
		return
	}

	time.Sleep(100000 * time.Second)
	nc.Close()
}

// 测试topic的发布：发布subject.demo1.test
func TestTopicPub1(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	pub(nc, "demo1", "subject.demo1.test")
}

// 测试topic的发布：发布subject.demo1.test
func TestTopicPub1_short(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	pub(nc, "demo1", "subject.demo1")
}

// 测试topic的发布：发布subject.demo2.test
func TestTopicPub2(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	pub(nc, "demo2", "subject.demo2.test")
}

func TestTopicPub3(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	pub(nc, "demo3", "subject.demo3.test")
}

// 测试topic的发布：发布subject.demo3.test1
func TestTopicPubNone(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	pub(nc, "demo_none", "subject..test")
}

// 测试topic的发布：发布subject.demo2.test
func TestTopicPub4(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	pub(nc, "root", "subject")
}

// 测试topic的发布：发布subject.demo2.test
func TestTopicPub5(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	for i := 0; i < 100; i++ {
		nc.Publish("subject1", []byte("root："+util.ToString(i)+" Hello World"))
		time.Sleep(500 * time.Millisecond)
	}
	nc.Close()
}

func pub(nc *nats.Conn, pre, subject string) {
	for i := 0; i < 10000; i++ {
		nc.Publish(subject, []byte(pre+"  "+util.ToString(i)+" Hello World"))
		time.Sleep(500 * time.Millisecond)
	}
}
