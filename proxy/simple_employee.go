package proxy

import "fmt"

type SimpleEmployee struct {
	LaidOff bool
}

func NewSimpleEmployee(LaidOff bool) AbstractEmployee {
	return &SimpleEmployee{LaidOff: LaidOff}
}

func (se *SimpleEmployee) DoWork() {
	fmt.Println("SimpleEmployee DoWork")
}

func (se *SimpleEmployee) CheckWork() {
	fmt.Println("SimpleEmployee CheckWork")
}
func (se *SimpleEmployee) GetLaidOff() bool {
	return se.LaidOff
}
func (se *SimpleEmployee) SetLaidOff(LaidOff bool) {
	se.LaidOff = LaidOff
}
