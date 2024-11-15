package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6}

	for _, n := range(s) {
		fmt.Println(n)
	}
}
