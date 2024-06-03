package creative

import "fmt"

// 抽象层
type Driver interface {
	Drive(car Car)
}
type Car interface {
	Run()
}

// 实现层  只能访问它的抽象层
type Benz struct {
}

func (b *Benz) Run() {
	fmt.Println("Benz Run")
}

type BMW struct {
}

func (b *BMW) Run() {
	fmt.Println("BMW Run")
}

type Zhang3 struct {
}

func (z *Zhang3) Drive(car Car) {
	fmt.Println("Zhang3 Drive Car")
	car.Run()
}

type Li4 struct {
}

func (l *Li4) Drive(car Car) {
	fmt.Println("Li4 Drive Car")
	car.Run()
}

type Wang5 struct {
}

func (w *Wang5) Drive(car Car) {
	fmt.Println("Wang5 Drive Car")
	car.Run()
}
