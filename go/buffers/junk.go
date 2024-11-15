package main

import (
	"bytes"
	"fmt"
	"os"
	"time"
)

func main() {
	iterations := 100000
	t0 := time.Now()

	s := ""
	for i := 0; i < iterations; i++ {
		s += "a"
	}
	fmt.Fprint(os.Stderr, s)

	fmt.Println("Buffer build up took", time.Since(t0))

	t1 := time.Now()

	var b bytes.Buffer
	for i := 0; i < iterations; i++ {
		b.WriteString("b")
	}
	fmt.Fprint(os.Stderr, b.String())

	fmt.Println("Buffer build up took", time.Since(t1))
}
