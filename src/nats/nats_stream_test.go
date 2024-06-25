package nats

import (
	"context"
	"fmt"
	baseTime "github.com/isyscore/isc-gobase/time"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/simonalong/gole/util"
	"testing"
	"time"
)

var streamName = "stream_test1"

func TestSteam(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream()

	// 新增
	js.AddStream(&nats.StreamConfig{
		Name: "stream_test1",
	})

	// 更新
	js.UpdateStream(&nats.StreamConfig{
		Name: "stream_test1_up",
	})

	// 删除
	js.DeleteStream("stream_test1_up")
}

// 流式：订阅
func TestStreamSubscribe(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream()
	defer nc.Close()

	info, _ := js.StreamInfo("stream_demo3")
	if nil == info {
		js.AddStream(&nats.StreamConfig{
			Name:     "stream_demo3",
			Subjects: []string{"topic.demo3"},
			NoAck:    false,
		})
	}

	//js.Publish("topic.demo2", []byte("你好1aaa"))

	sub, _ := js.Subscribe("topic.demo3", func(msg *nats.Msg) {
		//js.Subscribe("topic.demo", func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
		msg.Ack()
	}, nats.Durable("durable_test"))

	time.Sleep(30 * time.Second)
	sub.Unsubscribe()
}

func TestStreamSubscribe1(t *testing.T) {
	subjectName := "topic.test4"
	js := getJs(subjectName)

	js.Subscribe(subjectName, func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
		msg.Ack()
		//}, nats.Durable("durable_test1"))
	})

	time.Sleep(30999 * time.Second)
}

// 流式（持久消费者）：订阅
func TestStreamSubscribe2(t *testing.T) {
	subjectName := "topic.test4"
	js := getJs(subjectName)

	js.Subscribe(subjectName, func(msg *nats.Msg) {
		fmt.Println(baseTime.TimeToStringYmdHms(time.Now()), string(msg.Data))
	}, nats.Durable("durable_test4"), nats.ManualAck())
	//}, nats.Durable("durable_test4"))
	time.Sleep(30999 * time.Second)
}

// 流式（持久消费者）：订阅，顺序消息
func TestStreamSubscribeOrder(t *testing.T) {
	subjectName := "topic.test4"
	js := getJs(subjectName)

	js.Subscribe(subjectName, func(msg *nats.Msg) {
		fmt.Println(baseTime.TimeToStringYmdHms(time.Now()), string(msg.Data))
	}, nats.Durable("durable_test4"), nats.OrderedConsumer())
	time.Sleep(30999 * time.Second)
}

func TestStreamSubscribeMulti1(t *testing.T) {
	subjectName := "topic.test4"
	js := getJs(subjectName)

	js.Subscribe(subjectName, func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
		//time.Sleep(1 * time.Second)
	}, nats.Durable("durable_test4"))
	//}, nats.Durable("durable_test4"), nats.OrderedConsumer())

	time.Sleep(30999 * time.Second)
}

func TestStreamSubscribeMulti2(t *testing.T) {
	subjectName := "topic.test4"
	js := getJs(subjectName)

	js.Subscribe(subjectName, func(msg *nats.Msg) {
		fmt.Println(string(msg.Data))
		//time.Sleep(1 * time.Second)
	}, nats.Durable("durable_test4"))
	//}, nats.Durable("durable_test4"), nats.OrderedConsumer())

	time.Sleep(30999 * time.Second)
}

// 持久订阅（消息回放）：
func TestStreamSubscribeReply(t *testing.T) {
	streamName := "stream_test4"
	subjectName := "topic.test4"

	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Minute)
	defer cancel()

	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		fmt.Println(err)
	}
	s, err := js.CreateStream(ctx, jetstream.StreamConfig{
		Name:     streamName,
		Subjects: []string{subjectName},
	})
	if err != nil {
		fmt.Println(err)
	}

	// 过去12小时的数据
	startTime := time.Now().Add(-12 * time.Hour)
	// 这里也可以配置往前回拨的时间
	cons, err := s.OrderedConsumer(ctx, jetstream.OrderedConsumerConfig{
		MaxResetAttempts: 5,
		// 开始回放的的时间
		OptStartTime: &startTime,
		// 开始回放的序列
		//OptStartSeq: 3178293,
	})
	if err != nil {
		fmt.Println(err)
	}

	for {
		msgs, err := cons.Fetch(100)
		if err != nil {
			fmt.Println(err)
		}
		for msg := range msgs.Messages() {
			fmt.Println(string(msg.Data()))
			msg.Ack()
		}
		if msgs.Error() != nil {
			fmt.Println("Error fetching messages: ", err)
		}
	}
}

