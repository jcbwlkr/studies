package main

import "fmt"

func nothing() {
}

func add(a, b int) int {
	return a + b
}

func main() {
	nothing()

	x := add(1, 2)
	fmt.Println(x)
}
