package p018

import (
	"fmt"

	"github.com/washiil/euler-go/problems/registry"
)

var pyramid = [][]int{
	{75},
	{95, 64},
	{17, 47, 82},
	{18, 35, 87, 10},
	{20, 4, 82, 47, 65},
	{19, 1, 23, 75, 3, 34},
	{88, 2, 77, 73, 7, 63, 67},
	{99, 65, 4, 28, 6, 16, 70, 92},
	{41, 41, 26, 56, 83, 40, 80, 70, 33},
	{41, 48, 72, 33, 47, 32, 37, 16, 94, 29},
	{53, 71, 44, 65, 25, 43, 91, 52, 97, 51, 14},
	{70, 11, 33, 28, 77, 73, 17, 78, 39, 68, 17, 57},
	{91, 71, 52, 38, 17, 14, 91, 43, 58, 50, 27, 29, 48},
	{63, 66, 4, 68, 89, 53, 67, 30, 73, 16, 69, 87, 40, 31},
	{4, 62, 98, 27, 23, 9, 70, 98, 73, 93, 38, 53, 60, 4, 23},
}

var memo = make(map[string]int)

func recursive_solve(level int, index int, sum int) int {
	if level >= len(pyramid) || index > level || index < 0 {
		return sum
	}

	key := fmt.Sprintf("%d|%d", level, index)
	if val, found := memo[key]; found {
		return sum + val
	}

	// Dont change index becaues lowering a level causes it to switch naturally
	left_val := pyramid[level][index] + recursive_solve(level+1, index, sum)
	right_val := pyramid[level][index] + recursive_solve(level+1, index+1, sum)

	if left_val > right_val {
		memo[key] = left_val
		return left_val
	} else {
		memo[key] = right_val
		return right_val
	}
}

func Solve018() int {
	output := 0

	output = recursive_solve(0, 0, 0)

	return output
}

func init() {
	registry.Register("18", Solve018)
}
