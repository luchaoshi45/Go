package test

import (
	"Go/uml/aggregation"
	"fmt"
)

func Aggregation() {
	mit := aggregation.NewMIT("MIT", nil)

	zhang3 := aggregation.NewTeacher("zhang3")
	fmt.Println(zhang3.GetName())
	li4 := aggregation.NewTeacher("li4")
	fmt.Println(li4.GetName())

	mit.AddTeacher(zhang3)
	mit.AddTeacher(li4)
	fmt.Println(mit.GetName())

}
