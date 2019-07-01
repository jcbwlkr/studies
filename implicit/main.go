package main

import "fmt"

type number int

func useAsNumber(n number) {
	fmt.Println(n)
}

//func willNotCompile() {
//n := 2
//useAsNumber(n)
//}

type params map[string]interface{}

func useAsParams(p params) {
	fmt.Println(p)
}

func compiles() {
	m := map[string]interface{}{
		"Foo": "Bar",
		"Baz": 123,
	}
	useAsParams(m)
}

type user struct {
	id   int
	name string
}

func processUser(u user) {
	fmt.Println(u)
}

func compilesUser() {

	var u struct {
		id   int
		name string
	}

	processUser(u)
}
