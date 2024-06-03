package facade

import "fmt"

// MicroPhone 麦克风
type MicroPhone struct{}

func (m *MicroPhone) On() {
	fmt.Println("打开 麦克风")
}

func (m *MicroPhone) Off() {
	fmt.Println("关闭 麦克风")
}
