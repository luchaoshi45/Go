package inheritance

import (
	"fmt"
	"reflect"
)

// Animal interface
type Animal interface {
	Eat()
	Sleep()
}

// AbsFac interface
type AbsFac interface {
	CreateAnimal() Animal
}

// Human struct
type Human struct {
	name string
}

// HumanFac struct
type HumanFac struct {
}

func (this *HumanFac) CreateAnimal(name string) Animal {
	return &Human{name: name}
}

func (this *Human) Eat() {
	fmt.Println(reflect.TypeOf(this).Elem().Name(), "Eat")
}
func (this *Human) Sleep() {
	fmt.Println(reflect.TypeOf(this).Elem().Name(), "Sleep")
}

// Teacher struct
type Teacher struct {
	Human
	name string
}

// TeacherFac struct
type TeacherFac struct {
}

func (this *TeacherFac) CreateAnimal(name string) Animal {
	return &Teacher{name: name}
}

func (this *Teacher) Teaching() {
	fmt.Println("Teaching")
}

// Student struct
type Student struct {
	Human
	name string
}

// StudentFac struct
type StudentFac struct {
}

func (this *StudentFac) CreateAnimal(name string) Animal {
	return &Student{name: name}
}

func (this *Student) Studying() {
	fmt.Println("Studying")
}
