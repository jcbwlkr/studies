package main

import "fmt"

type User struct {
	Name string
}

func main() {

	const field = "Name"

	u := User{field: "Jacob"}

	fmt.Println(u)
}
