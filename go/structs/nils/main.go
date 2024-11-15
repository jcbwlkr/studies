package main

import (
	"fmt"

	"github.com/fatih/structs"
)

type update struct {
	Name    *string
	Enabled *bool
	age     *int
}

func main() {
	var u update

	m := structs.Map(u)
	fmt.Printf("%+v\n", m)
}
