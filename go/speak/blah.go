package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

var count = pflag.IntP("count", "c", 0, "number of times to talk")

func main() {

	pflag.Parse()

	for i := 0; i < *count; i++ {
		fmt.Println("hi")
	}
}
