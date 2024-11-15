package main

import "testing"

func setup(t *testing.T) {
	t.Log("reseed the database")
}

func TestSomething(t *testing.T) {
	setup(t)

	t.Log("Testing thing 1")
}

func TestOtherThing(t *testing.T) {
	setup(t)

	t.Log("Testing thing 2")
}

func TestWithSubTests(t *testing.T) {
	setup(t)

	t.Run("A=1", func(t *testing.T) {
		t.Log("Running sub test A=1")
	})
	t.Run("A=2", func(t *testing.T) {
		t.Log("Running sub test A=2")
	})
	t.Run("A=3", func(t *testing.T) {
		t.Log("Running sub test A=3")
	})
}
