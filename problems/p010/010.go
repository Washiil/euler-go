package p010

import (
	"math"

	"github.com/washiil/euler-go/problems/registry"
)

// Find the sum of all the primes below two million.

func checkPrime(n int) bool {
	if n < 2 {
		return false
	}
	limit := int(math.Floor(math.Sqrt(float64(n))))

	for i := 2; i <= limit; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func Solve010() int {
	sum := 0

	for i := 2; i < 2_000_000; i++ {
		if checkPrime(i) {
			sum += i
		}
	}

	return sum
}

func init() {
	registry.Register("10", Solve010)
}
