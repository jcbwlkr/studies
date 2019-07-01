package intreverse

import "strconv"

func withStrings(n int) int {
	s := strconv.Itoa(n)
	rev := []byte(s)
	for i, j := 0, len(rev)-1; i < j; i, j = i+1, j-1 {
		rev[i], rev[j] = rev[j], rev[i]
	}
	a, _ := strconv.Atoi(string(rev))

	return a
}

func withMath(n int) int {
	m := 0
	for ; n > 0; n /= 10 { // Chop the last digit off n each pass
		m *= 10     // Make room for a new digit on n
		m += n % 10 // Add the last digit of n to m
	}
	return m
}

func isPalindromeString(n int) bool {
	s := strconv.Itoa(n)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}

	return true
}

func isPalindromeMath(n int) bool {
	m, orig := 0, n
	for ; n > 0; n /= 10 { // Chop the last digit off n each pass
		m *= 10     // Make room for a new digit on n
		m += n % 10 // Add the last digit of n to m
	}
	return m == orig
}
