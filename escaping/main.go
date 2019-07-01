package main

import "fmt"

func main() {
	stays(20)
	escapes(20)
	a := 22
	fmt.Println(a)
}

func stays(x int) int {
	return x + 22
}

func escapes(x interface{}) int {
	switch x := x.(type) {
	case int:
		return x + 22
	default:
		return 42
	}
}
