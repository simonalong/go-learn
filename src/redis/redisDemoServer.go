package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

func main() {

	//fun()/**/

	//pubsub := RedisDb.Subscribe(ctx1, "mychannel1")
	//
	//// Wait for confirmation that subscription is created before publishing anything.
	//_, err := pubsub.Receive(ctx1)
	//if err != nil {
	//	panic(err)
	//}
	//
	//// Go channel which receives messages.
	//ch := pubsub.Channel()

	// Publish a message.
	//err := RedisDb.Publish(ctx1, "mychannel1", "hello").Err()
	//if err != nil {
	//	panic(err)
	//}

	//time.AfterFunc(time.Second, func() {
	//	// When pubsub is closed channel is closed too.
	//	_ = pubsub.Close()
	//})

	// Consume messages.
	//for msg := range ch {
	//	fmt.Println(msg.Channel, msg.Payload)
	//}

	//client:=redis.NewClient(&redis.Options{
	//	Addr:"127.0.0.1:6379",
	//	DB:0,
	//})

}

func fun() {

	// There is no error because go-redis automatically reconnects on error.
	pubsub := RedisDb.Subscribe(ctx1, "mychannel1")

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
