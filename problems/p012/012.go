package p012

import (
	"math"

	"github.com/washiil/euler-go/problems/registry"
)

func nextTriangleNumber() func() int {
	sum := 0
	current := 1

	return func() int {
		sum += current
		current += 1
		return sum
	}
}

func numOfFactors(n int) int {
	sqrtN := int(math.Sqrt(float64(n)))
	count := 0 // Start with one for itself

	for i := 1; i <= sqrtN; i++ {
		if n%i == 0 {
			if i*i == n {
				count += 1 // weird perfect square edge case
			} else {
				count += 2 // i and n/i
			}
		}
	}

	return count
}

func Solve012() int {
	output := 0
	generator := nextTriangleNumber()
	factors := 0

	for factors < 500 {
		output = generator()
		factors = numOfFactors(output)
	}

	return output
}

func init() {
	registry.Register("12", Solve012)
}
