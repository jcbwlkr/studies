package main

import "fmt"

func main() {

	a := make([]byte, 4)

	for i := range a {
		a[i] = byte(i)
	}

	fmt.Printf("%v\n", a)
}
