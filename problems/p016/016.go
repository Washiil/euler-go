package p016

import (
	"math/big"

	"github.com/washiil/euler-go/problems/registry"
)

func Solve016() int {
	exponent := big.NewInt(2)
	result := big.NewInt(1)

	for range 1_000 {
		result = result.Mul(result, exponent)
	}

	digits := result.String()
	sum := 0
	for _, r := range digits {
		sum += int(r - '0') // Convert rune to digit
	}

	return sum
}

func init() {
	registry.Register("16", Solve016)
}
