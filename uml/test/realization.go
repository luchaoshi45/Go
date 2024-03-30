package test

import "Go/uml/realization"

func Realization() {
	mouse := realization.NewMouse("mouse")
	mouse.Connect()
	mouse.Run()

	keyBoard := realization.NewKeyBoard("mouse")
	keyBoard.Connect()
	keyBoard.Run()
}
