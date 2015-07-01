## Ignoring return

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

## Forcing return to `_`

```
PASS
BenchmarkConcrete                 1000000000  2.10  ns/op  0  B/op  0  allocs/op
BenchmarkInterfacedNode           300000000   4.78  ns/op  0  B/op  0  allocs/op
BenchmarkInterfacedNodeStruct     300000000   4.48  ns/op  0  B/op  0  allocs/op
BenchmarkTypeAssertionNode        100000000   11.6  ns/op  0  B/op  0  allocs/op
BenchmarkTypeAssertionNodeStruct  100000000   12.3  ns/op  0  B/op  0  allocs/op
BenchmarkTypeSwitchingNode        100000000   14.4  ns/op  0  B/op  0  allocs/op
BenchmarkTypeSwitchingNodeStruct  100000000   14.7  ns/op  0  B/op  0  allocs/op
```
