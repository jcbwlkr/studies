package main

import "C"
import "fmt"

//export Greet
func Greet(name string) {
	fmt.Println(Greeting(name))
}

//export Greeting
func Greeting(name string) string {
	return "Hello, " + name + "!"
}

func main() {
}
