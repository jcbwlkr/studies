package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	process(nil)
	var slice []int
	process(slice)
}

func process(v interface{}) {
	fmt.Printf("process(%#v)\n", v)

	b, err := json.Marshal(v)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", b)
}
