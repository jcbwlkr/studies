package main

import (
	"fmt"
)

// nocopy does not do anything by itself but including it in another struct
// tells "go vet" that a type should not be copied.
type nocopy struct{}

func (t *nocopy) Lock()   {}
func (t *nocopy) Unlock() {}

// user represents something unique. It contains field of type nocopy so "go
// vet" knows we should never copy one of these records.
type user struct {
	Name string
	ID   int

	_ nocopy // include the nocopy field but name it _ so it is not accessible and its methods are not promoted
}

func main() {

	users := make([]user, 10)

	for _, u := range users { // range var u copies lock: main.user contains main.nocopy
		fmt.Println(u) // call of fmt.Println copies lock value: main.user contains main.nocopy
	}

	var u user
	work(u) // call of work copies lock value: main.user contains main.nocopy
}

func work(u user) { // work passes lock by value: main.user contains main.nocopy
	fmt.Println("Hi")
}
