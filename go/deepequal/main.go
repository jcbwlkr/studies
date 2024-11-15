package main

import (
	"fmt"
	"reflect"
)

func main() {

	a := map[string]interface{}{
		"foo":    "bar",
		"monkey": "banana",
		"age":    31,
	}

	b := map[string]interface{}{
		"age":    31,
		"monkey": "banana",
		"foo":    "bar",
	}

	fmt.Println(reflect.DeepEqual(a, b)) // true

	c := map[string]interface{}{
		"foo":   "bar",
		"names": []string{"larry", "curly", "moe", "shep"},
	}

	d := map[string]interface{}{
		"foo":   "bar",
		"names": []string{"larry", "curly", "moe", "shep"},
	}

	fmt.Println(reflect.DeepEqual(c, d)) // true

}
