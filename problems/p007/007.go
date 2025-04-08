package p007

import (
	"math"

	"github.com/washiil/euler-go/problems/registry"
)

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

func Solve007() int {
	count := 0
	num := 1

	for count < 10_001 {
		num++
		if checkPrime(num) {
			count += 1
		}
	}

	return num
}

func init() {
	registry.Register("7", Solve007)
}
