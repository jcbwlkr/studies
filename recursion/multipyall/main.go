package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
)

var input = []byte(`[1, 2, 3, [1, 2, 3, [1, 2, 3, [1, 2, 3]]]]`)

func main() {

	var nums interface{}
	if err := json.Unmarshal(input, &nums); err != nil {
		log.Fatal(err)
	}

	total, err := MultiplyAll(nums)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Total is", total)
}

func MultiplyAll(n interface{}) (float64, error) {

	switch v := n.(type) {
	case float64:
		// If n is a number then just return its value
		// I'm using float64 because that's the default format for numbers in json
		return v, nil
	case []interface{}:
		// If n is another slice then process each member
		var x float64
		for _, y := range v {
			z, err := MultiplyAll(y)
			if err != nil {
				return 0, err
			}
			x += z
		}
		return x, nil
	default:
		// If it's anything else then abort
		return 0, errors.New("Unknown thing!")
	}
}
