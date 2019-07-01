package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("vim-go")

	f := NewFoo("2", 2)
	fmt.Println(f)
}

type Foo struct {
	id  string
	num int
}

func NewFoo(id string, num int) Foo {
	return Foo{id: id, num: num}
}

func ids(n int) (string, int) {
	return strconv.Itoa(n), n
}
