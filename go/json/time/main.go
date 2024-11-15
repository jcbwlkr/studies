package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"
)

func main() {
	t := time.Now()

	b, err := json.Marshal(t)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}
