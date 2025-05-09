package syntax

import (
	"context"
	"fmt"
	"sync"
)

// 注意取消的 chan 不能读取数据了
func Print() {
	var wg sync.WaitGroup

	chan1 := make(chan int)
	chan2 := make(chan int)

	ctx, cancel := context.WithCancel(context.Background())

	nums := []int{1, 2, 3, 4, 5}
	nums_i := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-chan1:
				fmt.Println(nums[nums_i])
				nums_i++
				chan2 <- 0
			case <-ctx.Done():
				return
			}
		}
	}()

	strs := []byte{'a', 'b', 'c', 'd', 'e'}
	strs_i := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case <-chan2:
				fmt.Println(string(strs[strs_i]))
				strs_i++
				if strs_i >= 5 {
					cancel()
					return
				}
				chan1 <- 0
			}
		}
	}()

	chan1 <- 0
	wg.Wait()
}