func getJs2(subjectName string) (*nats.Conn, nats.JetStreamContext) {
	streamName := "stream_test4"
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream()

	info, _ := js.StreamInfo(streamName)
	if nil == info {
		js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{subjectName},
			NoAck:    false,
		})
	}
	return nc, js
}

func getJs(subjectName string) nats.JetStreamContext {
	streamName := "stream_test4"
	nc, _ := nats.Connect(nats.DefaultURL)
	js, _ := nc.JetStream()

	info, _ := js.StreamInfo(streamName)
	if nil == info {
		js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{subjectName},
			NoAck:    false,
		})
	}
	return js
}

// 流式：发布
func TestStreamPublish(t *testing.T) {
	subjectName := "topic.test4"
	js := getJs(subjectName)
	//js.Publish(subjectName, []byte("你好1241"))
	//js.Publish(subjectName, []byte("你好2241"))
	//js.Publish(subjectName, []byte("你好3241"))
	for i := 0; i < 10000; i++ {
		js.Publish(subjectName, []byte("hello "+util.ToString(i)))
		time.Sleep(500 * time.Millisecond)
	}
}

//type StreamConfig struct {
//	// 流的名字
//	Name string `json:"name"`
//	// 描述
//	Description string `json:"description,omitempty"`
//	// 监听的主题，支持多个
//	Subjects []string `json:"subjects,omitempty"`
//	// 保留策略
//	Retention RetentionPolicy `json:"retention"`
//	// 最大消费者
//	MaxConsumers int `json:"max_consumers"`
//	// 限制流中最多存储消息的条数
//	MaxMsgs int64 `json:"max_msgs"`
//	// 限制流的总字节大小
//	MaxBytes int64 `json:"max_bytes"`
//	// 限制单个消息的最大允许字节大小：这个字段对于控制消息的大小非常有用，以确保消息不会消耗过多的内存或存储资源，并防止过大的消息影响系统性能。
//	MaxMsgSize int32 `json:"max_msg_size,omitempty"`
//	// 限制流的消息的最大存活时间
//	MaxAge time.Duration `json:"max_age"`
//	// 限制每个主题的最大消息个数
//	MaxMsgsPerSubject int64 `json:"max_msgs_per_subject"`
//	// 丢弃策略
//	Discard DiscardPolicy `json:"discard"`
//	// 持久化存储
//	Storage StorageType `json:"storage"`
//
//	// 副本数：建议使用1、3、5；1性能很高，但是容错较低；3性能和风险持平，可以容忍一台机器数据丢失；5牺牲性能换取容错性，可以容忍两台机器数据丢失
//	Replicas int `json:"num_replicas"`
//	// 是否回应，这个是流级别配置，真正起作用在consumerConfig里面的NoAck会覆盖这个；选择是否回应，true-不回应则性能高，但是存在数据丢失风险
//	NoAck bool `json:"no_ack,omitempty"`
//	// 摸板名称，可以事先创建一个模板StreamTemplate，在创建流的时候直接通过templateName来使用对应的模板
//	Template string `json:"template_owner,omitempty"`
//	// 去重功能；在一段时间内保证消息的唯一性
//	Duplicates time.Duration `json:"duplicate_window,omitempty"`
//	// 用于指导流在集群JetStream中的放置
//	Placement *Placement `json:"placement,omitempty"`
//	// 镜像流，可以用来复制已经创建的流
//	Mirror *StreamSource `json:"mirror,omitempty"`
//	// 允许多个流镜像
//	Sources []*StreamSource `json:"sources,omitempty"`
//	// 流是否为密封流，密封流则为不允许修改的流
//	Sealed      bool `json:"sealed,omitempty"`
//	DenyDelete  bool `json:"deny_delete,omitempty"`
//	DenyPurge   bool `json:"deny_purge,omitempty"`
//	AllowRollup bool `json:"allow_rollup_hdrs,omitempty"`
//}
//
//// StreamConfig will determine the properties for a stream.
//// There are sensible defaults for most. If no subjects are
//// given the name will be used as the only subject.
//type StreamConfig struct {
//	Name        string          `json:"name"`
//	Description string          `json:"description,omitempty"`
//	Subjects    []string        `json:"subjects,omitempty"`
//	Retention   RetentionPolicy `json:"retention"`
//
//	MaxConsumers      int           `json:"max_consumers"`
//	MaxMsgs           int64         `json:"max_msgs"`
//	MaxBytes          int64         `json:"max_bytes"`
//	MaxAge            time.Duration `json:"max_age"`
//	MaxMsgsPerSubject int64         `json:"max_msgs_per_subject"`
//	MaxMsgSize        int32         `json:"max_msg_size,omitempty"`
//
//	Discard              DiscardPolicy `json:"discard"`
//	DiscardNewPerSubject bool          `json:"discard_new_per_subject,omitempty"`
//
//	Storage    StorageType     `json:"storage"`
//	Replicas   int             `json:"num_replicas"`
//	NoAck      bool            `json:"no_ack,omitempty"`
//	Duplicates time.Duration   `json:"duplicate_window,omitempty"`
//	Placement  *Placement      `json:"placement,omitempty"`
//	Mirror     *StreamSource   `json:"mirror,omitempty"`
//	Sources    []*StreamSource `json:"sources,omitempty"`
//
//	Sealed           bool                    `json:"sealed,omitempty"`
//	DenyDelete       bool                    `json:"deny_delete,omitempty"`
//	DenyPurge        bool                    `json:"deny_purge,omitempty"`
//	AllowRollup      bool                    `json:"allow_rollup_hdrs,omitempty"`
//	Compression      StoreCompression        `json:"compression"`
//	FirstSeq         uint64                  `json:"first_seq,omitempty"`
//	SubjectTransform *SubjectTransformConfig `json:"subject_transform,omitempty"`
//	RePublish        *RePublish              `json:"republish,omitempty"`
//	AllowDirect      bool                    `json:"allow_direct"`
//	MirrorDirect     bool                    `json:"mirror_direct"`
//	ConsumerLimits   StreamConsumerLimits    `json:"consumer_limits,omitempty"`
//	Metadata         map[string]string       `json:"metadata,omitempty"`
//	Template         string                  `json:"template_owner,omitempty"`
//}
//
//const (
//	//LimitsPolicy（默认）意味着消息将被保留，直到达到任何给定的限制。
//	//这可以是MaxMsgs、MaxBytes或MaxAge中的一个。
//	LimitsPolicy RetentionPolicy = iota
//	//InterestPolicy指定，当所有已知的可观察器都已确认消息时，可以删除该消息。
//	InterestPolicy
//	//WorkQueuePolicy指定当第一个工作者或订阅者确认消息时，可以将其删除。
//	WorkQueuePolicy
//)
//
//const (
//	//DiscardOld将删除较旧的消息以返回限制。这是
//	//默认值。
//	DiscardOld DiscardPolicy = iota
//	//DiscardNew将无法存储新邮件。
//	DiscardNew
//)
//
//const (
//	//FileStorage指定磁盘上的存储。这是默认设置。
//	FileStorage StorageType = iota
//	//MemoryStorage仅在内存中指定。
//	MemoryStorage
//)
//
//type StreamSource struct {
//	// 待复制的源stream的名字
//	Name string `json:"name"`
//	// 开始复制消息的序列号
//	OptStartSeq uint64 `json:"opt_start_seq,omitempty"`
//	// 开始复制消息的时间点
//	OptStartTime *time.Time `json:"opt_start_time,omitempty"`
//	// 复制的主题名字
//	FilterSubject string `json:"filter_subject,omitempty"`
//	// 外部集群流的源配置
//	External *ExternalStream `json:"external,omitempty"`
//}
//
//type ExternalStream struct {
//	// 通常是url，比如：https://external-nats.example.com:4443
//	APIPrefix string `json:"api"`
//	// 过滤的主题
//	DeliverPrefix string `json:"deliver"`
//}
