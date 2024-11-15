package foo_test

import (
	"testing"

	"github.com/jcbwlkr/studies/testing/blackbox/foo"
)

func TestSum(t *testing.T) {
	got := foo.Sum(2, 4)
	want := 6
	if got != want {
		t.Errorf("sum(2, 4) = %v, want %v", got, want)
	}
}
