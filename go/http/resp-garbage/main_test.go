package main

import (
	"net/http/httptest"
	"testing"
)

func BenchmarkWrite(b *testing.B) {
	a := NewApp()
	for i := 0; i < b.N; i++ {
		r := httptest.NewRequest("GET", "/write", nil)
		w := httptest.NewRecorder()
		a.ServeHTTP(w, r)
	}
}

func BenchmarkWriteString(b *testing.B) {
	a := NewApp()
	for i := 0; i < b.N; i++ {
		r := httptest.NewRequest("GET", "/writestring", nil)
		w := httptest.NewRecorder()
		a.ServeHTTP(w, r)
	}
}

func BenchmarkStringtobytes(b *testing.B) {
	a := NewApp()
	for i := 0; i < b.N; i++ {
		r := httptest.NewRequest("GET", "/stringtobytes", nil)
		w := httptest.NewRecorder()
		a.ServeHTTP(w, r)
	}
}

func BenchmarkBytestostring(b *testing.B) {
	a := NewApp()
	for i := 0; i < b.N; i++ {
		r := httptest.NewRequest("GET", "/bytestostring", nil)
		w := httptest.NewRecorder()
		a.ServeHTTP(w, r)
	}
}
