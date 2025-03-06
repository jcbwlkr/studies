package main

import "fmt"

type document struct {
	name  string
	count int
}

func main() {

	docs := []document{
		{name: "a", count: 10},
		{name: "b", count: 24},
		{name: "c", count: 15},
		{name: "d", count: 33},
		{name: "e", count: 18},
	}

	limit := 50

	total := 0
	for i, d := range docs {
		if total+d.count > limit {
			docs = docs[:i]
			break
		}
		total += d.count
	}

	for i, d := range docs {
		fmt.Println(i, d)
	}

}
