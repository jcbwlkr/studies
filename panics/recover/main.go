package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Operation |", "Result")
	fmt.Println("----------------------")
	fmt.Println("pass      |", myFunc("pass"))
	fmt.Println("fail      |", myFunc("fail"))
	fmt.Println("panic     |", myFunc("panic"))
	fmt.Println("assert    |", myFunc("assert"))

	useMyRecover()

	fmt.Println("All done!")
}

func myFunc(op string) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("myFunc panicked: %v", r)
		}
	}()

	switch op {
	case "fail":
		return errors.New("you failed normally")
	case "panic":
		panic("OH NO! DISASTER!")
	case "assert":
		var i interface{}
		i = "hello"
		d := i.(int)
		fmt.Println(d)
	}

	return nil
}

func useMyRecover() {
	defer myRecover()

	fmt.Println("About to die")
	panic("UH OH")
}

func myRecover() {
	if err := recover(); err != nil {
		fmt.Println("Recovered from panic in my special function:", err)
	}
}
