package p023

import (
	"github.com/washiil/euler-go/problems/registry"
)

func sumOfDivisors(n int) int {
	sum := 0
	for i := 1; i < n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func isAbundant(n int) bool {
	return sumOfDivisors(n) > n
}

func Solve023() int {
	output := 0
	abundantNums := make([]int, 100)

	for i := 1; i < 28123; i++ {
		if isAbundant(i) {
			abundantNums = append(abundantNums, i)
		}
	}

	return output
}

func init() {
	registry.Register("023", Solve023)
}
