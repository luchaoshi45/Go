package adapter

import "fmt"

// 适配者
type Mac struct {
}

func (m *Mac) LightningPort() {
	fmt.Println("Mac LightningPort")
}

type MacAdapter struct {
	m *Mac
}

func (a *MacAdapter) USBPort() {
	fmt.Print("Adapter USBPort : ")
	a.m.LightningPort()
}

func NewMacAdapter() Computer {
	return &MacAdapter{m: new(Mac)}
}
