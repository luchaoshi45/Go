package test

import "Go/creative"

// 逻辑层
func LowCoupling() {
	// 注意业务层 只能访问 抽象层
	var benz creative.Car // 这里必须是 Car 类型的
	benz = new(creative.Benz)
	var bmw creative.Car // 这里必须是 Car 类型的
	bmw = new(creative.BMW)

	var zhang3 creative.Driver
	zhang3 = new(creative.Zhang3)
	zhang3.Drive(benz)

	var li4 creative.Driver
	li4 = new(creative.Li4)
	li4.Drive(bmw)

	var wang5 creative.Driver
	wang5 = new(creative.Wang5)
	wang5.Drive(benz)
	wang5.Drive(bmw)
}
