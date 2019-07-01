package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("runtime.NumCPU() =", runtime.NumCPU())
	fmt.Println("runtime.GOMAXPROCS(-1) =", runtime.GOMAXPROCS(-1))
}
