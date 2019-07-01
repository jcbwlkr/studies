package main

import "fmt"

func main() {
	for k, v := range mapFunc() {
		fmt.Println(k, "=>", v)
	}
}

func mapFunc() map[string]string {
	return map[string]string{
		"f": "Foo",
		"b": "Bar",
		"z": "Baz",
	}
}
