package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		fmt.Println("G1")

		work()

		fmt.Println("G2")
	}()

	wg.Wait()
	fmt.Println("Done")
}

func work() {
	fmt.Println("W1")
	defer fmt.Println("WD1")

	work2()

	defer fmt.Println("WD1")
	fmt.Println("W2")
}
func work2() {
	fmt.Println("WW1")
	defer fmt.Println("WWD1")

	runtime.Goexit()
}
