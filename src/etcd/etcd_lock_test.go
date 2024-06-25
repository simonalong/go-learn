package test

//
//import (
//	"context"
//	"fmt"
//	clientv3 "go.etcd.io/etcd/client/v3"
//	"testing"
//	"time"
//)
//
//package main
//
//import (
//"context"
//"fmt"
//"time"
//
//"go.etcd.io/etcd/clientv3"
//)
//
//var (
//	client *clientv3.Client
//	cfg    clientv3.Config
//	err    error
//	lease clientv3.Lease
//	ctx context.Context
//	cancelFunc context.CancelFunc
//	leaseId clientv3.LeaseID
//	leaseGrantResponse *clientv3.LeaseGrantResponse
//	leaseKeepAliveChan <-chan *clientv3.LeaseKeepAliveResponse
//	leaseKeepAliveResponse *clientv3.LeaseKeepAliveResponse
//	txn clientv3.Txn
//	txnResponse *clientv3.TxnResponse
//	kv clientv3.KV
//)
//
//func TestLock(t *testing.T) {
//
//	cfg = clientv3.Config{
//		Endpoints:   []string{"youwebsite:2379"},
//		DialTimeout: time.Second * 5,
//	}
//	if client, err = clientv3.New(cfg); err != nil {
//		fmt.Println(err)
//		return
//	}
//
//	lease = clientv3.NewLease(client)
//	if leaseGrantResponse,err = lease.Grant(context.TODO(),5);err!=nil{
//		fmt.Println(err)
//		return
//	}
//	leaseId = leaseGrantResponse.ID
//
//	//租约自动过期，立刻过期。//cancelfunc 取消续租，而revoke 则是立即过期
//	ctx,cancelFunc = context.WithCancel(context.TODO())
//	defer cancelFunc()
//	defer lease.Revoke(context.TODO(),leaseId)
//
//	if leaseKeepAliveChan,err = lease.KeepAlive(ctx,leaseId);err!=nil{
//		fmt.Println(err)
//		return
//	}
//
//	//启动续租协程，每秒续租一次
//	go func() {
//		for {
//			select {
//			case leaseKeepAliveResponse = <-leaseKeepAliveChan:
//				if leaseKeepAliveResponse != nil{
//					fmt.Println("续租成功,leaseID :",leaseKeepAliveResponse.ID)
//				}else {
//					fmt.Println("续租失败")
//				}
//
//			}
//			time.Sleep(time.Second*1)
//		}
//	}()
//
//	//锁逻辑。
//	kv = clientv3.NewKV(client)
//	txn = kv.Txn(context.TODO())
//
//	txn.If(clientv3.Compare(clientv3.CreateRevision("/dev/lock"),"=",0)).Then(
//		clientv3.OpPut("/dev/lock","占用",clientv3.WithLease(leaseId))).Else(
//		clientv3.OpGet("/dev/lock"))
//	if txnResponse,err = txn.Commit();err!=nil{
//		fmt.Println(err)
//		return
//	}
//	if txnResponse.Succeeded {
//		fmt.Println("抢到锁了")
//	}else {
//		fmt.Println("没抢到锁",txnResponse.Responses[0].GetResponseRange().Kvs[0].Value)
//	}
//	time.Sleep(time.Second * 10 )
//}
