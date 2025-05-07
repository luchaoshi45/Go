package syntax

import (
	"fmt"
	"sync"
)

func Func() {
	ans := func(nums []int) bool {
		if len(nums) == 0 {
			return false
		} else {
			return true
		}
	}
	fmt.Println(ans([]int{0, 1, 2}))

	var wg sync.WaitGroup
	wg.Add(1)
	go func(s string) {
		defer wg.Done()
		fmt.Println(s)
	}("Hello, Goroutine!")

	wg.Wait()
}
