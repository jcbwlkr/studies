package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Call this with the maximum number to inspect like\n\twondrous 42")
		os.Exit(1)
	}

	n, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Println(err)
		os.Exit(2)
	}

	fmt.Println(maxSequential(n))
}

type result struct {
	n   int
	seq []int
}

func wondrous(n int) result {
	r := result{n: n}
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		r.seq = append(r.seq, n)
	}

	return r
}

func wondrousLength(n int) int {
	l := 0
	for n != 1 {
		if n%2 == 0 {
			n /= 2
		} else {
			n = n*3 + 1
		}
		l++
	}

	return l
}
