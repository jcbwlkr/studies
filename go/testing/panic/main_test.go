package main

import (
	"fmt"
	"testing"
)

func TestFatal(t *testing.T) {
	fmt.Println("Begin")
	defer fmt.Println("End")

	t.Fatal("Die")
}

func TestPanic(t *testing.T) {
	fmt.Println("Begin")
	defer fmt.Println("End")

	panic("Die")
}
