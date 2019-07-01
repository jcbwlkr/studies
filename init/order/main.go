package main

import "fmt"
import _ "./pkg"

var n, _ = fmt.Println("in var")

func init() {
	fmt.Println("in init")
}

func main() {
	fmt.Println("in main")
}
