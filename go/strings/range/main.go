package main

import "fmt"

func main() {

	name := "zoë"
	for i, r := range name {
		fmt.Printf("%d %c\n", i, r)
	}
}
