package test

import (
	"Go/structural/proxy"
	"fmt"
	"sync"
)

func Proxy() {
	ceo := proxy.NewCEO(proxy.NewSimpleEmployee(true))
	ceo.DoWork()
	ceo.CheckWork()

	CEOFac := new(proxy.CEOFactory)
	ceo2 := CEOFac.CrateEmployee(proxy.NewSimpleEmployee(true))
	ceo2.DoWork()
	ceo2.CheckWork()

	ceo3 := proxy.NewCEO(proxy.NewAdvanceEmployee(true))

	ceo3.DoWork()
	ceo3.SetLaidOff(false)
	ceo3.CheckWork()
	ceo3.SetLaidOff(true)
	ceo3.CheckWork()

	var wg sync.WaitGroup

	for i := 0; i < 10000000; i++ { // 1kw 并发
		wg.Add(1)
		go func() {
			ceox := proxy.NewCEO(proxy.NewAdvanceEmployee(true))
			//ceox.DoWork()
			ceox.SetLaidOff(false)
			//ceox.CheckWork()
			ceox.SetLaidOff(true)
			//ceox.CheckWork()
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("所有任务完成")
}
