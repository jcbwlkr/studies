package main

import "testing"

func TestError(t *testing.T) {
	t.Log("Start")
	defer t.Log("Defer after error")

	t.Error("Just error")

}

func TestFatal(t *testing.T) {
	t.Log("Start")
	defer t.Log("Defer after fatal")

	t.Fatal("Die!")
}
