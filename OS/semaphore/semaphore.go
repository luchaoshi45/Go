package semaphore

import (
	"container/list"
	"context"
	"sync"
)

// waiter 结构体定义了一个等待信号量的goroutine
type waiter struct {
	n     int64           // 请求的权重
	ready chan<- struct{} // 在信号量被获取后会关闭
}

// Weighted 结构体提供了一个有权重的信号量实现
type Weighted struct {
	size    int64      // 信号量的最大权重
	cur     int64      // 当前已经被获取的权重
	mu      sync.Mutex // 互斥锁，保护下面的字段
	waiters list.List  // 存储等待获取信号量的goroutine的链表
}

// NewWeighted 创建一个新的有权重的信号量，n为最大的并发访问权重
func NewWeighted(n int64) *Weighted {
	w := &Weighted{size: n}
	return w
}

// Acquire 获取一个权重为n的信号量，阻塞直到资源可用或ctx被取消
func (s *Weighted) Acquire(ctx context.Context, n int64) error {
	done := ctx.Done()

	s.mu.Lock()
	select {
	case <-done:
		// 如果ctx已经取消，则立即返回
		s.mu.Unlock()
		return ctx.Err()
	default:
	}
	// 如果当前已经获取的权重+请求的权重小于等于信号量的最大权重，并且等待队列为空，则获取信号量成功
	if s.size-s.cur >= n && s.waiters.Len() == 0 {
		s.cur += n
		s.mu.Unlock()
		return nil
	}

	if n > s.size {
		// 如果请求的权重大于信号量的最大权重，立即返回失败
		s.mu.Unlock()
		<-done
		return ctx.Err()
	}

	ready := make(chan struct{})
	w := waiter{n: n, ready: ready}
	elem := s.waiters.PushBack(w)
	s.mu.Unlock()

	select {
	case <-done:
		// 如果ctx已经取消，则处理等待队列中的goroutine
		s.mu.Lock()
		select {
		case <-ready:
			// 如果在ctx取消之前获取了信号量，则回退并唤醒其他等待的goroutine
			s.cur -= n
			s.notifyWaiters()
		default:
			// 如果在ctx取消之前没有获取信号量，则从等待队列中移除
			isFront := s.waiters.Front() == elem
			s.waiters.Remove(elem)
			// 如果当前等待队列的第一个元素并没有获取信号量，则唤醒其他等待的goroutine
			if isFront && s.size > s.cur {
				s.notifyWaiters()
			}
		}
		s.mu.Unlock()
		return ctx.Err()

	case <-ready:
		// 如果成功获取了信号量，则检查ctx是否已经取消
		select {
		case <-done:
			// 如果ctx在获取信号量之后被取消，则释放信号量并返回错误
			s.Release(n)
			return ctx.Err()
		default:
		}
		return nil
	}
}

// TryAcquire 尝试获取一个权重为n的信号量，不会阻塞
func (s *Weighted) TryAcquire(n int64) bool {
	s.mu.Lock()
	success := s.size-s.cur >= n && s.waiters.Len() == 0
	if success {
		s.cur += n
	}
	s.mu.Unlock()
	return success
}

// Release 释放一个权重为n的信号量
func (s *Weighted) Release(n int64) {
	s.mu.Lock()
	s.cur -= n
	if s.cur < 0 {
		s.mu.Unlock()
		panic("semaphore: released more than held")
	}
	s.notifyWaiters()
	s.mu.Unlock()
}

// notifyWaiters 唤醒等待队列中的goroutine
func (s *Weighted) notifyWaiters() {
	for {
		next := s.waiters.Front()
		if next == nil {
			break // 没有等待的goroutine
		}

		w := next.Value.(waiter)
		if s.size-s.cur < w.n {
			// 如果当前剩余的信号量不够下一个等待的goroutine，则停止唤醒等待的goroutine
			break
		}

		s.cur += w.n
		s.waiters.Remove(next)
		close(w.ready)
	}
}
