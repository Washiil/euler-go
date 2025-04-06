// main.go
package main

import (
	"fmt"
	"os"

	"github.com/washiil/euler-go/problems"
)

func main() {
	problemMap := map[string](func() int){
		"1": problems.Solve001,
		"2": problems.Solve002,
		"3": problems.Solve003,
		"4": problems.Solve004,
	}

	if len(os.Args) < 2 {
		for i, fn := range problemMap {
			result := fn()
			fmt.Printf("Problem [%s] : %d\n", i, result)
		}
		return
	}

	arg := os.Args[1]
	if fn, ok := problemMap[arg]; ok {
		result := fn()
		fmt.Printf("Problem [%s] : %d", arg, result)
	} else {
		fmt.Printf("Problem %s not implemented.\n", arg)
	}
}
