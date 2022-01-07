package test

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"
	"testing"
	"time"
)

//func tree() {
//  ctx1 := context.Background()
//  ctx2, _ := context.WithCancel(ctx1)
//  ctx3, _ := context.WithTimeout(ctx2, time.Second * 5)
//  ctx4, _ := context.WithTimeout(ctx3, time.Second * 3)
//  ctx5, _ := context.WithTimeout(ctx3, time.Second * 6)
//  ctx6 := context.WithValue(ctx5, "userID", 12)
//}

func TestName1(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}

func TestName2(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go handle(ctx, 1500*time.Millisecond)
	select {
	case <-ctx.Done():
		fmt.Println("main", ctx.Err())
	}
}

func TestBackground(t *testing.T) {
	ctx1 := context.Background()
	//ctx2, _ := context.WithCancel(ctx1)
	//ctx3, _ := context.WithTimeout(ctx2, time.Second*5)
	//ctx4, _ := context.WithTimeout(ctx3, time.Second*3)
	//ctx5, _ := context.WithTimeout(ctx3, time.Second*6)
	//ctx6 := context.WithValue(ctx5, "userID", 12)
	//ctx6 := context.WithValue(ctx5, "userID", 12)
	ctx1.Done()
}

var logg *log.Logger

func someHandler() {
	ctx, cancel := context.WithCancel(context.Background())
	go doStuff(ctx)

	//10秒后取消doStuff
	time.Sleep(2 * time.Second)
	cancel()
}

//每1秒work一下，同时会判断ctx是否被取消了，如果是就退出
func doStuff(ctx context.Context) {
	for {
		time.Sleep(1 * time.Second)
		select {
		case <-ctx.Done():
			logg.Printf("done")
			return
		default:
			logg.Printf("work")
		}
	}
}

func TestContext1(t *testing.T) {
	logg = log.New(os.Stdout, "", log.Ltime)
	someHandler()
	logg.Printf("down")
}

func TestContext2(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	go fun1(ctx, cancel)
	go fun2(ctx, cancel)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)
	fmt.Println("test finish")
}

func fun1(ctx context.Context, cancel context.CancelFunc) {
	select {
	case <-ctx.Done():
		fmt.Println("fun1 out finish")
	default:
		fmt.Println("fun1 run")
		time.Sleep(3 * time.Second)
	}

	fmt.Println("fun1 finish")
}

func fun2(ctx context.Context, cancel context.CancelFunc) {
	select {
	case <-ctx.Done():
		fmt.Println("fun2 out finish")
	default:
		fmt.Println("fun2 run")
		time.Sleep(4 * time.Second)
	}
	fmt.Println("fun2 finish")
}

// 今天向后延迟固定时间之后自动调用cancel接口
func TestContext3(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go fun1(ctx, cancel)
	go fun2(ctx, cancel)
	time.Sleep(5 * time.Second)
}

// 具体时间之后自动执行cancel接口
//func TestContext4(t *testing.T) {
//	ctx := context.Background()
//	ctx, cancel := context.WithDeadline(ctx, time.Date(2022, 1, 7, 13, 56, 00, 0, time.Local))
//	go fun1(ctx, cancel)
//	go fun2(ctx, cancel)
//	go fun3(ctx, cancel)
//	time.Sleep(500 * time.Second)
//}

func TestContext5(t *testing.T) {
	ctx := context.Background()
	childCtx := context.WithValue(ctx, "key1", 12)
	go fun4(childCtx)
	time.Sleep(5 * time.Second)
}

func fun4(ctx context.Context) {
	fmt.Println(ctx.Value("key1"))
}

// 任何一个执行完毕，则返回
// 全部执行完毕，则返回
// 任何一个执行失败，则返回失败
// 全部执行失败，则返回失败

// 任何一个执行完毕，则返回
func TestContext6(t *testing.T) {
	ctx := context.Background()
	//childCtx, cancel := context.WithCancel(ctx)
	go fun5(ctx)
	go fun6(ctx)
	go fun7(ctx)
	time.Sleep(7 * time.Second)
	fmt.Println("test finish")
	time.Sleep(10 * time.Second)
}

func fun5(ctx context.Context) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("fun5 task finish")
		}
	}()

	time.Sleep(3 * time.Second)
	ctx.Done()
	fmt.Println("fun5 finish")
}

func fun6(ctx context.Context) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("fun6 task finish")
		}
	}()

	time.Sleep(2 * time.Second)
	ctx.Done()
	fmt.Println("fun6 finish")
}

func fun7(ctx context.Context) {
	go func() {
		select {
		case <-ctx.Done():
			fmt.Println("fun7 task finish")
		}
	}()
	time.Sleep(4 * time.Second)
	ctx.Done()
	fmt.Println("fun7 finish")
}

