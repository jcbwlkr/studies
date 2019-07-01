package main

import (
	"fmt"
	"reflect"
)

func main() {
	username := "jcbwlkr"

	v := reflect.ValueOf(&username)
	newName := "anna"
	v.SetString(newName)

	fmt.Println(username)

}
