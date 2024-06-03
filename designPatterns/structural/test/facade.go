package test

import "Go/designPatterns/structural/facade"

func Facade() {
	hmp := facade.NewHomePlayerFacade()
	hmp.DoKTV()
	hmp.DoGame()
}
