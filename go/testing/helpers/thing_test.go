package main

import (
	"testing"
)

func TestHelloWorld(t *testing.T) {
	n := 4
	isEven(t, n)

	n = 5
	isEven(t, n)
}

func isEven(t *testing.T, n int) {
	t.Helper()
	if n%2 != 0 {
		t.Errorf("argument %d is not even", n)
	}
}

func TestFoobar(t *testing.T) {
	tt := NewTestThing(t, "http://google.broken")
	body := tt.Get("index")
	body2 := tt.Get("fail")

	if len(body) == 0 {
		t.Errorf("got empty body")
	}
	if len(body2) == 0 {
		t.Errorf("got empty body 2")
	}
}

type TestThing struct {
	url string
	t   *testing.T
}

func NewTestThing(t *testing.T, url string) *TestThing {
	t.Helper()

	if url == "" {
		t.Fatal("Blank URL passed to NewTestThing")
	}

	return &TestThing{t: t}
}

func (tt *TestThing) Get(id string) []byte {
	tt.t.Helper()

	if id == "fail" {
		tt.t.Fatal("forcing an error")
	}

	return nil
}
