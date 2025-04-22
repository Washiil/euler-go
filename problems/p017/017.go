package p017

import (
	"github.com/divan/num2words"
	"github.com/washiil/euler-go/problems/registry"
)

func Solve017() int {
	total_sum := 0

	for i := 1; i <= 1_000; i++ {
		str := num2words.ConvertAnd(i)
		for _, r := range str {
			if r != ' ' && r != '-' {
				total_sum += 1
			}
		}
	}

	return total_sum
}

func init() {
	registry.Register("17", Solve017)
}
