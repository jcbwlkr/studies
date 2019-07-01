package main

import "fmt"

type User struct {
	ID   string `json:"-"`
	Name string `json:"name"`
}

type UserTest struct {
	User

	ID string `json:"id"`
}

func main() {
	a := User{}

	b := UserTest(a)

	fmt.Println(b)
}
