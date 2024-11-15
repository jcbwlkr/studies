package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type node struct {
	ID    string
	Nodes []node
}

func main() {

	var program struct {
		Nodes []node
	}

	program.Nodes = addParts(program.Nodes)

	b, err := json.MarshalIndent(program, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(b))
}

func addParts(nodes []node) []node {
	part1 := node{ID: "part 1", Nodes: []node{}}
	part1.Nodes = addModules(part1.Nodes)
	nodes = append(nodes, part1)

	return nodes
}

func addModules(nodes []node) []node {
	module1 := node{ID: "module 1", Nodes: []node{}}
	nodes = append(nodes, module1)

	return nodes
}
