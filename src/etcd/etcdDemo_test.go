package test

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"testing"
	"time"
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
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}

	if etcdClient, err = clientv3.New(config); err != nil {
		fmt.Println(err)
		return
	}
}

func TestEtcd1(t *testing.T) {
	res, err := etcdClient.Put(Ctx, "test:key1", "v1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

	res1, err := etcdClient.Get(Ctx, "test:key1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(res1.Kvs[0].Value))
}

// 配置包含过期时间的
func TestEtcd2(t *testing.T) {
	// 创建契约
	lease := clientv3.NewLease(etcdClient)

	leaseRes, _ := lease.Grant(Ctx, 3)
	etcdClient.Put(Ctx, "test:k2", "v2", clientv3.WithLease(leaseRes.ID))

	res, _ := etcdClient.Get(Ctx, "test:k2")

	fmt.Println(res)
	time.Sleep(2 * time.Second)
	// 续约，其实就是又往后延迟了3秒
	lease.KeepAliveOnce(Ctx, leaseRes.ID)
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
