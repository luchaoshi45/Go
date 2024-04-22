package command

type Waiter struct {
	CmdList []Command
}

func (w *Waiter) Notify() {
	if w.CmdList == nil {
		return
	}

	for _, cmd := range w.CmdList {
		cmd.Work() // 订单多态
	}
}
