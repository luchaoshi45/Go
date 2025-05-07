package syntax

import (
	"fmt"
	"sync"
)

func Wait() {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Func1")
	}()

	wg.Add(1)
	go func(str string) {
		defer wg.Done()
		fmt.Println(str)
	}("Func2")

	wg.Wait()
}
