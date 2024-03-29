package association

import "fmt"

// AbstractNode 抽象Node
type AbstractNode interface {
	GetName() string
	SetName(Name string)
	GetParentNode() AbstractNode
	SetParentNode(SubNode AbstractNode)
	PrintToRoot()
}

// Node 实现Node
type Node struct {
	Name       string
	ParentNode AbstractNode
}

func NewNode(Name string, ParentNode AbstractNode) AbstractNode {
	return &Node{Name: Name, ParentNode: ParentNode}
}

func (this *Node) GetName() string {
	return this.Name
}

func (this *Node) SetName(Name string) {
	this.Name = Name
}

func (this *Node) GetParentNode() AbstractNode {
	return this.ParentNode
}

func (this *Node) SetParentNode(ParentNode AbstractNode) {
	this.ParentNode = ParentNode
}

func (this *Node) PrintToRoot() {
	fmt.Println(this.Name)
	if this.ParentNode != nil {
		this.ParentNode.PrintToRoot()
	}
}
