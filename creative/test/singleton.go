package test

import (
	"Go/creative/singleton"
	"fmt"
	"sync"
)

func Singlet() {
	st := singleton.GetInstance()
	st.Show()
	fmt.Printf("st address: %p\n", st)

	st2 := singleton.GetInstance()
	st2.Show()
	fmt.Printf("st2 address: %p\n", st2)

	var wg sync.WaitGroup
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			instance := singleton.GetInstance()
			fmt.Printf("Instance address: %p\n", instance)
		}()
	}
	wg.Wait()
}
