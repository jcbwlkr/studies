package main

import "fmt"

func main() {

	names := []string{
		"Larry", "Curly", "Moe", "Shep",
	}

	addrs := []*string{}
	for _, name := range names {
		addrs = append(addrs, &name)
	}

	for i, addr := range addrs {
		fmt.Printf("%2d: %v %q\n", i, addr, *addr)
	}
}
