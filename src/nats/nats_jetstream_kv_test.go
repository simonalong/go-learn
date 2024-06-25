package nats

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"testing"
	"time"
)

func TestKv1(t *testing.T) {
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	kv := GetBucket(js, "bucket", 3*time.Second)

	// put
	kv.Put("key.k1", []byte("value12"))

	// delete
	kv.Delete("key.k1")
	// purge，跟delete一样，但是会清理的更彻底
	kv.Purge("key.k1")

	// 更新
	kv.Update("key.k1", []byte("value_chg"), 0)

	// get
	data, _ := kv.Get("key.k1")
	if nil != data {
		fmt.Println(string(data.Value()))
	} else {
		fmt.Println("null")
	}

	// keys
	//keys1, _ := kv.Keys()

	// watchs
	keyWatcher, _ := kv.WatchAll()
	for {
		select {
		case kvs := <-keyWatcher.Updates():
			if nil != kvs {
				fmt.Println(string(kvs.Key()) + " = " + string(kvs.Value()))
			}
		}
	}
}

func TestKv2(t *testing.T) {
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	kv := GetBucket(js, "bucket", 3*time.Second)

	// put
	kv.Put("key.k2", []byte("value12"))

	// delete
	kv.Delete("key.k2")
	// purge，跟delete一样，但是会清理的更彻底
	kv.Purge("key.k2")

	data, _ := kv.Get("key.k2")
	if nil != data {
		fmt.Println(string(data.Value()))
	} else {
		fmt.Println("null")
	}

	kv.Create("key.k3", []byte("nihao"))

	data, _ = kv.Get("key.k3")
	if nil != data {
		fmt.Println(string(data.Value()))
	} else {
		fmt.Println("null")
	}
}

func TestKv3(t *testing.T) {
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	kv := GetBucket(js, "bucket", 3*time.Second)

	// put
	kv.Put("key.k1", []byte("value12"))
	kv.Put("key.k2", []byte("value12"))

	keys, _ := kv.Keys()
	// [key.k1 key.k2]
	fmt.Println(keys)
}

func TestKv4(t *testing.T) {
	nc, _ := nats.Connect("localhost:4222")
	js, _ := nc.JetStream()
	kv := GetBucket(js, "bucket", 3*time.Second)

	// put
	kv.Put("key.k1", []byte("value12"))
	kv.Put("key.k2", []byte("value12"))

	keyWatcher, _ := kv.WatchAll()
	for {
		select {
		case kvs := <-keyWatcher.Updates():
			if nil != kvs {
				fmt.Println(string(kvs.Key()) + " = " + string(kvs.Value()))
			}
		}
	}
}

func GetBucket(js nats.JetStreamContext, name string, ttl time.Duration) nats.KeyValue {
	if kv, _ := js.KeyValue(name); nil != kv {
		return kv
	}
	kv, _ := js.CreateKeyValue(&nats.KeyValueConfig{
		// 桶名字
		Bucket: name,
		// 保存key的实效性
		TTL: ttl,
	})
	return kv
}
