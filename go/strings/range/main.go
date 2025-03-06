package main

import "fmt"

func main() {

	name := "zoë is here"
	for i, r := range name {
		fmt.Printf("%d %c\n", i, r)
	}
}
