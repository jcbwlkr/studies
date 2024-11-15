package main

import "fmt"

func main() {
	body := []int{1, 2, 3}

	fmt.Println("At the beginning body is", body)

	for i := 0; i < 5; i++ {
		pre := []int{i, i, i}

		pre = append(pre, body...)

		fmt.Println("Pass:", i)
		fmt.Println("pre is", pre)
	}

	fmt.Println("At the end body is", body)
}
