package test

import aggregation "Go/uml/combination"

func Combination() {
	order := aggregation.NewOrder("order", nil)
	order.Show()
	item1 := aggregation.NewOrderItem("item1")
	item2 := aggregation.NewOrderItem("item2")
	item3 := aggregation.NewOrderItem("item3")
	order.AddOrderItem(item1)
	order.AddOrderItem(item2)
	order.AddOrderItem(item3)
	order.Show()
}
