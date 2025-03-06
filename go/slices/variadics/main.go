package main

import "fmt"

func main() {

	fmt.Println("Separate args")
	fn("a", "b", "c")

	names := []string{"Jenn", "Jacob", "Kell", "Krylex", "Carter", "Rory", "Olive", "Nixie"}

	fmt.Println("Slice expansion")
	fn(names...)

	fmt.Println("Subslice expansion")
	fn(names[0:3]...)

	fmt.Println("Empty")
	var empty []string
	fn(empty...)
}

func fn(s ...string) {

	fmt.Printf("Len(%d) Cap(%d)\n", len(s), cap(s))
	for i, s := range s {
		fmt.Println(i, s)
	}

	if len(s) != cap(s) {
		fmt.Println("GROWTH HACK!")
		s = s[0:cap(s)]
		fmt.Println(s)
	}
}
