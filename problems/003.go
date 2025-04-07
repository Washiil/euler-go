package problems

// Largest Prime Factor of 600851475143

func largestPrimeFactor(n int) int {
	i := 2
	for i*i < n {
		for n%i == 0 {
			n /= i
		}
		i += 1
	}
	return n
}

func Solve003() int {
	bigNum := 600851475143

	factor := largestPrimeFactor(bigNum)
	return factor
}

func init() {
	Register("3", Solve003)
}
