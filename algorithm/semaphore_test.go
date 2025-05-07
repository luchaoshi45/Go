package algorithm

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// 测试信号量的并发行为
func TestSemaphore(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 初始化信号量，最多允许 3 个并发任务
	sem := &Weighted{tokens: 3}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ { // 启动 10 个任务
		wg.Add(1)
		go tryWorker(i, sem, &wg)
	}

	// 等待所有任务完成
	go func() {
		wg.Wait()
		fmt.Println("All workers completed.")
	}()

	// 如果超时未完成，则失败
	select {
	case <-ctx.Done():
		t.Error("Test timed out")
	case <-time.After(5 * time.Second):
		// 正常完成
	}
}

// tryWorker 是每个并发任务的执行逻辑
func tryWorker(id int, s *Weighted, wg *sync.WaitGroup) {
	defer wg.Done()

	// 获取信号量
	s.Acquire(1)
	fmt.Printf("Worker %d acquired semaphore\n", id)

	// 模拟任务处理
	time.Sleep(time.Second)
	fmt.Printf("Worker %d finished processing\n", id)

	// 释放信号量
	s.Release(1)
}
