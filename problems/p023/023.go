package p023

import (
	"github.com/washiil/euler-go/problems/registry"
)

func sumOfDivisors(n int) int {
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum
}

func isAbundant(n int) bool {
	return sumOfDivisors(n) > n
}

const limit = 28123

func Solve023() int {
	output := 0

	var abundantNumbers []int
	for i := 12; i <= limit; i++ {
		if isAbundant(i) {
			abundantNumbers = append(abundantNumbers, i)
		}
	}

	canBeWrittenAsAbundantSum := make([]bool, limit+1)
	for i := 0; i < len(abundantNumbers); i++ {
		for j := i; j < len(abundantNumbers); j++ {
			sum := abundantNumbers[i] + abundantNumbers[j]
			if sum <= limit {
				canBeWrittenAsAbundantSum[sum] = true
			}
		}
	}

	for i := 1; i <= limit; i++ {
		if !canBeWrittenAsAbundantSum[i] {
			output += i
		}
	}

	return output
}

func init() {
	registry.Register("23", Solve023)
}
