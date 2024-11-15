package main

import "fmt"

// Could be defined in another file
var example = `{
	"foo": "bar",
	"baz": [
		"apple",
		"banana",
		"cherries"
	]
}`

func main() {
	// When they init a project write the example config to a file
	fmt.Println("example config")
	fmt.Println(example)
}
