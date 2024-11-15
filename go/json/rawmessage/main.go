package main

import (
	"encoding/json"
	"fmt"
	"log"
)

const input = `[
	{
		"id": 1234,
		"name": "Jacob"
	},
	{
		"name": "Anna"
	},
	"just a string",
	1234
]`

func main() {

	var batch []json.RawMessage

	if err := json.Unmarshal([]byte(input), &batch); err != nil {
		log.Fatal(err)
	}

	fmt.Println(len(batch))
}
