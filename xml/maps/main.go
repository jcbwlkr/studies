package main

import (
	"encoding/xml"
	"fmt"
	"log"
)

type Doc struct {
	Row map[string]string
}

func main() {
	d := Doc{
		Row: map[string]string{
			"A": "Apple",
			"B": "Banana",
			"C": "Cherry",
		},
	}
	b, err := xml.Marshal(d)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
