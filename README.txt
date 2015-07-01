PASS

BenchmarkConcrete					1000000000	        2.13 ns/op	       0 B/op	       0 allocs/op
BenchmarkInterfacedNode				30000000	        54.2 ns/op	       8 B/op	       1 allocs/op
BenchmarkInterfacedNodeStruct		20000000	        76.9 ns/op	      16 B/op	       1 allocs/op
BenchmarkTypeSwitchingNode			20000000	        62.8 ns/op	       8 B/op	       1 allocs/op
BenchmarkTypeSwitchingNodeStruct	100000000	        12.9 ns/op	       0 B/op	       0 allocs/op

ok  	github.com/hgfischer/benchinterf	8.268s
