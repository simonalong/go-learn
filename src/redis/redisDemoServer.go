package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

var ctx1 = context.Background()

func main() {

	fun()
	time.Sleep(time.Second * 3100)

	//redisdb := redis.NewClient(&redis.Options{
	//	Addr:     "192.168.137.18:6379",
	//	Password: "", // no password set
	//	DB:       0,  // use default DB
	//})
	////rdb.AddHook()
	//
	//pubsub := redisdb.Subscribe("mychannel1")
	//
	//// Wait for confirmation that subscription is created before publishing anything.
	//_, err := pubsub.Receive()
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Go channel which receives messages.
	//ch := pubsub.Channel()
	//
	//// Publish a message.
	//err = redisdb.Publish("mychannel1", "hello").Err()
	//if err != nil {
	//	panic(err)
	//}
	//
	//time.AfterFunc(time.Second, func() {
	//	// When pubsub is closed channel is closed too.
	//	_ = pubsub.Close()
	//})
	//
	//// Consume messages.
	//for msg := range ch {
	//	fmt.Println(msg.Channel, msg.Payload)
	//}
}

func fun() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.30.30.78:26379",
		Password: "ZljIsysc0re123", // no password set
		//DB:       0,  // use default DB
	})

	// There is no error because go-redis automatically reconnects on error.
	pubsub := rdb.Subscribe(ctx1, "mychannel1")

	// Close the subscription when we are done.
	defer func(pubsub *redis.PubSub) {
		err := pubsub.Close()
		if err != nil {
			panic(err)
		}
	}(pubsub)

	ch := pubsub.Channel()

	for msg := range ch {
		fmt.Println(msg.Channel, msg.Payload)
	}
}
