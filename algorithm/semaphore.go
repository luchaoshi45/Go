package algorithm

import (
	"sync"
	"time"
)

// Weighted 是一个简单的加权信号量实现
type Weighted struct {
	mu     sync.Mutex
	tokens int
}

// Acquire 尝试获取指定数量的令牌
func (s *Weighted) Acquire(tokens int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	for s.tokens < tokens {
		s.mu.Unlock()
		time.Sleep(time.Millisecond) // 等待令牌可用
		s.mu.Lock()
	}
	s.tokens -= tokens
}

// Release 释放指定数量的令牌
func (s *Weighted) Release(tokens int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.tokens += tokens
}
