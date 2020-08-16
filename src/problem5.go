package main

import "fmt"

func main() {
	fmt.Printf("Result: %v", problem5(20))
}

// problem5 gives a solution for project euler problem #5.
// See: https://projecteuler.net/problem=5.
// It returns the smallest positive number that is evenly divisible by all of the numbers below a given maxDivisor.
func problem5(maxDivisor int) int {
	maxOddDivisor := maxDivisor
	if maxDivisor%2 != 0 {
		maxOddDivisor--
	}

	for evenFactor := 2; true; evenFactor += 2 {
		// Solution has to be even
		solution := evenFactor * maxOddDivisor

		// Solution has to be dividable by 5 and all other divider
		if (maxDivisor >= 5 && solution%5 == 0) && dividableByAllBelow(maxDivisor, solution) {
			return solution
		}
	}

	return -1
}

// dividableByAllBelow checks if a number is dividable by all numbers below a given maxDivisor
func dividableByAllBelow(maxDivisor int, number int) bool {
	for divisor := maxDivisor; divisor > 1; divisor-- {
		if number%divisor != 0 {
			return false
		}
	}

	return true
}
