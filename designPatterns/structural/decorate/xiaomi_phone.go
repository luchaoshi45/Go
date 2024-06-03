package decorate

import "fmt"

type XiaoMiPhone struct {
	Name string
}

func (xmp *XiaoMiPhone) Show() {
	fmt.Println("XiaoMiPhone")
	fmt.Println("Name: ", xmp.Name)
}

func NewXiaoMiPhone(Name string) AbstractPhone {
	return &XiaoMiPhone{Name: Name}
}
