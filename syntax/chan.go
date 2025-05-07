package syntax

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 优雅退出 waitGroup chan context go select
func Chan() {
	var wg sync.WaitGroup

	chan1 := make(chan string)
	chan2 := make(chan int, 10)

	ctx, cancel := context.WithCancel(context.Background())

	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case ch1 := <-chan1:
				fmt.Println(ch1)
			case ch2 := <-chan2:
				fmt.Println(ch2)
			case <-ctx.Done():
				fmt.Println("Done")
				return // 注意返回
			}
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		chan1 <- "1"
		chan2 <- 2
		time.Sleep(1 * time.Second)
		cancel()
	}()

	wg.Wait()
}
