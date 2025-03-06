package main

import "fmt"

type user struct {
	name string
}

func main() {
	slice := []*user{
		{"Jenn"},
		{"Jacob"},
	}

	fmt.Println(slice)
}
