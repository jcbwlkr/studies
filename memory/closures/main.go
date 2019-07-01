package main

import "fmt"

func main() {

	x := 42

	fmt.Printf("main:\t&x = %p  x = %d\n", &x, x)

	f := func() {
		fmt.Printf("clsr:\t&x = %p  x = %d\n", &x, x)
	}

	f()

	x = 84

	f()

	fmt.Printf("main:\t&x = %p  x = %d\n", &x, x)
}
