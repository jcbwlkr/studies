package main

import "fmt"

func main() {
	s := []byte("abcdefghi")
	fmt.Println(s, len(s), cap(s))
}
