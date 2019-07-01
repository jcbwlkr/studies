package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano()) // Try changing this number!

	quote_file, err := os.Open("quotes.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(quote_file)
	scanner.Split(onDollarSign)

	var quotes []string
	// I think this will scan the file and append all the parsed quotes into quotes
	for scanner.Scan() {
		quotes = append(quotes, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}

	fmt.Println("There are", len(quotes), "quotes")
	fmt.Println("quote 1:", quotes[rand.Intn(len(quotes))])
	fmt.Println("quote 2:", quotes[rand.Intn(len(quotes))])
	fmt.Println("quote 3:", quotes[rand.Intn(len(quotes))])
}

// define split function
func onDollarSign(data []byte, atEOF bool) (advance int, token []byte, err error) {

	// If we are at the end of the file and there's no more data then we're done
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// If we are at the end of the file and there IS more data return it
	if atEOF {
		return len(data), data, nil
	}

	// If we find a $ then check if the next rune after is also a $. If so we
	// want to advance past the second $ and return a token up to but not
	// including the first $.
	if i := bytes.IndexByte(data, '$'); i >= 0 {
		if len(data) > i && data[i+1] == '$' {
			return i + 2, data[0:i], nil
		}
	}

	// Request more data.
	return 0, nil, nil
}
