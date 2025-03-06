package main

import (
	"fmt"
	"reflect"
)

func main() {
	username := "jcbwlkr"

	v := reflect.ValueOf(&username)
	newName := "jenn"
	v.SetString(newName)

	fmt.Println(username)

}
