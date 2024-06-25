package test

//
//import (
//	"fmt"
//	"sync"
//	"testing"
//)
//
//const queueCapacity = 10 // 假设队列的容量为10
//
//// GoroutinePool 结构体定义
//type GoroutinePool struct {
//	workers int
//	queue   chan func()
//	wg      sync.WaitGroup
//}
//
//// NewGoroutinePool 创建一个新的协程池
//func NewGoroutinePool(workers int) *GoroutinePool {
//	return &GoroutinePool{
//		workers: workers,
//		queue:   make(chan func(), queueCapacity), // 设置队列容量
//	}
//}
//
//// Execute 在协程池中执行任务
//func (p *GoroutinePool) Execute(job func()) bool {
//	select {
//	case p.queue <- job:
//		p.wg.Add(1)
//		return true
//	default:
//		return false // 队列已满，任务未被接受
//	}
//}
//
//// Start 启动协程池
//func (p *GoroutinePool) Start() {
//	for i := 0; i < p.workers; i++ {
//		go func() {
//			for {
//				job, more := <-p.queue
//				if !more {
//					return
//				}
//				job()
//				p.wg.Done()
//			}
//		}()
//	}
//}
//
//// Wait 等待所有任务完成
//func (p *GoroutinePool) Wait() {
//	p.wg.Wait()
//}
//
//// IsFull 判断协程池的任务队列是否已满
//func (p *GoroutinePool) IsFull() bool {
//	return len(p.queue) == cap(p.queue)
//}
//
//func TestName(t *testing.T) {
//
//	// 创建一个具有3个工作协程的协程池
//	pool := NewGoroutinePool(5)
//	pool.Start()
//
//	// 向协程池提交任务
//	for i := 0; i < 15; i++ {
//		if pool.IsFull() {
//			fmt.Println("协程池已满，任务将被丢弃或返回错误")
//		} else {
//			task := func() {
//				fmt.Println("执行任务:", i)
//			}
//			success := pool.Execute(task)
//			if !success {
//				fmt.Println("任务未能提交到协程池，队列已满")
//			}
//		}
//	}
//
//	// 等待所有任务完成
//	pool.Wait()
//
//	fmt.Println("所有任务执行完毕")
//}
