package main

import (
	"fmt"
	"os"
	"testing"
)

func init() {
	fmt.Println("func init in test file")
}

func TestMain(m *testing.M) {
	fmt.Println("in TestMain")
	code := m.Run()
	os.Exit(code)
}

func TestThing(t *testing.T) {
	t.Log("A test")
}
