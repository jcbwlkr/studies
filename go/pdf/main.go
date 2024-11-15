package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	m := flag.String("messages", "", "Path to the messages file")

	flag.Parse()

	if *m == "" {
		log.Fatal("-messages is required")
	}

	fmt.Println("You said " + *m)
}
