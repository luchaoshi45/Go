package algorithm

import "fmt"

func DoubleLink(){
	dlink := node{value: -1} // dumpy
	dlink.AddHead(2)
	dlink.AddTail(3)
	dlink.AddHead(1)
	dlink.AddTail(4)
	dlink.PrintList()
	fmt.Println("-----")
	dlink.Reversal()
	dlink.PrintList()
}

type node struct{
	value int
	pre   *node
	next  *node
}

func (n node)PrintList(){
	cur := n.next
	for cur != nil{
		fmt.Println(cur.value)
		cur = cur.next
	}
}

func (n *node)AddHead(value int){
	cur := n.next
	newHead := &node{value: value}
	if cur != nil{
		n.next = newHead
		newHead.pre = n
		newHead.next = cur
		cur.pre = newHead

	} else{
		n.next = newHead
		newHead.pre = n
	}
}

func (n *node)AddTail(value int){
	cur := n.next
	newHead := &node{value: value}
	if cur == nil{
		n.next = newHead
		newHead.pre = n
		return
	}
	for cur.next != nil {
		cur = cur.next // 最后
	}
	cur.next = newHead
	newHead.pre = n
}

func (n *node)Reversal(){
	var pre *node // 初始为空
	var cur = n.next // 指首元素
	if cur == nil{ // 空链表不需要反转
		return
	}

	// 迭代
	for cur != nil{
		next := cur.next
		cur.next = pre
		cur.pre = next
		pre = cur
		cur = next
	}
	n.next = pre // n 指向原来的尾元素
	pre.pre = n
}