package p006

import "github.com/washiil/euler-go/problems/registry"

func Solve006() int {
	n := 100

	sumOfSquares := 0
	sum := 0

	for i := 1; i <= n; i++ {
		sumOfSquares += i * i
		sum += i
	}

	squareOfSum := sum * sum
	difference := squareOfSum - sumOfSquares

	return difference
}

func init() {
	registry.Register("6", Solve006)
}
