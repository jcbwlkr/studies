package main

import "fmt"

func main() {

	var s interface{} = "hello"

	var key []byte
	switch val := s.(type) {
	case string:
		fmt.Println("It's a string")
		key = []byte(val)
	case []byte:
		fmt.Println("It's a []byte")
		key = val
	default:
		fmt.Println("It's something")
		// TODO return error
	}

	fmt.Println(key)
}
