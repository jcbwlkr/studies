package main

/*
$ go test -bench . -benchmem
goos: darwin
goarch: amd64
pkg: github.com/jcbwlkr/studies/regex/bench
BenchmarkSubmatches-8            1000000              1773 ns/op              32 B/op          1 allocs/op
BenchmarkMatchString-8           1000000              1677 ns/op               0 B/op          0 allocs/op
BenchmarkParser-8               50000000                29.6 ns/op             0 B/op          0 allocs/op
BenchmarkStrings-8              50000000                27.7 ns/op             0 B/op          0 allocs/op
BenchmarkTanner-8                 500000              3670 ns/op             128 B/op          3 allocs/op
BenchmarkBr0xen-8                2000000               836 ns/op              56 B/op          3 allocs/op
BenchmarkAaron-8                50000000                33.9 ns/op             0 B/op          0 allocs/op
*/

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"
)

var verRE = regexp.MustCompile(`^v[0-9]+:.+$`)

func verMatchString(raw string) int {

	if !verRE.MatchString(raw) {
		return 0
	}

	// We know it's v1234: and we just want the 1234. Find the index of the :
	// and pull out the number part.
	pos := strings.Index(raw, ":")
	version, err := strconv.Atoi(raw[1:pos])
	if err != nil {
		return 0
	}

	return version
}

var verRE2 = regexp.MustCompile(`^v([0-9]+):.+$`)

func verSubmatches(raw string) int {

	matches := verRE2.FindStringSubmatch(raw)
	if len(matches) != 2 {
		return 0
	}

	version, err := strconv.Atoi(matches[1])
	if err != nil {
		return 0
	}

	return version
}

func verParser(raw string) int {
	if len(raw) < 4 {
		return 0
	}

	if raw[0] != 'v' {
		return 0
	}

	// Find the position of the :
	var pos int

loop:
	for pos = 1; pos < len(raw); pos++ {
		r := raw[pos]
		switch {
		// Let pos advance
		case '0' <= r && r <= '9':
		// Done advancing
		case r == ':':
			break loop

		// Unrecognized character in version space
		default:
			return 0
		}
	}

	// If there was no numbers, no :, or nothing after the : that's an error
	if pos == 1 || pos == len(raw)-1 {
		return 0
	}

	ver, err := strconv.Atoi(raw[1:pos])
	if err != nil {
		return 0
	}
	return ver
}

func verStrings(raw string) int {

	// Must be at least v1:x
	if len(raw) < 4 {
		return 0
	}

	// Must start with a v
	if raw[0] != 'v' {
		return 0
	}

	// Find the :
	pos := strings.Index(raw, ":")

	// If there were no numbers, no :, or nothing after the : that's an error
	if pos <= 1 || pos == len(raw)-1 {
		return 0
	}

	// Parse after the v up to before the :
	ver, err := strconv.Atoi(raw[1:pos])
	if err != nil {
		return 0
	}
	return ver
}

var re = regexp.MustCompile("^v([0-9]+):[a-z]+$")

func SternRottenMan(s string) int {
	//re := regexp.MustCompile("^v([0-9]+):[a-z]+$")
	if !re.MatchString(s) {
		return 0
	}

	subMatches := re.FindSubmatch([]byte(s))
	numberString := string(subMatches[1])
	n, _ := strconv.Atoi(numberString)
	return n
}

func br0xen(s string) int {
	var v int
	idx := strings.Index(s, ":")
	if idx < 0 || idx == len(s)-1 {
		return 0
	}
	r := strings.NewReader(s)
	fmt.Fscanf(r, "v%d", &v)
	return v
}

func GetVersionAaron(s string) int {
	idx := -1
	for i, r := range s {
		if r == ':' {
			idx = i
			break
		}
	}
	if idx == -1 || s[0] != 'v' || idx == len(s)-1 {
		return 0
	}
	n, err := strconv.Atoi(s[1:idx])
	if err != nil {
		return 0
	}
	return n
}

func TestString(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{"normal", "v1234:blahblah", 1234},
		{"not versioned", "blahblah", 0},
		{"missing version", "v:foobar", 0},
		{"missing v", "1234:foobar", 0},
		{"missing colon", "v1234foobar", 0},
		{"missing after", "v1234:", 0},
		{"missing number", "v:foobar", 0},
		{"missing lots", "v:", 0},
		{"number is letters", "vabc:foobar", 0},
		{"blank", "", 0},
	}

	funcs := []struct {
		name string
		fn   func(string) int
	}{
		{"jacob-regexp-submatches", verSubmatches},
		{"jacob-regexp-match", verMatchString},
		{"jacob-parser", verParser},
		{"jacob-strings", verStrings},
		{"tanner's", SternRottenMan},
		{"brian's", br0xen},
		{"aaron's", GetVersionAaron},
	}

	for _, test := range tests {
		for _, fn := range funcs {
			got := fn.fn(test.in)
			if got != test.want {
				t.Errorf("%s(%q) = %d, want %d", fn.name, test.in, got, test.want)
			}
		}
	}
}

var result int

func BenchmarkSubmatches(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = verSubmatches("v1234567890:blahblahblahblahblahblahblahblahblahblah")
	}
	result = n
}

func BenchmarkMatchString(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = verMatchString("v1234567890:blahblahblahblahblahblahblahblahblahblah")
	}
	result = n
}

func BenchmarkParser(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = verParser("v1234567890:blahblahblahblahblahblahblahblahblahblah")
	}
	result = n
}

func BenchmarkStrings(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = verStrings("v1234567890:blahblahblahblahblahblahblahblahblahblah")
	}
	result = n
}

func BenchmarkTanner(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = SternRottenMan("v1234567890:blahblahblahblahblahblahblahblahblahblah")
	}
	result = n
}

func BenchmarkBr0xen(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = br0xen("v1234567890:blahblahblahblahblahblahblahblahblahblah")
	}
	result = n
}

func BenchmarkAaron(b *testing.B) {
	var n int
	for i := 0; i < b.N; i++ {
		n = GetVersionAaron("v1234567890:blahblahblahblahblahblahblahblahblahblah")
	}
	result = n
}
