package realization

import "fmt"

// USB interface
type USB interface {
	Connect()
	Run()
}

// Mouse struct
type Mouse struct {
	name string
}

func NewMouse(name string) USB {
	return &Mouse{name: name}
}

func (this *Mouse) Connect() {
	fmt.Println("Mouse Connect")
}

func (this *Mouse) Run() {
	this.Click()
}

func (this *Mouse) Click() {
	fmt.Println("Mouse Click")
}

// KeyBoard struct
type KeyBoard struct {
	name string
}

func NewKeyBoard(name string) USB {
	return &KeyBoard{name: name}
}

func (this *KeyBoard) Connect() {
	fmt.Println("KeyBoard Connect")
}

func (this *KeyBoard) Run() {
	this.Tap()
}

func (this *KeyBoard) Tap() {
	fmt.Println("KeyBoard Tap")
}
