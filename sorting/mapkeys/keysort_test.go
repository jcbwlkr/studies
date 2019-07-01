package main

import (
	"reflect"
	"testing"
)

//BenchmarkTopKeys-8       1000000              1522 ns/op             440 B/op         14 allocs/op
//BenchmarkTopSlice-8      1000000              1289 ns/op             720 B/op         11 allocs/op

func BenchmarkTopKeys(b *testing.B) {
	for i := 0; i < b.N; i++ {
		top(3, data)
	}
}

func BenchmarkTopSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		topSlice(3, data)
	}
}

func BenchmarkTopSliceMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		topSliceMap(3, data)
	}
}

var (
	data = map[string]int{
		"apple":      1234,
		"banana":     234,
		"cherry":     999,
		"date":       423,
		"elderberry": 1234,
	}

	want    = []string{"apple", "elderberry", "cherry"}
	wantAlt = []string{"elderberry", "apple", "cherry"}
)

func TestTop(t *testing.T) {
	// Don't panic if we have fewer values than requested
	_ = top(33, data)
	got := top(3, data)

	if !reflect.DeepEqual(got, want) && !reflect.DeepEqual(got, wantAlt) {
		t.Errorf("top(%d, %v) failed", 3, data)
		t.Errorf("got %v", got)
		t.Errorf("want %v", want)
		t.Errorf("or want %v", wantAlt)
	}
}

func TestTopSlice(t *testing.T) {
	// Don't panic if we have fewer values than requested
	_ = topSlice(33, data)
	got := topSlice(3, data)

	if !reflect.DeepEqual(got, want) && !reflect.DeepEqual(got, wantAlt) {
		t.Errorf("topSlice(%d, %v) failed", 3, data)
		t.Errorf("got %v", got)
		t.Errorf("want %v", want)
		t.Errorf("or want %v", wantAlt)
	}
}

func TestTopSliceMap(t *testing.T) {
	// Don't panic if we have fewer values than requested
	_ = topSliceMap(33, data)
	got := topSliceMap(3, data)

	if !reflect.DeepEqual(got, want) && !reflect.DeepEqual(got, wantAlt) {
		t.Errorf("topSliceMap(%d, %v) failed", 3, data)
		t.Errorf("got %v", got)
		t.Errorf("want %v", want)
		t.Errorf("or want %v", wantAlt)
	}
}
