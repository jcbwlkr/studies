package main

import "fmt"

func main() {

	// Original map
	m := map[string]string{
		"jacob": "Jacob Walker",
		"jenn":  "Jenn Walker",
	}
	fmt.Println(m)

	// This only copies the map header so deleting from `x` affects `m`
	x := m
	delete(x, "jacob")
	fmt.Println(m)

	// This makes a whole new map and copies the values so modifying `y` does not affect `m`
	y := make(map[string]string)
	for k, v := range m {
		y[k] = v
	}
	delete(y, "jenn")
	fmt.Println(m)
}
