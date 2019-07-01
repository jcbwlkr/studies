package main

import (
	"fmt"
	"sync"
)

// User is a blah blah. It should not be copied.
type User struct {
	ID   string
	Name string

	*nocopy
	*sync.Mutex
}

type nocopy struct{}

func (*nocopy) Lock()   {}
func (*nocopy) Unlock() {}

func main() {
	u := User{ID: "abc", Name: "Jacob"}

	process(u)
}

func process(u User) {
	fmt.Println(u)
}
