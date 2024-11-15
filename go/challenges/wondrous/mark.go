package main

func max_mark(n int) result {
	maxSequence := []int{}
	maxLength := 0
	maxN := 0

	for i := 1; i <= n; i++ {
		seqLength := 0
		seqLength = hotpo_length(i, seqLength, true)

		if seqLength > maxLength {
			maxLength = seqLength
			maxN = i
		}
	}

	// Get the max_sequence
	maxSequence = hotpo_sequence(maxN, maxSequence)

	return result{
		n:   maxN,
		seq: maxSequence,
	}
}

func hotpo_length(n int, seq_length int, begin bool) int {
	if !begin && n == 1 {
		return seq_length
	} else if n%2 == 0 {
		n = n / 2
	} else {
		n = (n * 3) + 1
	}

	seq_length += 1
	return hotpo_length(n, seq_length, false)
}

func hotpo_sequence(n int, sequence []int) []int {
	if n == 1 {
		return sequence
	} else if n%2 == 0 {
		n = n / 2
	} else {
		n = (n * 3) + 1
	}

	sequence = append(sequence, n)
	return hotpo_sequence(n, sequence)
}
