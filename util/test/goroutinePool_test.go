package test

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

const (
	maxWorkers = 2 // 最大工作协程数
	queueSize  = 1 // 任务队列大小
)

// GoroutinePool 结构体包含工作协程数、任务队列和同步原语
type GoroutinePool struct {
	taskQueue chan func() error // 任务队列
	results   chan error        // 存放任务执行结果的channel
	wg        sync.WaitGroup
}

// NewGoroutinePool 创建并启动一个新的协程池
func NewGoroutinePool() *GoroutinePool {
	pool := &GoroutinePool{
		taskQueue: make(chan func() error, queueSize),
		results:   make(chan error, queueSize), // 确保结果channel与任务队列大小一致
	}
	// 启动工作协程
	pool.wg.Add(maxWorkers)
	for i := 0; i < maxWorkers; i++ {
		go pool.worker()
	}
	return pool
}

// worker 是工作协程的主体
func (gp *GoroutinePool) worker() {
	defer gp.wg.Done()
	for {
		// 阻塞等待任务
		job, ok := <-gp.taskQueue
		if !ok { // 如果任务队列已关闭，退出工作协程
			break
		}
		// 执行任务，并将结果发送到结果channel
		err := job()
		gp.results <- err
	}
}

// Submit 提交一个任务到协程池
func (gp *GoroutinePool) Submit(job func() error) (error, bool) {
	// 尝试将任务放入任务队列
	select {
	case gp.taskQueue <- job:
		// 任务放入任务队列，从结果channel获取执行结果
		result := <-gp.results
		return result, true
	default:
		// 任务队列已满，任务提交失败
		return nil, false
	}
}

// Close 关闭协程池，不再接受新任务
func (gp *GoroutinePool) Close() {
	// 关闭任务队列，通知所有工作协程退出
	close(gp.taskQueue)
	gp.wg.Wait()
	// 清空结果channel
	for len(gp.results) > 0 {
		<-gp.results
	}
	close(gp.results)
}

func TestName(t *testing.T) {
	pool := NewGoroutinePool()

	// 提交任务
	for i := 0; i < 6; i++ {
		jobIndex := i
		job := func() error {
			//fmt.Printf("执行任务: %d\n", jobIndex)
			time.Sleep(time.Second)
			// 模拟任务执行，这里没有错误返回nil
			return nil
		}
		go func() {
			result, ok := pool.Submit(job)
			if ok {
				fmt.Printf("任务 %d 提交成功，结果：%v\n", jobIndex, result)
			} else {
				fmt.Printf("任务 %d 提交失败，任务队列已满\n", jobIndex)
			}
		}()
	}

	time.Sleep(time.Second * 100)

	// 关闭协程池
	pool.Close()
}
