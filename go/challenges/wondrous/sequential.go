package main

func maxSequential(n int) result {
	m := 0
	for i := 1; i <= n; i++ {
		got := wondrousLength(i)
		if got > m {
			m = i
		}
	}
	return wondrous(m)
}
