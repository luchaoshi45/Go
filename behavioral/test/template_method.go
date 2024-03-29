package test

import "Go/behavioral/template_method"

func TemplateMethod() {
	mt := template_method.NewMakeTea("tea1")
	mbt := template_method.NewMakeBeverageTemplate(mt)
	mbt.Run()

	mbt.SetName("tttt")
}
