package main

import "fmt"

func main() {
	var count int

	defer fmt.Println("Count A", count)

	defer func() {
		fmt.Println("In defer", count)
	}()

	fmt.Println("In main", count)
	count++
	fmt.Println("In main", count)
	count++
	fmt.Println("In main", count)
	count++
}
