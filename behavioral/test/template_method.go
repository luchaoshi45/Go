package test

import "Go/behavioral/template_method"

func TemplateMethod() {
	mt := template_method.NewMakeTea("tea")
	mbt := template_method.NewMakeBeverageTemplate(mt)
	mbt.Run()
	mbt.SetName("ttt")

	mc := template_method.NewMakeCoffee("coffee")
	mbc := template_method.NewMakeBeverageTemplate(mc)
	mbc.Run()
	mbc.SetName("ccc")
}
