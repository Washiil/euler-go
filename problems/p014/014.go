package p014

import (
	"github.com/washiil/euler-go/problems/registry"
)

func Solve014() int {
	maximumLength := -1
	largestValue := -1

	collatzSequences := make(map[int]int)

	for i := 1; i < 1_000_000; i++ {
		sequenceLength := 1
		val, ok := collatzSequences[i]

		if ok {
			sequenceLength += val
		} else {
			curr := i
			for curr != 1 {
				// Check if we already know the sequence length from this point
				if val, ok := collatzSequences[curr]; ok && curr != i {
					sequenceLength += val - 1
					break
				}

				if curr%2 == 0 {
					curr = curr / 2
				} else {
					curr = (3 * curr) + 1
				}
				sequenceLength += 1
			}
			collatzSequences[i] = sequenceLength
		}
		if sequenceLength > maximumLength {
			maximumLength = sequenceLength
			largestValue = i
		}
	}

	return largestValue
}

func init() {
	registry.Register("14", Solve014)
}
