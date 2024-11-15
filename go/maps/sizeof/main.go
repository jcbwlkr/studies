package main

import (
	"fmt"
	"unsafe"
)

func main() {
	m := make(map[rune]bool)
	fmt.Println(unsafe.Sizeof(m))
	m['a'] = true
	m['b'] = true
	m['c'] = true
	m['d'] = true
	fmt.Println(unsafe.Sizeof(m))
}
