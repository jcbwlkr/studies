package main

import "fmt"

type span struct {
	id string
}

func (sp *span) print() {
	fmt.Println("Span ID is", sp.id)
}

func main() {

	var sp *span

	sp = &span{id: "a"}
	defer sp.print()

	sp = &span{id: "b"}
	defer sp.print()

	sp = &span{id: "c"}
	defer sp.print()

	fmt.Println("pointer defer test")
}
