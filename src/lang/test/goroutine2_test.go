package test

import (
	"fmt"
	"sync"
	"testing"
)

const queueCapacity = 10 // 假设队列的容量为10

// GoroutinePool 结构体定义
type GoroutinePool struct {
	workers int
	queue   chan func()
	wg      sync.WaitGroup
}

// NewGoroutinePool 创建一个新的协程池
func NewGoroutinePool(workers int) *GoroutinePool {
	return &GoroutinePool{
		workers: workers,
		queue:   make(chan func(), queueCapacity), // 设置队列容量
	}
}

// Execute 在协程池中执行任务
func (p *GoroutinePool) Execute(job func()) {
	// 使用非阻塞方式向队列发送任务
	select {
	case p.queue <- job:
		p.wg.Add(1)
	default:
		// 如果队列已满，这里可以选择打印日志或者处理任务被拒绝的情况
		fmt.Println("任务提交失败，协程池已满")
	}
}

// Start 启动协程池
func (p *GoroutinePool) Start() {
	for i := 0; i < p.workers; i++ {
		go func() {
			for job := range p.queue { // 从队列中接收任务
				job() // 执行任务
				p.wg.Done()
			}
		}()
	}
}

// Wait 等待所有任务完成
func (p *GoroutinePool) Wait() {
	p.wg.Wait()
}

// IsFull 判断协程池的任务队列是否已满
func (p *GoroutinePool) IsFull() bool {
	return len(p.queue) == cap(p.queue)
}

func TestName(t *testing.T) {
	// 创建一个具有3个工作协程的协程池
	pool := NewGoroutinePool(3)
	pool.Start()

	// 向协程池提交任务
	for i := 0; i < 15; i++ {
		task := func() {
			fmt.Println("执行任务:", i)
		}
		// 检查协程池是否满，但注意这里检查可能存在竞态条件，因为队列状态可能瞬间改变
		if pool.IsFull() {
			fmt.Println("协程池已满，任务将被丢弃")
		} else {
			pool.Execute(task)
		}
	}

	// 等待所有任务完成
	pool.Wait()

	fmt.Println("所有任务执行完毕")
}
