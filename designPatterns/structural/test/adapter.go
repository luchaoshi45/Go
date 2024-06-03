package test

import (
	"Go/designPatterns/structural/adapter"
	"fmt"
)

func Adapter() {
	macAdapter := adapter.NewMacAdapter()
	macAdapter.USBPort()
	windows := adapter.NewWindows()
	windows.USBPort()

	fmt.Println("_______________")
	usbPrinter := adapter.NewUSBPrinter("p1")
	usbPrinter.Connect(macAdapter)
	usbPrinter.Connect(windows)
	usbPrinter.Run()
}
