```
PASS
BenchmarkConcrete                 1000000000  2.15  ns/op  0   B/op  0  allocs/op
BenchmarkInterfacedNode           30000000    55.4  ns/op  8   B/op  1  allocs/op
BenchmarkInterfacedNodeStruct     20000000    76.9  ns/op  16  B/op  1  allocs/op
BenchmarkTypeAssertionNode        20000000    62.3  ns/op  8   B/op  1  allocs/op
BenchmarkTypeAssertionNodeStruct  100000000   12.8  ns/op  0   B/op  0  allocs/op
BenchmarkTypeSwitchingNode        20000000    65.1  ns/op  8   B/op  1  allocs/op
BenchmarkTypeSwitchingNodeStruct  200000000   6.60  ns/op  0   B/op  0  allocs/op
```
