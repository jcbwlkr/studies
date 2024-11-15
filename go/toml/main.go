package main

import (
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

type bills struct {
	// TODO Stuff
}

func parseBills(r io.Reader) (bills, error) {
	// TODO Pass r to the toml parser
}

func main() {

	// In func main() we open the actual file and pass it to parseBills
	// We can do this because os.Open returns an *os.File which happens to
	// implement the io.Reader interface.
	f, err := os.Open("bills.toml")
	if err != nil {
		log.Fatal(err)
	}

	bills, err := parseBills(f)
	if err != nil {
		log.Fatal(err)
	}
}

func TestParseBills(t *testing.T) {

	// in tests we can use the strings package to make a reader which always yields this static value
	r := strings.NewReader(`minutes = 35.00
messages = 8.00
megabytes = 20.00
devices = 42.00
extras = 1.00
fees = 12.84`)

	bills, err := parseBills(r)
	if err != nil {
		t.Fatalf("parseBills should be able to parse this input, got %v", err)
	}
}
