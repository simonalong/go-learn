package test

import (
	"context"
	"encoding/base64"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/simonalong/tools/util"
	"math/rand"
	"testing"
	"time"
)

var ctx1 = context.Background()

var RedisDb = redis.NewClient(&redis.Options{
	Addr:     "10.30.30.78:26379",
	Password: "ZljIsysc0re123",
	DB:       0,
	PoolSize: 20,
})

func TestSub(t *testing.T) {
	pubsub := RedisDb.Subscribe(ctx1, "channel")
	defer pubsub.Close()

	go watch(pubsub)

	pug()
	time.Sleep(12000 * time.Second)
}

func pug() {
	for i := 0; i < 20; i++ {
		time.Sleep(time.Second)
		rand.Seed(time.Now().Unix())
		num := rand.Intn(1000)
		n, _ := RedisDb.Publish(ctx1, "channel", num).Result()
		fmt.Println(n)
	}
}

func watch(pubsub *redis.PubSub) {
	for msg := range pubsub.Channel() {
		fmt.Println(msg.Payload)
	}
}

type ConfigCenterMessage struct {

	// 跟踪的id
	TraceId string

	// 链路跟踪id
	RpcId string

	Profile   string
	AppId     int64
	Group     string
	VersionId int64
	Key       string
	Value     string
	KeyDesc   string
}

var Ctx = context.Background()

func TestPpub(t *testing.T) {
	var configCenterMessage = ConfigCenterMessage{}
	configCenterMessage.Profile = "sdf"
	configCenterMessage.AppId = 12312

	fmt.Println("发送配置变更")
	err := RedisDb.Publish(Ctx, "/redis/isyscore/os/config", Base64Encode([]byte(util.ToJsonString(configCenterMessage)))).Err()
	if err != nil {
		panic(err)
	}
}

func TestOpenRedisSubscribe(t *testing.T) {
	pubsub := RedisDb.Subscribe(Ctx, "/redis/isyscore/os/config")
	defer func(pubsub *redis.PubSub) {
		err := pubsub.Close()
		if err != nil {

		}
	}(pubsub)

	go watch(pubsub)

	time.Sleep(10000 * time.Second)
}

func watch2(pubsub *redis.PubSub) {
	for msg := range pubsub.Channel() {
		data, _ := Base64Decode(msg.Payload)
		message := ConfigCenterMessage{}
		err := util.StrToObject(string(data), &message)
		if err != nil {
			fmt.Println("receive message config chang err", err.Error())
			return
		}

		fmt.Println("收到配置变更")
	}
}

func Base64Encode(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

func Base64Decode(dst string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(dst)
}
