package test

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func TestRedisMode(t *testing.T) {
	// 先尝试单机
	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{"10.30.30.47:6379"},
		Password: "ZljIsysc0re123", // no password set
	})

	redisMode, err := getRedisMode(rdb)
	if nil != err {
		fmt.Println("error", err)
		return
	}

	fmt.Println(redisMode)

	// rdb is *redis.Client.

	// // rdb is *redis.ClusterClient.
	// rdb := NewUniversalClient(&redis.UniversalOptions{
	// 	Addrs: []string{":6379", ":6380"},
	// })

	// // rdb is *redis.FailoverClient.
	// rdb := NewUniversalClient(&redis.UniversalOptions{
	// 	Addrs:      []string{":6379"},
	// 	MasterName: "mymaster",
	// })
}

func getRedisMode(rdb redis.UniversalClient) (string, error) {
	cmd := rdb.Info(ctx, "server")
	dataRs, err := cmd.Result()
	if err != nil {
		return "", err
	}
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

// // 获取redis客户端
// func getRedisClient() redis.UniversalClient {
// 	rdb := redis.NewUniversalClient(&redis.UniversalOptions{
// 		Addrs: []string{"10.30.30.47:6379"},
// 		// Addrs:    []string{"10.30.30.78:26379"},
// 		Password: "ZljIsysc0re123", // no password set
// 	})

// 	redisMode, err := getRedisMode(rdb)
// 	if nil != err {
// 		fmt.Println("error", err)
// 		return
// 	}

// 	fmt.Println(redisMode)

// }

// // 获取redis的单机或者主从的客户端
// func getStandaloneRedisClient() redis.UniversalClient {

// }

// // 获取redis的哨兵客户端
// func getSentinelRedisClient() redis.UniversalClient {

// }

// // 获取redis的集群的客户端
// func getClusterRedisClient() redis.UniversalClient {

// }

// 测试单机版
func TestClient1(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:16379",
	})

	ctx := context.Background()

	rdb.Set(ctx, "single-k", "v", time.Hour)

	v := rdb.Get(ctx, "single-k")
	fmt.Println(v.Result())
}

// 测试单机版
func TestClient2(t *testing.T) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.30.30.78:26379",
		Password: "ZljIsysc0re123",
	})

	redisMode, _ := getRedisMode(rdb)
	fmt.Println(redisMode)
}

// 测试主从：go-redis里面没有找到主从的支持姿势，我这边尝试了如下
func TestMasterSlave(t *testing.T) {
	// 主
	rdbWriteRdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6371",
	})

	// 多个从节点
	readWriteRdb := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"shard1": "localhost:6372",
			"shard2": "localhost:6373",
		},
	})

	rdbWriteRdb.Set(context.Background(), "k2", "ddx2", 4*time.Hour)

	v := readWriteRdb.Get(context.Background(), "k2")
	fmt.Println(v.Result())
}

// 测试哨兵time
func TestSentinel(t *testing.T) {
	rdb := redis.NewFailoverClient(&redis.FailoverOptions{
		MasterName:    "mymaster",
		SentinelAddrs: []string{"localhost:26379", "localhost:26380", "localhost:26381"},
	})

	rdb.Set(context.Background(), "k2", "ddx2", 4*time.Hour)

	v := rdb.Get(context.Background(), "k2")
	fmt.Println(v.Result())
}

// 测试集群
func TestCluster(t *testing.T) {
	rdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"localhost:6381", "localhost:6382", "localhost:6383", "localhost:6384", "localhost:6385", "localhost:6386"},
	})

	rdb.Set(context.Background(), "k2", "ddx2", 4*time.Hour)

	v := rdb.Get(context.Background(), "k2")
	fmt.Println(v.Result())

	rdb1 := redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs: []string{":6379"},
	})
}
