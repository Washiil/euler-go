package p015

import (
	"fmt"

	"github.com/washiil/euler-go/problems/registry"
)

var memo = make(map[string]int)

func recursive_solve(x int, y int, width int, height int) int {
	key := fmt.Sprintf("%d,%d", x, y)
	if val, found := memo[key]; found {
		return val
	}
	// 0, 0 represents the top left of the matrix
	// From wherever we are we can increment our x or our y but never subtract
	// Thus we only need to perform a simply bounds check
	if x == width && y == height {
		return 1
	} else if x > width || y > height {
		return 0
	}

	paths := 0
	if x+1 <= width {
		paths += recursive_solve(x+1, y, width, height)
	}

	if y+1 <= height {
		paths += recursive_solve(x, y+1, width, height)
	}
	memo[key] = paths
	return paths
}

func Solve015() int {
	paths := 0
	paths = recursive_solve(0, 0, 20, 20)
	return paths
}

func init() {
	registry.Register("15", Solve015)
}
