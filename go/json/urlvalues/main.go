package main

import (
	"fmt"
	"log"

	"github.com/gorilla/schema"
)

type values map[string][]string

type person struct {
	Name string
	Kids []string
}

func main() {

	v := values{
		"Name": {"Jacob"},
		"Kids": {"Kell", "Carter", "Rory"},
	}

	var p person

	if err := schema.NewDecoder().Decode(&p, v); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%+v\n", p)
}
