package adapter

import "fmt"

// 适配者
type Windows struct {
}

func (m *Windows) USBPort() {
	fmt.Println("Windows USBPort")
}

func NewWindows() Computer {
	return new(Windows)
}
