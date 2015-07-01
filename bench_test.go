package bench

import "testing"

func BenchmarkConcrete(b *testing.B) {
	n := Node(17)
	for i := 0; i < b.N; i++ {
		concrete(n)
	}
}

func BenchmarkInterfacedNode(b *testing.B) {
	n := Node(123)
	for i := 0; i < b.N; i++ {
		interfaced(n)
	}
}

func BenchmarkInterfacedNodeStruct(b *testing.B) {
	n := NodeStruct{17, 123}
	for i := 0; i < b.N; i++ {
		interfaced(n)
	}
}

func BenchmarkTypeSwitchingNode(b *testing.B) {
	n := Node(123)
	for i := 0; i < b.N; i++ {
		typeSwitching(n)
	}
}

func BenchmarkTypeSwitchingNodeStruct(b *testing.B) {
	n := &NodeStruct{17, 123}
	for i := 0; i < b.N; i++ {
		typeSwitching(n)
	}
}
