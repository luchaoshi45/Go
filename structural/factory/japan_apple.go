package factory

import "fmt"

type JapanApple struct{}
type JapanFactory struct{}

func (ca *JapanApple) Show() {
	fmt.Println("JapanApple")
}

func (cf *JapanFactory) CrateApple() AbstractApple {
	var apple AbstractApple
	apple = new(JapanApple)
	return apple
}
