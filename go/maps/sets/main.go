package main

import "fmt"

func main() {

	seen := make(map[string]struct{})

	input := []string{
		"bar", "bar", "bar", "barbara", "ann",
	}

	for _, thing := range input {
		if _, ok := seen[thing]; !ok {
			fmt.Println("New thing", thing)
			seen[thing] = struct{}{}
		}
	}

}
