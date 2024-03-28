package test

type AbstractApple interface {
	Show()
}

type AbstractFactory interface {
	CrateApple() AbstractApple
}
