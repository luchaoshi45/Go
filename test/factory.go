package test

import "Go/factory"

func Factory() {
	var appleFac factory.AbstractFactory
	var apple factory.AbstractApple

	appleFac = new(factory.ChinaFactory)
	apple = appleFac.CrateApple()
	apple.Show()

	appleFac = new(factory.JapanFactory)
	apple = appleFac.CrateApple()
	apple.Show()

}
