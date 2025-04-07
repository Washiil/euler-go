package problems

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
