package main

import "testing"

var result bool

func BenchmarkStringEquals(b *testing.B) {
	type strBench struct {
		name string
		a    string
		b    string
	}

	benches := []strBench{
		{"big-same", "abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyz"},
		{"big-similar", "abcdefghijklmnopqrstuvwxyz", "abcdefghijklmnopqrstuvwxyX"},
		{"big-different", "abcdefghijklmnopqrstuvwxyz", "axxxxxxxxxxxxxxxxxxxxxxxxx"},
		{"big-lengths", "abcdefghijklmnopqrstuvwxyz", "a"},
	}

	for _, bench := range benches {
		b.Run(bench.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				x := bench.a == bench.b
				result = x
			}
		})
	}
}
