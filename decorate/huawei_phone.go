package decorate

import "fmt"

type HuaWeiPhone struct {
	Name string
}

func (hwp *HuaWeiPhone) Show() {
	fmt.Println("HuaWeiPhone")
	fmt.Println("Name: ", hwp.Name)
}

func NewHuaWeiPhone(Name string) AbstractPhone {
	return &HuaWeiPhone{Name: Name}
}
