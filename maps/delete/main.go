package main

import "fmt"

func main() {

	m := map[string]string{
		"jcbwlkr": "Jacob Walker",
		"anna":    "Anna Walker",
	}

	delete(m, "jcbwlkr")
	delete(m, "rory")

	fmt.Println(m)
}
