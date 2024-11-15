package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

func main() {
	size := 10000000

	var wg sync.WaitGroup
	wg.Add(2)

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	a := strings.Repeat("A", size)
	b := strings.Repeat("B", size)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Fprintln(f, a)
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Fprintln(f, b)
		}
		wg.Done()
	}()

	wg.Wait()
}