// 模拟一个最小执行时间的阻塞函数
func inc(a int) int {
	res := a + 1                 // 虽然我只做了一次简单的 +1 的运算,
	time.Sleep(10 * time.Second) // 但是由于我的机器指令集中没有这条指令,
	// 所以在我执行了 1000000000 条机器指令, 续了 1s 之后, 我才终于得到结果。B)
	return res
}

// 向外部提供的阻塞接口
// 计算 a + b, 注意 a, b 均不能为负
// 如果计算被中断, 则返回 -1
func Add(ctx context.Context, a, b int) int {
	res := 0
	for i := 0; i < a; i++ {
		res = inc(res)
		select {
		case <-ctx.Done():
			return -1
		default:
		}
	}
	for i := 0; i < b; i++ {
		res = inc(res)
		select {
		case <-ctx.Done():
			return -1
		default:
		}
	}
	return res
}

func TestContext7(t *testing.T) {
	{
		// 使用开放的 API 计算 a+b
		a := 1
		b := 2
		timeout := 1 * time.Second
		ctx, _ := context.WithTimeout(context.Background(), timeout)
		// 相当于整个函数这边只等1秒，但是如果任务执行耗费了10秒，也是会等到该函数执行完再返回，但是返回过来是已经结束的数据
		res := Add(ctx, 1, 2)
		fmt.Printf("Compute: %d+%d, result: %d\n", a, b, res)
	}
	//{
	//	// 手动取消
	//	a := 1
	//	b := 2
	//	ctx, cancel := context.WithCancel(context.Background())
	//	go func() {
	//		time.Sleep(2 * time.Second)
	//		cancel() // 在调用处主动取消
	//	}()
	//	res := Add(ctx, 1, 2)
	//	fmt.Printf("Compute: %d+%d, result: %d\n", a, b, res)
	//}
}

// 数据接收服务主协程同子协程同步变量
var wg sync.WaitGroup

func run(i int) {
	fmt.Println("start 任务ID：", i)
	time.Sleep(time.Second * 1)
	wg.Done() // 每个goroutine运行完毕后就释放等待组的计数器
}

func TestContext8(t *testing.T) {
	countThread := 8    //runtime.NumCPU()
	wg.Add(countThread) // 需要开启的goroutine等待组的计数器
	for i := 0; i < countThread; i++ {
		go run(i)
	}

	//等待所有的任务都释放
	wg.Wait()
	fmt.Println("任务全部结束,退出")
}

func run2(stop chan bool) {
	for {
		select {
		case <-stop:
			fmt.Println("任务1结束退出")
			return
		default:
			fmt.Println("任务1正在运行中")
			time.Sleep(time.Second * 2)
		}
	}
}

// 相当于一个定时任务了
func TestContext9(t *testing.T) {
	stop := make(chan bool)
	go run2(stop) // 开启goroutine

	// 运行一段时间后停止
	time.Sleep(time.Second * 10)
	fmt.Println("停止任务1。。。")
	stop <- true
	time.Sleep(time.Second * 3)
	return
}

func run3(ctx context.Context, id int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("任务%v结束退出\n", id)
			return
		default:
			fmt.Printf("任务%v正在运行中\n", id)
			time.Sleep(time.Second * 2)
		}
	}
}

func TestContext10(t *testing.T) {
	//管理启动的协程
	ctx, cancel := context.WithCancel(context.Background())

	// 开启多个goroutine，传入ctx
	go run3(ctx, 1)
	go run3(ctx, 2)

	// 运行一段时间后停止
	time.Sleep(time.Second * 10)
	fmt.Println("停止任务")
	cancel() // 使用context的cancel函数停止goroutine

	// 为了检测监控过是否停止，如果没有监控输出，表示停止
	time.Sleep(time.Second * 3)
	return
}

func coroutine(ctx context.Context, duration time.Duration, id int, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Printf("协程 %d 退出\n", id)
			wg.Done()
			return
		case <-time.After(duration):
			fmt.Printf("消息来自协程 %d\n", id)
		}
	}
}

func TestContext11(t *testing.T) {
	// 使用WaitGroup等待所有的goroutine执行完毕，在收到<-ctx.Done()的终止信号后使wg中需要等待的goroutine数量减一。
	// 因为context只负责取消goroutine，不负责等待goroutine运行，所以需要配合一点辅助手段
	// 管理启动的协程
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go coroutine(ctx, 1*time.Second, i, wg)
	}
	time.Sleep(2 * time.Second)
	cancel()
	wg.Wait()
}
