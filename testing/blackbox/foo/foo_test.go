package foo

import "testing"

func TestInternal_sum(t *testing.T) {
	got := sum(2, 4)
	want := 6
	if got != want {
		t.Errorf("sum(2, 4) = %v, want %v", got, want)
	}
}
