package p002

import "github.com/washiil/euler-go/problems/registry"

// By considering the terms in the Fibonacci sequence whose values
// do not exceed four million, find the sum of the even-valued terms.

func Solve002() int {
	sum := 0
	a, b := 1, 2

	for b < 4_000_000 {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}

	return sum
}

func init() {
	registry.Register("2", Solve002)
}
