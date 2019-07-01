// +build group

package main

import (
	"errors"
	"fmt"
)

type ErrorGroup struct {
	funcs  []func() error
	Errors []error
}

func NewErrorGroup(f ...func() error) *ErrorGroup {
	return &ErrorGroup{
		funcs: f,
	}
}

func (e *ErrorGroup) Run() {
	for _, fn := range e.funcs {
		err := fn()
		if err != nil {
			e.Errors = append(e.Errors, err)
		}
	}
}

func main() {

	g := NewErrorGroup(
		func() error { return doSomething("foo") },
		func() error { return doSomething("bar") },
		func() error { return doSomething("baz") },
		func() error { return doSomething("qux") },
	)

	g.Run()

	fmt.Println(g.Errors)
}

func doSomething(x string) error {
	fmt.Println("doing x: ", x)
	if x == "bar" {
		return errors.New("OH NO it's Bar")
	}
	return nil
}
