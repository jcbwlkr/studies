package main

import (
	"reflect"
	"testing"
)

var results = []result{
	{6, []int{3, 10, 5, 16, 8, 4, 2, 1}},
	{19, []int{58, 29, 88, 44, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1}},
}

func TestWondrous(t *testing.T) {
	for i, want := range results {
		got := wondrous(want.n)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("%d: wondrous(%d) got:\n%v\nwant:\n%v", i, want.n, got, want)
		}
	}
}

var lengths = []struct{ n, l int }{
	{6, 8},
	{19, 20},
}

func TestWondrousLengths(t *testing.T) {
	for i, want := range lengths {
		got := wondrousLength(want.n)
		if got != want.l {
			t.Errorf("%d: wondrousLength(%d) got:\n%v\nwant:\n%v", i, want.n, got, want)
		}
	}
}

var maxTests = []struct {
	n    int
	want result
}{
	{10, result{9, []int{28, 14, 7, 22, 11, 34, 17, 52, 26, 13, 40, 20, 10, 5, 16, 8, 4, 2, 1}}},
	{100, result{97, []int{292, 146, 73, 220, 110, 55, 166, 83, 250, 125, 376, 188, 94, 47, 142, 71, 214, 107, 322, 161, 484, 242, 121, 364, 182, 91, 274, 137, 412, 206, 103, 310, 155, 466, 233, 700, 350, 175, 526, 263, 790, 395, 1186, 593, 1780, 890, 445, 1336, 668, 334, 167, 502, 251, 754, 377, 1132, 566, 283, 850, 425, 1276, 638, 319, 958, 479, 1438, 719, 2158, 1079, 3238, 1619, 4858, 2429, 7288, 3644, 1822, 911, 2734, 1367, 4102, 2051, 6154, 3077, 9232, 4616, 2308, 1154, 577, 1732, 866, 433, 1300, 650, 325, 976, 488, 244, 122, 61, 184, 92, 46, 23, 70, 35, 106, 53, 160, 80, 40, 20, 10, 5, 16, 8, 4, 2, 1}}},
}

func TestMaxSequential(t *testing.T) {
	for i, test := range maxTests {
		got := maxSequential(test.n)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: max(%d) got:\n%v\nwant:\n%v", i, test.n, got, test.want)
		}
	}
}

func TestMaxConcurrent(t *testing.T) {
	for i, test := range maxTests {
		got := maxConcurrent(test.n)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: max(%d) got:\n%v\nwant:\n%v", i, test.n, got, test.want)
		}
	}
}

func TestMaxMark(t *testing.T) {
	for i, test := range maxTests {
		got := max_mark(test.n)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%d: max(%d) got:\n%v\nwant:\n%v", i, test.n, got, test.want)
		}
	}
}

func BenchmarkWondrous1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		wondrous(1000)
	}
}

func BenchmarkMaxSequential1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxSequential(1000000)
	}
}

func BenchmarkMaxConcurrent1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxConcurrent(1000000)
	}
}

func BenchmarkMaxMark1000000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		max_mark(1000000)
	}
}
