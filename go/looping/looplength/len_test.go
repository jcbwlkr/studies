package main

import "testing"

var (
	input []int
	lastX int
)

func init() {
	length := 10000
	input = make([]int, length)
	for i := 0; i < length; i++ {
		input[i] = i
	}
}

func BenchmarkLenInLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := 0
		for j := 0; j < len(input); j++ {
			x = input[j]
		}
		lastX = x
	}
}

func BenchmarkLenBeforeLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x := 0
		l := len(input)
		for j := 0; j < l; j++ {
			x = input[j]
		}
		lastX = x
	}
}
