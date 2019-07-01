package main

import (
	"flag"
	"fmt"
	"runtime/debug"
)

func main() {
	defer func() {
		err := recover()
		if err == nil {
			return
		}

		fmt.Println("VVVVVVVVVVVVVVVVVVVVV")
		fmt.Println(err)
		debug.PrintStack()
		fmt.Println("^^^^^^^^^^^^^^^^^^^^^")
		panic(err)
	}()

	var a, b bool

	flag.BoolVar(&a, "a", true, "something")
	flag.BoolVar(&b, "a", true, "something")

	fmt.Println(a, b)
}
