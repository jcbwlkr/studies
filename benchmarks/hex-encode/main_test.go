package main

import (
	"encoding/hex"
	"fmt"
	"testing"
)

var result string

func BenchmarkHexEncoding(b *testing.B) {

	data := []byte{65, 66, 67, 68, 69, 70, 71, 72}

	b.Run("fmt.Sprintf", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := fmt.Sprintf("%x", data)
			result = s
		}
	})
	b.Run("hex.EncodeToString", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			s := hex.EncodeToString(data)
			result = s
		}
	})
}
