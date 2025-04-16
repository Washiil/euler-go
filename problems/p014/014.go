package p014

import (
	"github.com/washiil/euler-go/problems/registry"
)

func Solve014() int {
	maximum_length := -1
	largest_value := -1

	collatz_sequences := make(map[int]int)

	for i := 1; i < 1_000_000; i++ {
		sequence_length := 1
		val, ok := collatz_sequences[i]

		if ok {
			sequence_length += val
		} else {
			curr := i
			for curr != 1 {
				// Check if we already know the sequence length from this point
				if val, ok := collatz_sequences[curr]; ok && curr != i {
					sequence_length += val - 1
					break
				}

				if curr%2 == 0 {
					curr = curr / 2
				} else {
					curr = (3 * curr) + 1
				}
				sequence_length += 1
			}
			collatz_sequences[i] = sequence_length
		}
		if sequence_length > maximum_length {
			maximum_length = sequence_length
			largest_value = i
		}
	}

	return largest_value
}

func init() {
	registry.Register("14", Solve014)
}
