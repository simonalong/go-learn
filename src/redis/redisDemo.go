package main

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.30.30.78:26379",
		Password: "ZljIsysc0re123", // no password set
		//DB:       0,  // use default DB
	})

	redisMode, err := getRedisMode(rdb)
	if nil != err {

	}
	fmt.Println(redisMode)

	//time.Sleep(time.Second * 3)
	//
	//// There is no error because go-redis automatically reconnects on error.
	//pubsub := rdb.Subscribe(ctx, "mychannel1")
	//
	//// Close the subscription when we are done.
	//defer pubsub.Close()

	// err = rdb.Set(ctx, "key1", "12", 2*time.Second).Err()
	// if err != nil {
	// 	panic(err)
	// }
	// val, err := rdb.Get(ctx, "key1").Result()
	// if err != nil {
	// 	fmt.Println("yichang")
	// }
	// fmt.Println("key1", val)

	// fmt.Println(runtime.GOMAXPROCS)
	// fmt.Println(runtime.NumCPU())
	// fmt.Println(runtime.GOMAXPROCS(2))
	// fmt.Println(runtime.GOMAXPROCS(runtime.NumCPU()))

	//
	//val1, err := rdb.Get(ctx, "key1").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("key-ex", val1)

	//val2, err := rdb.Get(ctx, "key2").Result()
	//if err == redis.Nil {
	//	fmt.Println("key2 does not exist")
	//} else if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("key2", val2)
	//}
	// Output: key value
	// key2 does not exist
	//
	//pipe := rdb.TxPipeline()
	//
	//_ = pipe.HSet(ctx, "hkey", "key1", 1).Err()
	//pipe.Expire("key", time.Hour)
	//_, exer := pipe.Exec()
	//if exer != nil {
	//	fmt.Println(exer)
	//}
	//
	//
	//_ = rdb.HSet(ctx, "hkey", "key2", 3).Err()
	//_ = rdb.HSet(ctx, "hkey", "key3", 4).Err()
	//_ = rdb.HSet(ctx, "hkey", "key4", 2).Err()
	//_ = rdb.HSet(ctx, "hkey", "key4", 6).Err()
	//_ = rdb.HSet(ctx, "hkey", "key4", 6).Err()
	//

	//if err != nil {
	//	panic(err)
	//}

	//val1, err := rdb.HGet(ctx, "hkey", "key1").Result()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("hkey, key1", val1)
	//
	//
	//dataMap, _ := rdb.HGetAll(ctx, "hkey").Result()
	//var count int
	//for _, value := range dataMap {
	//	num, _ := strconv.Atoi(value)
	//	count += num
	//}
	//
	//fmt.Println(count)
}

func getRedisMode(rdb *redis.Client) (string, error) {
	cmd := rdb.Info(ctx, "server")
	dataRs, err := cmd.Result()
	if err != nil {
		return "", err
	}
	fmt.Println(dataRs)
	datas := strings.Split(dataRs, "\n")
	var mode = ""
	for _, data := range datas {
		if strings.Contains(data, "redis_mode") {
			modes := strings.Split(data, ":")
			mode = strings.TrimSpace(modes[1])
		}
	}
	return mode, nil
}
