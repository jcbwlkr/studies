package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	const input = `[[1, 2, [3, 4, 5, [6, 7, 8], 9], 10, 11, [12, 13, 14]], 15]`

	var data interface{}
	err := json.Unmarshal([]byte(input), &data)
	if err != nil {
		log.Fatal(err)
	}

	out := process(data, nil)
	fmt.Println(out)
}

func process(in interface{}, nums []float64) []float64 {
	switch val := in.(type) {
	case float64:
		nums = append(nums, val)
	case []interface{}:
		for _, v := range val {
			nums = process(v, nums)
		}
	default:
		log.Panicf("unknown value in input: %T", val)
	}

	return nums
}
