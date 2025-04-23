package p021

import (
	"github.com/washiil/euler-go/problems/registry"
)

func sum_of_proper_divisors(n int) int {
	sum := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

func Solve021() int {
	output := 0
	visited := make(map[int]bool)

	for a := 1; a < 10_000; a++ {
		if visited[a] {
			continue
		}
		b := sum_of_proper_divisors(a)
		if a != b && sum_of_proper_divisors(b) == a {
			// A is guranteed to not be over 10_000 due to our loop
			// So only check b
			if b < 10_000 {
				output += a + b
			} else {
				output += a
			}
			// don't repeat numbers
			visited[a] = true
			visited[b] = true
		}
	}

	return output
}

func init() {
	registry.Register("21", Solve021)
}
