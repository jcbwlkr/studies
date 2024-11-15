package sprintf

import (
	"fmt"
	"strconv"
	"testing"
)

func TestInts(t *testing.T) {
	a := fmt.Sprintf("%d", 123456789)
	b := strconv.Itoa(123456789)
	if a != b {
		t.Errorf("fmt.Sprintf: %q != strconv.Itoa: %q", a, b)
	}
}

var result string

func BenchmarkIntSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = fmt.Sprintf("%d", 123456789)
	}
}

func BenchmarkIntStrconv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result = strconv.Itoa(123456789)
	}
}
