package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var a, b string
	_, err := fmt.Scanln(&a, &b)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("a", a)
	fmt.Println("b", b)
}

func xmain() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)

	for s.Scan() {
		fmt.Println(s.Text())
	}

	if err := s.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done!")
}
