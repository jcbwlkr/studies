package main

import (
	"fmt"
	"time"
)

func main() {
	d := time.Date(2017, time.November, 30, 0, 0, 0, 0, time.UTC)
	fmt.Println("Starting ==", d)
	for i := 0; i < 12; i++ {
		d = d.AddDate(0, 1, 0)
		fmt.Println("+1 month ==", d)
	}
}
