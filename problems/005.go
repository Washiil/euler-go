package problems

// What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20 ?

func Solve005() int {

	for i := 1; ; i++ {
		divisible := true
		for divisor := 1; divisor <= 20; divisor++ {
			if i%divisor != 0 {
				divisible = false
				break
			}
		}

		if divisible {
			return i
		}
	}
}
