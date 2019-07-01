package main

import (
	"fmt"
	"sort"
)

// This code comes from the documentation of sort.Search
func main() {
	fmt.Println("Pick an integer from 0 to 100.")

	var s string
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})

	fmt.Printf("Your number is %d.\n", answer)
}
