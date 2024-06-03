package factory

type AbstractApple interface {
	Show()
}

type AbstractFactory interface {
	CrateApple() AbstractApple
}
