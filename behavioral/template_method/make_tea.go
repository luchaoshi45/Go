package template_method

import "fmt"

type MakeTea struct {
	Name string
}

func NewMakeTea(Name string) MakeBeverage {
	return &MakeTea{Name: Name}
}

func (mk *MakeTea) BoilWater() {
	fmt.Println("MakeTea BoilWater")
}
func (mk *MakeTea) Brew() {
	fmt.Println("MakeTea Brew")
}
func (mk *MakeTea) PourCup() {
	fmt.Println("MakeTea PourCup")
}
func (mk *MakeTea) AddThings() {
	fmt.Println("MakeTea AddThings")
}

func (mk *MakeTea) GetName() string {
	return mk.Name
}
func (mk *MakeTea) SetName(Name string) {
	mk.Name = Name
}
