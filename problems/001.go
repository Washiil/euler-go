package problems

/*
Find the sum of all multiples of 3 or 5 below 1000
*/

func Solve001() int {
	sum := 0
	for i := range 1000 {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	return sum
}
