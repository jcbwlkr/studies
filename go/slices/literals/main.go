package main

import "fmt"

type user struct {
	name string
}

func main() {
	slice := []*user{
		{"Anna"},
		{"Jacob"},
	}

	fmt.Println(slice)
}
