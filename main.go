package main

import "Go/test"

func main() {
	var appleFac test.AbstractFactory
	var apple test.AbstractApple

	appleFac = new(test.ChinaFactory)
	apple = appleFac.CrateApple()
	apple.Show()

	appleFac = new(test.JapanFactory)
	apple = appleFac.CrateApple()
	apple.Show()

}
