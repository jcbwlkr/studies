package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	fmt.Println("Num CPUS:", runtime.NumCPU())

	for {
		select {
		case <-take2():
			fmt.Println("Select 2")
		case <-take4():
			fmt.Println("Select 4")
		}
	}

}

func take4() chan int {
	fmt.Println("In take4")
	time.Sleep(4 * time.Second)
	fmt.Println("Leaving take4")

	c := make(chan int, 2)
	c <- 4
	return c
}

func take2() chan int {
	fmt.Println("In take2")
	time.Sleep(2 * time.Second)
	fmt.Println("Leaving take2")

	c := make(chan int, 2)
	c <- 2
	return c
}
