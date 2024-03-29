package test

import "Go/uml/association"

func SelfAssociation() {
	nodeRoot := association.NewNode("root", nil)
	node1 := association.NewNode("node1", nil)
	node2 := association.NewNode("node2", nil)
	node3 := association.NewNode("node3", nil)
	node4 := association.NewNode("node4", nil)
	node5 := association.NewNode("node5", nil)
	node6 := association.NewNode("node6", nil)

	node3.SetParentNode(node1)
	node4.SetParentNode(node1)
	node5.SetParentNode(node2)
	node6.SetParentNode(node2)
	node1.SetParentNode(nodeRoot)
	node2.SetParentNode(nodeRoot)
	nodeRoot.SetName("nodeRoot")
	node6.PrintToRoot()
}
