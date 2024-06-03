package template_method

import "fmt"

type MakeCoffee struct {
	Name string
}

func NewMakeCoffee(Name string) MakeBeverage {
	return &MakeCoffee{Name: Name}
}

func (mc *MakeCoffee) BoilWater() {
	fmt.Println("MakeCoffee BoilWater")
}
func (mc *MakeCoffee) Brew() {
	fmt.Println("MakeCoffee Brew")
}
func (mc *MakeCoffee) PourCup() {
	fmt.Println("MakeCoffee PourCup")
}
func (mc *MakeCoffee) AddThings() {
	fmt.Println("MakeCoffee AddThings")
}

func (mc *MakeCoffee) GetName() string {
	return mc.Name
}
func (mc *MakeCoffee) SetName(Name string) {
	mc.Name = Name
}
