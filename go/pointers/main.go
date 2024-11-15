package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := new(int)
	*x = 8

	y := new(int)
	*y = 8

	fmt.Println("x == y", x == y)
	fmt.Println("reflect.DeepEqual(x, y)", reflect.DeepEqual(x, y))
}
