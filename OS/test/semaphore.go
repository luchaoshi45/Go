package test

import (
	"Go/OS/semaphore"
	"context"
	"fmt"
	"sync"
	"time"
)

func tryWorker(id int, s *semaphore.Weighted, wg *sync.WaitGroup) {
	defer wg.Done()

	n := int64(2) // 请求权重为 2

	// 获取信号量，不会阻塞
	if !s.TryAcquire(n) {
		fmt.Println("Failed to TryAcquire ", id)
		return
	}
	defer s.Release(n) // 使用 defer 释放信号量

	fmt.Println("Success to TryAcquire ", id)

	// 模拟工作
	time.Sleep(time.Second)
	fmt.Println("Worker Completed ", id)
}

func TrySemaphore() {
	maxWeight := int64(5) // 最大权重为 5
	s := semaphore.NewWeighted(maxWeight)

	var wg sync.WaitGroup
	maxTaskNum := 10 // 最大任务为 10

	for i := 0; i < maxTaskNum; i++ {
		wg.Add(1)
		go tryWorker(i, s, &wg)
	}
	wg.Wait()
}

func worker(ctx context.Context, id int, s *semaphore.Weighted, wg *sync.WaitGroup) {
	defer wg.Done()

	n := int64(2) // 请求权重为 2

	// 获取信号量，阻塞
	if err := s.Acquire(ctx, n); err != nil {
		fmt.Println("Failed to Acquire ", id)
		return
	}
	defer s.Release(n) // 使用 defer 释放信号量

	fmt.Println("Success to Acquire ", id)

	// 模拟工作
	time.Sleep(time.Second)
	fmt.Println("Worker Completed ", id)
}

func Semaphore() {
	ctx := context.Background()
	maxWeight := int64(5) // 最大权重为 5
	s := semaphore.NewWeighted(maxWeight)

	var wg sync.WaitGroup
	maxTaskNum := 10 // 最大任务为 10

	for i := 0; i < maxTaskNum; i++ {
		wg.Add(1)
		go worker(ctx, i, s, &wg)
	}
	wg.Wait()
}
