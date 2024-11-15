package main

import "testing"

var result bool

func BenchmarkTooShort(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = "bb" == secret
	}
	result = r
}

func BenchmarkTooLong(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = "badcababcde" == secret
	}
	result = r
}

func BenchmarkRightLength0of6(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = "xxxxxx" == secret
	}
	result = r
}

func BenchmarkRightLength1of6(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = "bxxxxx" == secret
	}
	result = r
}

func BenchmarkRightLength2of6(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = "baxxxx" == secret
	}
	result = r
}

func BenchmarkRightLength3of6(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = "badxxx" == secret
	}
	result = r
}

func BenchmarkRightLength4of6(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = "badcxx" == secret
	}
	result = r
}

func BenchmarkRightLength5of6(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = "badcax" == secret
	}
	result = r
}

func BenchmarkCorrect(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = "badcab" == secret
	}
	result = r
}

func BenchmarkRightLength5of6Backwards(b *testing.B) {
	var r bool
	for i := 0; i < b.N; i++ {
		r = "xadcab" == secret
	}
	result = r
}
