package syntax

import (
	"fmt"
	"sync"
)

func Mutex() {
	var mutex sync.Mutex
	var wg sync.WaitGroup
	var count int

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mutex.Lock()
			count++
			mutex.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println(count)
}
