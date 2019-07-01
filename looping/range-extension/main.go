package main

import "fmt"

func main() {

	var x []int
	for i := 0; i < 10; i++ {
		x = append(x, i+1)
	}

	//for i, v := range x {
	for i := 0; i < len(x); i++ {
		v := x[i]
		if i < 5 {
			x = append(x, (i+1)*100)
		}
		fmt.Println(i, v)
	}

	fmt.Println("done:", x)
}
