package scanlines

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func ExampleScanlines() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line is %[1]T %[1]q\n", line)
	}
	if err := scanner.Err(); err != nil {
		log.Println("reading standard input:", err)
	}

	// Output:
	// line is string "foo"
	// line is string "bar"
	// line is string "baz"
	// line is string "banana"
}
