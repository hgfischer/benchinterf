package bench

func concrete(n *Node) {
	_ = n.ID()
}

func interfaced(n NodeInterface) {
	_ = n.ID()
}

func typeAssertion(n NodeInterface) {
	if c, ok := n.(*NodeStruct); ok {
		_ = c.ID()
	} else {
		_ = n.ID()
	}
}

func typeSwitching(n NodeInterface) {
	switch n := n.(type) {
	case *Node:
		_ = n.ID()
	case *NodeStruct:
		_ = n.ID()
	}
}
