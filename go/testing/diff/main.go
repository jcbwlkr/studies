package main

import (
	"fmt"
	"log"
	"time"

	"github.com/google/go-cmp/cmp"
)

func main() {
	a := time.Date(2019, 01, 01, 0, 0, 0, 0, time.UTC)

	b, err := time.Parse(time.RFC822Z, "01 Jan 19 00:00 +0000")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("a == b         ", a == b)
	fmt.Println("a.Equal(b)     ", a.Equal(b))
	fmt.Println("cmp.Equal(a, b)", cmp.Equal(a, b))
}
