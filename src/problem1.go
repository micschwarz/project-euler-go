package main

import "fmt"

func main() {
	fmt.Printf("Result: %v", problem1(1000))
}

// problem1 gives a solution for project euler problem #1.
// See: https://projecteuler.net/problem=1.
// It returns the sum of all the multiples of 3 or 5 below a given number.
func problem1(max int) int {
	sum := 0

	for i := 1; i*3 < max; i++ {
		// All multiples of 3 are added
		sum += i*3
		// Add multiple of 5 if its less than the given number and not also a multiple of 3
		if mul5 := i*5; mul5 < max && mul5 % 3 != 0 {
			sum += mul5
		}
	}

	return sum
}