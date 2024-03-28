package proxy

import "fmt"

type CEO struct {
	LaidOff bool
	ae      AbstractEmployee
}

// 简单工厂 模式
func NewCEO(ae AbstractEmployee) AbstractEmployee {
	return &CEO{LaidOff: true, ae: ae}
}

type CEOFactory struct{}

func (ceof *CEOFactory) CrateEmployee(ae AbstractEmployee) AbstractEmployee {
	return &CEO{LaidOff: true, ae: ae}
}

func (ceo *CEO) DoWork() {
	ceo.ae.DoWork()
}

func (ceo *CEO) CheckWork() {
	ceo.ae.CheckWork()
	ceo.layoffs()
}

func (ceo *CEO) layoffs() {
	if ceo.ae.GetLaidOff() == false {
		fmt.Println("员工被开除")
	} else {
		fmt.Println("员工在工作")
	}
}

func (ceo *CEO) GetLaidOff() bool {
	return ceo.ae.GetLaidOff()
}
func (ceo *CEO) SetLaidOff(LaidOff bool) {
	ceo.ae.SetLaidOff(LaidOff)
}
