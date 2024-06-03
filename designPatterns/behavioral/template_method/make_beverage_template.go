package template_method

import "fmt"

type MakeBeverageTemplate struct {
	mb MakeBeverage
}

func NewMakeBeverageTemplate(mb MakeBeverage) *MakeBeverageTemplate {
	return &MakeBeverageTemplate{mb: mb}
}

func (mbt *MakeBeverageTemplate) Run() {
	mbt.mb.BoilWater()
	mbt.mb.Brew()
	mbt.mb.PourCup()
	mbt.mb.AddThings()
	fmt.Println("Make ", mbt.mb.GetName(), " Complete")
}

func (mbt *MakeBeverageTemplate) SetName(Name string) {
	mbt.mb.SetName(Name)
	fmt.Println("New Name: ", mbt.mb.GetName())
}
