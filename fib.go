package main

import (
	. "fmt"
	"log"
	"math/big"
	"time"
)

// Dynamic Programming (Quy hoach dong)
func fibDP(n int) *big.Int {
	fn := make([]*big.Int, n+1)
	for i := 0; i <= n; i++ {
		var f = big.NewInt(0)
		if i <= 2 {
			f.SetUint64(1)
		} else {
			f = f.Add(fn[i-1], fn[i-2])
		}
		fn[i] = f
	}
	return fn[n]
}

// De qui
func fibRecursion(n int) *big.Int {
	var f = big.NewInt(0)
	if n <= 2 {
		return f.SetUint64(1)
	} else {
		return f.Add(fibRecursion(n-1), fibRecursion(n-2))
	}
}

func main() {
	Println("Calculate fibonacci using dynamic programming")
	start := time.Now()
	Println(fibDP(3000))
	elapsed := time.Since(start)
	log.Println("Fibonacci (DP) took ", elapsed)

	Println("Calculate fibonacci using recursion")
	start2 := time.Now()
	Println(fibRecursion(3000))
	elapsed2 := time.Since(start2)
	log.Println("Fibonacci (Recursion) took ", elapsed2)
}
