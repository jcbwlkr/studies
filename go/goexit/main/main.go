package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Starting up")
	go work()

	fmt.Println("Killing main")
	runtime.Goexit()
}

func work() {
	for i := 0; i < 10; i++ {
		fmt.Println("Tick")
		time.Sleep(100 * time.Millisecond)
	}

	fmt.Println("Done with work")
}
