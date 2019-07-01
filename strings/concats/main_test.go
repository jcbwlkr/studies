package main

import "testing"

var gs string

func BenchmarkConcat(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = "foo" + "bar" + "baz" + "banana"
	}

	gs = s
}

func BenchmarkBytes(b *testing.B) {
	var s string
	for i := 0; i < b.N; i++ {
		s = "foo" + "bar" + "baz" + "banana"
	}

	gs = s
}
