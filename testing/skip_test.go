package main

import "testing"

func TestSkip(t *testing.T) {
	t.Skip("Nope")
	t.Log("Happens?")
}
