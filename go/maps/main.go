package main

import "fmt"

func main() {

	m := map[string]interface{}{
		"ed": "hi",
	}

	modify(m)

	fmt.Println(m)
}

func modify(m map[string]interface{}) {
	m["foo"] = "bar"
}
