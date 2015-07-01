```
PASS
BenchmarkConcrete                 1000000000  2.13  ns/op  0  B/op  0  allocs/op
BenchmarkInterfacedNode           300000000   4.83  ns/op  0  B/op  0  allocs/op
BenchmarkInterfacedNodeStruct     300000000   4.52  ns/op  0  B/op  0  allocs/op
BenchmarkTypeAssertionNode        100000000   11.7  ns/op  0  B/op  0  allocs/op
BenchmarkTypeAssertionNodeStruct  100000000   12.4  ns/op  0  B/op  0  allocs/op
BenchmarkTypeSwitchingNode        100000000   14.6  ns/op  0  B/op  0  allocs/op
BenchmarkTypeSwitchingNodeStruct  100000000   14.8  ns/op  0  B/op  0  allocs/op
```
