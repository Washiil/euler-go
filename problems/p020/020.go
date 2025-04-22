package p020

import (
	"math/big"

	"github.com/washiil/euler-go/problems/registry"
)

func Solve020() int {
	output := 0

	// Compute large factorial
	largeFactorial := big.NewInt(1)
	for i := 100; i > 0; i-- {
		largeFactorial.Mul(largeFactorial, big.NewInt(int64(i)))
	}

	str := largeFactorial.String()
	for _, r := range str {
		output += int(r - '0')
	}

	return output
}

func init() {
	registry.Register("20", Solve020)
}
