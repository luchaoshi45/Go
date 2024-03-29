package test

import "Go/structural/facade"

func Facade() {
	hmp := facade.NewHomePlayerFacade()
	hmp.DoKTV()
	hmp.DoGame()
}
