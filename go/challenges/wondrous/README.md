# Wondrous numbers

Benchmarks here are for finding the longest sequence of wondrous numbers for
every n in the range of 1 - 1000.

Sequential implementation for `max`: **669k ns/op**

```
$ go test -bench .
BenchmarkWondrous1000-8          2000000               781 ns/op
BenchmarkMax1000-8                  2000            669514 ns/op
PASS
ok      github.com/jcbwlkr/studies/challenges/wondrous  3.781s
```

Initial naive concurrent `max`: **846k ns/op**

```
$ go test -bench .
BenchmarkWondrous1000-8          2000000               789 ns/op
BenchmarkMax1000-8                  2000            846068 ns/op
PASS
ok      github.com/jcbwlkr/studies/challenges/wondrous  4.178s
```

A worker implementation that uses one goroutine per cpu to do the calculations,
another to feed the workers, while main collates the results then waits for the
workers to finish: **504k ns/op**

I added the `-benchtime` flag to get a larger sample size to even out some
spiky results I was seeing.

```
$ go test -benchtime 30s -bench ".*Max"
BenchmarkMax1000-8        100000            504246 ns/op
PASS
ok      github.com/jcbwlkr/studies/challenges/wondrous  55.754s
```

My conclusion is although I *thought* this would be a good application for
concurrency since each test is independent of the others, it really isn't worth
the complexity. If it took a lot longer to do the individual calculations or if
each worker might get blocked on some external resource then maybe concurrency
would make more sense.
