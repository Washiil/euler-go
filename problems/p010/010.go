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
	const limit = 2_000_000

	// First consider all numbers as prime
	isPrime := make([]bool, limit)
	for i := 2; i < limit; i++ {
		isPrime[i] = true
	}

	// Implement Sieve of Eratosthenes
	for i := 2; i < limit; i++ {
		if !isPrime[i] {
			continue
		}
		// Remove numbers who are multiples of prime as they cannot be prime
		for j := i * i; j < limit; j += i {
			isPrime[j] = false
		}
	}

	// Sum prime numbers
	sum := 0
	for i := 2; i < limit; i++ {
		if isPrime[i] {
			sum += i
		}
	}

	return sum
}

func init() {
	registry.Register("10", Solve010)
}
