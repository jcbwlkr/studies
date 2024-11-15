package intreverse

import "testing"

var tests = []struct {
	in  int
	out int
	pal bool
}{
	{123456, 654321, false},  // even, not palindrome
	{432234, 432234, true},   // even, palindrome
	{12345, 54321, false},    // odd, not palindrome
	{4321234, 4321234, true}, // odd, palindrome
}

func TestReverseWithString(t *testing.T) {
	for i, test := range tests {
		actual := withStrings(test.in)
		if actual != test.out {
			t.Errorf("Test %d: withStrings(%d) = %d, expected %d", i, test.in, actual, test.out)
		}
	}
}

func TestReverseWithMath(t *testing.T) {
	for i, test := range tests {
		actual := withMath(test.in)
		if actual != test.out {
			t.Errorf("Test %d: withMath(%d) = %d, expected %d", i, test.in, actual, test.out)
		}
	}
}

func TestIsPalindromeWithStrings(t *testing.T) {
	for i, test := range tests {
		actual := isPalindromeString(test.in)
		if actual != test.pal {
			t.Errorf("Test %d: isPalindromeString(%d) = %v, expected %v", i, test.in, actual, test.pal)
		}
	}
}

func TestIsPalindromeWithMath(t *testing.T) {
	for i, test := range tests {
		actual := isPalindromeMath(test.in)
		if actual != test.pal {
			t.Errorf("Test %d: isPalindromeMath(%d) = %v, expected %v", i, test.in, actual, test.pal)
		}
	}
}

func BenchmarkReverseWithStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withStrings(123456789)
	}
}

func BenchmarkReverseWithMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		withMath(123456789)
	}
}

func BenchmarkBigPalindromePassWithStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindromeString(98765432123456789)
	}
}

func BenchmarkBigPalindromePassWithMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindromeMath(98765432123456789)
	}
}

func BenchmarkBigPalindromeFailWithStrings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindromeString(1234567890123456789)
	}
}

func BenchmarkBigPalindromeFailWithMath(b *testing.B) {
	for i := 0; i < b.N; i++ {
		isPalindromeMath(1234567890123456789)
	}
}
