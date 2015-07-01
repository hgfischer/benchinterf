package bench

type NodeInterface interface {
	ID() int64
}

type Node int64

func (n Node) ID() int64 {
	return int64(n)
}

type NodeStruct struct {
	Id   int64
	Data int64
}

func (n NodeStruct) ID() int64 {
	return n.Id
}
