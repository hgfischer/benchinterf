package bench

func concrete(n *Node) {
	n.ID()
}

func interfaced(n NodeInterface) {
	n.ID()
}

func typeAssertion(n NodeInterface) {
	if c, ok := n.(*NodeStruct); ok {
		c.ID()
	} else {
		n.ID()
	}
}

func typeSwitching(n NodeInterface) {
	switch n := n.(type) {
	case *Node:
		n.ID()
	case *NodeStruct:
		n.ID()
	}
}
