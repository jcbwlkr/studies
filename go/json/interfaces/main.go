package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	err := work()

	b, _ := json.Marshal(err)

	fmt.Println(string(b))
}

func work() error {
	return MyError{Foo: "f", Bar: "b"}
}

type MyError struct {
	Foo, Bar string
}

func (e MyError) Error() string {
	return "ERROR: " + e.Foo + " " + e.Bar
}
