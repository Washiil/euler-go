package p009

import (
	"github.com/washiil/euler-go/problems/registry"
)

// a < b < c and a + b + c = 1000
// a^2 + b^2 = c^2
// Find the product a * b * c

func Solve009() int {
	for a := 1; a < 333; a++ {
		for b := a + 1; b < 666; b++ {
			c := 1000 - a - b
			if c > b && ((a*a)+(b*b)) == (c*c) {
				return a * b * c
			}
		}
	}

	return -1
}

func init() {
	registry.Register("9", Solve009)
}
