package aggregation

import "fmt"

// AbstractOrder interface
type AbstractOrder interface {
	GetName() string
	AddOrderItem(orderItem AbstractOrderItem)
	Show()
}

// Order concrete
type Order struct {
	name      string
	orderItem []AbstractOrderItem
}

func NewOrder(name string, orderItem []AbstractOrderItem) AbstractOrder {
	return &Order{name: name, orderItem: orderItem}
}

func (this *Order) GetName() string {
	return this.name
}

func (this *Order) AddOrderItem(orderItem AbstractOrderItem) {
	this.orderItem = append(this.orderItem, orderItem)
}

func (this *Order) Show() {
	fmt.Println(this.GetName())
	for _, item := range this.orderItem {
		fmt.Println(item.GetInfo())
	}
}

// AbstractOrderItem interface
type AbstractOrderItem interface {
	GetInfo() string
}

// OrderItem concrete
type OrderItem struct {
	info string
}

func NewOrderItem(info string) AbstractOrderItem {
	return &OrderItem{info: info}
}

func (this *OrderItem) GetInfo() string {
	return this.info
}
