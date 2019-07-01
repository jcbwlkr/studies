package main

import "fmt"

func main() {
	fmt.Println("vim-go")
	c := make(chan int, 8)

	c <- 1
	c <- 2
	fmt.Printf("len: %d, cap: %d\n", len(c), cap(c))
	c <- 3
	c <- 4
	fmt.Printf("len: %d, cap: %d\n", len(c), cap(c))
	c <- 4
	c <- 4
	fmt.Printf("len: %d, cap: %d\n", len(c), cap(c))
	c <- 4
	c <- 4
	fmt.Printf("len: %d, cap: %d\n", len(c), cap(c))
}
