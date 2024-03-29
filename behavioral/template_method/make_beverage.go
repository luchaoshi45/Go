package template_method

type MakeBeverage interface {
	BoilWater()
	Brew()
	PourCup()
	AddThings()
	GetName() string
	SetName(Name string)
}
