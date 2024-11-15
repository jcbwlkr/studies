package main

import "fmt"

func main() {
	fibUpTo(0)
	fibUpTo(1)
	fibUpTo(4)
	fibUpTo(8)
}

func fibUpTo(input int) {

	fmt.Println("====== Printing fib sequence up to and including", input)
	first := 0
	second := 1

	fmt.Printf("\n%d, %d", first, second)

	for first+second <= input {
		answer := first + second
		fmt.Printf(", %d", answer)

		first = second
		second = answer
	}

	fmt.Print("\n\n")
}
