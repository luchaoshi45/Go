package test

import "fmt"

type ChinaApple struct{}
type ChinaFactory struct{}

func (ca *ChinaApple) Show() {
	fmt.Println("ChinaApple")
}

func (cf *ChinaFactory) CrateApple() AbstractApple {
	var apple AbstractApple
	apple = new(ChinaApple)
	return apple
}
