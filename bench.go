package bench

func concrete(n Node) {
	n.ID()
}

func interfaced(n NodeInterface) {
	n.ID()
}

func typeSwitching(n NodeInterface) {
	if c, ok := n.(*NodeStruct); ok {
		c.ID()
	} else {
		n.ID()
	}
}
