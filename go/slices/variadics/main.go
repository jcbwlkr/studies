package main

import "fmt"

func main() {

	fmt.Println("Separate args")
	fn("a", "b", "c")

	names := []string{"Anna", "Jacob", "Kell", "Carter", "Rory"}

	fmt.Println("Slice expansion")
	fn(names...)

	fmt.Println("Subslice expansion")
	fn(names[0:3]...)
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
