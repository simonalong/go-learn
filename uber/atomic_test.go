package test

import (
	"fmt"
	"go.uber.org/atomic"
	"testing"
	"time"
)

// Semaphore 结构体使用 atomic.Value 来存储当前的计数
type Semaphore struct {
	count *atomic.Uint32
}

// NewSemaphore 创建一个新的信号量
func NewSemaphore(initial int) *Semaphore {
	return &Semaphore{
		count: atomic.NewUint32(uint32(initial)),
	}
}

// Acquire 方法尝试获取资源，如果信号量的值为0，则阻塞等待
func (s *Semaphore) Acquire() {
	for {
		current := s.count.Load()
		if current == 0 {
			// 如果信号量的值为0，则阻塞等待
			time.Sleep(10 * time.Millisecond)
			continue
		}
		// 尝试将信号量的值减1
		if s.count.CAS(current, current-1) {
			break
		}
	}
}

func (s *Semaphore) tryAcquire() bool {
	current := s.count.Load()
	if current == 0 {
		return false // 信号量已无可用资源
	}
	// 尝试原子地将信号量的值减1
	return s.count.CAS(current, current-1)
}

// Release 方法释放资源，将信号量的值加1
func (s *Semaphore) Release() {
	s.count.Add(1)
}

func TestName(t *testing.T) {

	sem := NewSemaphore(3) // 创建一个初始值为3的信号量

	// 启动一些协程来模拟对信号量的访问
	for i := 0; i < 5; i++ {
		go func(id int) {
			fmt.Printf("Goroutine #%d is trying to acquire the semaphore\n", id)
			sem.Acquire()
			fmt.Printf("Goroutine #%d has acquired the semaphore\n", id)
			// 模拟使用资源
			time.Sleep(500 * time.Millisecond)
			sem.Release()
			fmt.Printf("Goroutine #%d has released the semaphore\n", id)
		}(i)
	}

	// 给所有协程足够的时间来完成
	time.Sleep(5 * time.Second)
}
