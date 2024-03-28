package proxy

import "fmt"

type AdvanceEmployee struct {
	LaidOff bool
}

func NewAdvanceEmployee(LaidOff bool) AbstractEmployee {
	return &AdvanceEmployee{LaidOff: LaidOff}
}

func (ae *AdvanceEmployee) DoWork() {
	fmt.Println("AdvanceEmployee DoWork")
}

func (ae *AdvanceEmployee) CheckWork() {
	fmt.Println("AdvanceEmployee CheckWork")
}

func (ae *AdvanceEmployee) GetLaidOff() bool {
	return ae.LaidOff
}
func (ae *AdvanceEmployee) SetLaidOff(LaidOff bool) {
	ae.LaidOff = LaidOff
}
