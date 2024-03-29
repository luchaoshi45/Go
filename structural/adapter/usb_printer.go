package adapter

import "fmt"

type USBPrinter struct {
	name string
}

func NewUSBPrinter(name string) *USBPrinter {
	return &USBPrinter{name: name}
}

func (usbp *USBPrinter) Connect(computer Computer) {
	computer.USBPort()
}

func (usbp *USBPrinter) Run() {
	fmt.Println("Run")
}
