package main

import "fmt"

func main() {
	fmt.Printf("Result: %v", problem2(4_000_000))
}

// problem2 gives a solution for project euler problem #2.
// See: https://projecteuler.net/problem=2.
// It returns the sum of the even-valued fibonacci-terms below a given number.
func problem2(max int) int {
	sum := 0
	// fibonacci-sequence starts with 1,1
	fib := [3]int{1,1}

	for fib[2] < max {
		fib[2] = fib[0] + fib[1]

		// Add term if even-valued
		if fib[2] % 2 == 0 {
			sum += fib[2]
		}

		// Shift fibonacci-sequence left
		fib[0] = fib[1]
		fib[1] = fib[2]
	}

	return sum
}
