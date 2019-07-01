Benchmarking some different ways to reverse ints (`1234` => `4321`) and for
checking if ints are palindromes (`4321234` etc).

```
ᐳ go test -bench .
PASS
BenchmarkReverseWithStrings-8           10000000               156 ns/op
BenchmarkReverseWithMath-8              50000000                27.2 ns/op
BenchmarkBigPalindromePassWithStrings-8 20000000                91.8 ns/op
BenchmarkBigPalindromePassWithMath-8    30000000                54.3 ns/op
BenchmarkBigPalindromeFailWithStrings-8 20000000                88.0 ns/op
BenchmarkBigPalindromeFailWithMath-8    30000000                59.1 ns/op
ok      github.com/jcbwlkr/studies/intreverse   10.429s
```
