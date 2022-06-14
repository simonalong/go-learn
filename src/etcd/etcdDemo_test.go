package test

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/simonalong/gole/util"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	config     clientv3.Config
	etcdClient *clientv3.Client
	err        error
)

var Ctx = context.Background()

func init() {
	// 客户端配置
	config = clientv3.Config{
		Endpoints:   []string{"10.30.30.78:22379"},
		DialTimeout: 5 * time.Second,
		Username:    "root",
		Password:    "ZljIsysc0re123",
	}

	if etcdClient, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ok")
}

func TestEtcd1(t *testing.T) {
	authRsp, _ := etcdClient.AuthStatus(Ctx)
	fmt.Println(authRsp)

	//res, err := etcdClient.Put(Ctx, "key4", "v1")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(res)
	//
	//res1, err := etcdClient.Get(Ctx, "key4")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(string(res1.Kvs[0].Value))
}

func TestKeys(t *testing.T) {
	etcdClient.Put(Ctx, "test:key1", "v2")
	etcdClient.Put(Ctx, "test:key1/k1", "v2")
	etcdClient.Put(Ctx, "test:key1/k3", "v2")
	etcdClient.Put(Ctx, "test:k2", "v2")
	etcdClient.Put(Ctx, "test:k3", "v2")
	etcdClient.Put(Ctx, "test:k4", "v2")

	rsp, _ := etcdClient.Get(Ctx, "test:*")
	fmt.Println(util.ToString(rsp))
}

// 配置包含过期时间的
func TestEtcd2(t *testing.T) {
	// 创建契约
	//lease := clientv3.NewLease(etcdClient)

	// 单位是秒
	//leaseRes, _ := lease.Grant(Ctx, 3)
	etcdClient.Put(Ctx, "test:k2", "v2")

	res, _ := etcdClient.Get(Ctx, "test:k2")

	fmt.Println(util.ToJsonString(res))
	time.Sleep(2 * time.Second)
	// 续约，其实就是又往后延迟了3秒
	//lease.KeepAliveOnce(Ctx, leaseRes.ID)
	res, _ = etcdClient.Get(Ctx, "test:k2")
	fmt.Println(res)

	time.Sleep(2 * time.Second)
	res, _ = etcdClient.Get(Ctx, "test:k2")
	fmt.Println(res)

	time.Sleep(2 * time.Second)
	res, _ = etcdClient.Get(Ctx, "test:k2")
	fmt.Println(res)

	time.Sleep(1 * time.Second)
	res, _ = etcdClient.Get(Ctx, "test:k2")
	fmt.Println(res)
}

// 事务操作
func TestEtcd3(t *testing.T) {
	watchRsp := etcdClient.Watch(Ctx, "test:key1")
	for res := range watchRsp {
		fmt.Println(res.Header)
		fmt.Println(res.Events)
		fmt.Println(res.Canceled)
		fmt.Println(res.Created)
		fmt.Println(res.CompactRevision)
		value := res.Events[0].Kv.Value

		fmt.Println("now", time.Now(), "watchConfig", string(value))
	}

}

func TestEtcdChg(t *testing.T) {
	etcdClient.Put(Ctx, "test:key1", "v2")
}

var data int64 = 1
var num int64 = 120
var count int64 = 12

func TestEtcdReadPress(t *testing.T) {
	var start = time.Now()
	var i int64
	var index int64

	for index = 0; index < count; index++ {
		for i = 0; i < data*num; i++ {
			read()
		}

		time.Sleep(1 * time.Second)
		fmt.Println(index + 1)
	}

	var end = time.Now()
	var pre = (end.UnixMilli() - start.UnixMilli())
	var result = util.ToString((data * num * count * 1000) / pre)
	fmt.Printf("read finish, latency=%s, qps=%s", util.ToString(pre), result)
}

func TestEtcdWritePress(t *testing.T) {
	var start = time.Now()
	var i int64
	for i = 0; i < data*num; i++ {
		read()
	}
	var end = time.Now()
	var pre = (end.UnixMilli() - start.UnixMilli())
	var result = util.ToString((data * num * 1000) / pre)
	fmt.Printf("write finish, qps=%s", result)
}

func read() {
	_, err := etcdClient.Get(Ctx, "/key1")
	if err != nil {
		fmt.Println(err)
		return
	}
}

func write(index int) {
	_, err := etcdClient.Put(Ctx, "key4"+util.ToString(index), "v1"+util.ToString(index))
	if err != nil {
		fmt.Println(err)
		return
	}
}

func TestName(t *testing.T) {

	dataMap := map[string]string{}
	dataMap["a"] = "a"
	dataMap["b"] = "b"

	bytes, err := json.Marshal(12)
	if err != nil {
		fmt.Printf("%v", err.Error())
	}
	fmt.Println(string(bytes))
}
