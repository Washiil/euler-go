package problems

// Find the largest palindrome made from the product of two -digit numbers.

func intIsPalindrome(n int) bool {
	if n < 0 {
		return false
	}

	if n < 10 {
		return true
	}

	original := n
	reversed := 0

	for n > 0 {
		digit := n % 10

		// Build the reversed number by taking the existing reversed number,
		// multiplying it by 10 (to shift the digits to the left), and then
		// adding the current digit to the rightmost position.
		reversed = reversed*10 + digit

		n /= 10
	}

	return original == reversed
}

func Solve004() int {
	limit := 999

	palindromes := []int{} // Initialize an empty slice to store palindromes

	for i := limit; i > 100; i-- {
		for j := limit; j > i-1; j-- {
			x := (i * j)
			if intIsPalindrome(x) {
				palindromes = append(palindromes, x)
			}
		}
	}

	maximum := palindromes[0]
	for _, num := range palindromes {
		if num > maximum {
			maximum = num
		}
	}

	return maximum
}

func init() {
	Register("4", Solve004)
}
