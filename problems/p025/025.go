package p025

import (
	"math/big"

	"github.com/washiil/euler-go/problems/registry"
)

func fibonacciGenerator() <-chan *big.Int {
	ch := make(chan *big.Int)
	go func() {
		a := big.NewInt(0)
		b := big.NewInt(1)
		for {
			ch <- a
			a.Add(a, b)
			a, b = b, a
		}
	}()
	return ch
}

func Solve025() int {
	fib := fibonacciGenerator()
	num := <-fib
	count := 1

	for len(num.String()) < 1000 {
		num = <-fib
		count += 1
	}

	return count
}

func init() {
	registry.Register("025", Solve025)
}
