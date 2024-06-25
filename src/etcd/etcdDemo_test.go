package test

import (
	"context"
	"fmt"
	mik "go-learn/src/mk"
	"go.etcd.io/etcd/client/v3/concurrency"
	"reflect"
	"testing"
	"time"
	"unsafe"

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
		Endpoints:   []string{"10.30.30.78:52379"},
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
	//authRsp, _ := etcdClient.AuthStatus(Ctx)
	//fmt.Println(authRsp)

	res, err := etcdClient.Put(Ctx, "key4", "v1")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(res)

	res1, err := etcdClient.Get(Ctx, "key4")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(res1.Kvs[0].Value))
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

	//dataMap := map[string]string{}
	//dataMap["a"] = "a"
	//dataMap["b"] = "b"
	//
	//bytes, err := json.Marshal(12)
	//if err != nil {
	//	fmt.Printf("%v", err.Error())
	//}
	//fmt.Println(string(bytes))

	//fmt.Println(*flag.String("a", "b","c"))
}

func TestLock1(t *testing.T) {

	//为锁生成session
	s1, err := concurrency.NewSession(etcdClient, concurrency.WithTTL(5))
	if err != nil {
		t.Log(err.Error())
		return
	}
	defer s1.Close()

	ctx := context.Background()
	locker := concurrency.NewMutex(s1, "test/name")
	fmt.Println("acquiring lock")
	if err := locker.TryLock(ctx); err != nil {
		fmt.Println("尝试失败", err.Error())
		return
	}

	//  请求锁

	fmt.Println("acquired lock")

	time.Sleep(time.Second * 10)

	//time.Sleep(time.Duration(rand.Intn(30))*time.Second)
	locker.Unlock(ctx) //释放锁
	fmt.Println("released lock")
}

func TestLock2(t *testing.T) {

	//为锁生成session
	s1, err := concurrency.NewSession(etcdClient, concurrency.WithTTL(5))
	if err != nil {
		t.Log(err.Error())
		return
	}
	defer s1.Close()
	fmt.Println("acquiring lock")
	ctx := context.Background()
	locker := concurrency.NewMutex(s1, "test/name")
	if err := locker.TryLock(ctx); err != nil {
		fmt.Println("尝试失败", err.Error())
		return
	}

	//  请求锁
	fmt.Println("acquired lock")

	time.Sleep(time.Second * 10)

	//time.Sleep(time.Duration(rand.Intn(30))*time.Second)
	locker.Unlock(ctx) //释放锁
	fmt.Println("released lock")
}

type Demo struct {
	private        string
	youCannotSeeMe int
	Trick          bool
}

func TestPrivate(t *testing.T) {
	//d := Demo{private: "hahaha", youCannotSeeMe: 110, Trick: true}

	//type Header struct {
	//	NotPrivate  string
	//	YouCanSeeMe int
	//}
	//
	//h := *(*Header)(uintptr(unsafe.Pointer(&d)))
	//h.YouCanSeeMe = 32
	//
	//fmt.Printf("%+v", h)
	//fmt.Printf("%+v", d)

	//priV := reflect.ValueOf(d).FieldByName("private")
	//priV = reflect.NewAt(priV.Type(), unsafe.Pointer(priV.UnsafeAddr())).Elem()
	//v := priV.Interface()
	//pvv := v.(string)
	//pvv = "sdfsdf"
}

type TestPointer struct {
	A int
	b int // 私有变量
}

func (T *TestPointer) OouPut() {
	fmt.Println("TestPointer OouPut:", T.A, T.b)
}

func TestPrivate2(t *testing.T) {
	T := TestPointer{A: 1}
	pb := (*int)(unsafe.Pointer(uintptr(unsafe.Pointer(&T)) + 8))
	/*
	   Tmp := uintptr(unsafe.Pointer(&T)) + 8)
	   pb := (*int)(unsafe.Pointer(Tmp)
	   千万不能出现这种用临时变量中转一下的情况。因为GC可能因为优化内存碎片的原因移动了这个对象。只保留了指针的地址是没有意义的。
	*/
	*pb = 2
	T.OouPut() //1 2
}

//// Index 获取首页信息
//func (receiver *Home) Index(req *controllers.HomeRequest, ctx *gin.Context) (*controllers.HomeResponse, error) {
//	value := reflect.ValueOf(ctx).Elem()
//
//	engine := value.FieldByName("engine")
//	// rf can't be read or set.
//	engine = reflect.NewAt(engine.Type(), unsafe.Pointer(engine.UnsafeAddr())).Elem()
//	v := engine.Interface()
//	context := v.(*gin.Engine)
//	help.Dump(context)
//
//	return &controllers.HomeResponse{}, nil
//}

type TestStruct struct {
	mik.TestStruct33
	testField int
}

func (pt TestStruct) Namess() []string {
	return pt.TestStruct33.Namess()
}

func TestPrivate3(t *testing.T) {
	//var s = TestStruct{testField: 100, name: []string{"z", "haha"}}
	//
	//// 获取旧数据
	//ov := GetPrivateFieldValue(reflect.ValueOf(&s), "name")
	//// 100
	//fmt.Println(ov)
	//
	//ovV := ov.([]string)
	//newDatas := append(ovV, "kkk")
	//// 修改为新数据
	//var data []string
	//data = newDatas
	//SetFieldPrivateValue(reflect.ValueOf(&s), "name", reflect.ValueOf(&data))
	//// 21
	//fmt.Println(s)
}

// 获取对象的私有属性
func GetPrivateFieldValue(objPtrValue reflect.Value, fieldName string) interface{} {
	fieldValue := objPtrValue.Elem().FieldByName(fieldName)
	return reflect.NewAt(fieldValue.Type(), unsafe.Pointer(fieldValue.UnsafeAddr())).Elem().Interface()
}

// 给对象的属性设置值
func SetFieldPrivateValue(objPtrValue reflect.Value, fieldName string, fieldNewValue reflect.Value) {
	fieldValue := objPtrValue.Elem().FieldByName(fieldName)
	fieldValue = reflect.NewAt(fieldValue.Type(), unsafe.Pointer(fieldValue.UnsafeAddr())).Elem()
	fieldValue.Set(fieldNewValue.Elem())
}
