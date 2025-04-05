// main.go
package main

import (
	"fmt"
	"os"

	"github.com/washiil/euler-go/problems"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go [problem number]")
		return
	}

	problemMap := map[string](func() int){
		"1": problems.Solve001,
		"2": problems.Solve002,
	}

	arg := os.Args[1]
	if fn, ok := problemMap[arg]; ok {
		result := fn()
		fmt.Printf("Problem [%s] : %d", arg, result)
	} else {
		fmt.Printf("Problem %s not implemented.\n", arg)
	}
}
