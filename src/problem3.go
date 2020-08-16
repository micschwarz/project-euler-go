package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Result: %v", problem3(600851475143))
}

// problem2 gives a solution for project euler problem #3.
// See: https://projecteuler.net/problem=3.
// It returns the largest prime factor of a given number.
// The solution is not very nice and needs cleanup.
func problem3(max int) int {
	maxPrime := -1

	if max <= 1 {
		return maxPrime
	}

	rangeMax := int(math.Ceil(math.Sqrt(float64(max))))
	primes := createPrimesArray(rangeMax)

	for number := 1; number < rangeMax; number++ {

		if primes[number] {
			for numberKill := number * number; numberKill < rangeMax; numberKill += number {
				primes[numberKill] = false
			}
		}
	}

	for number, isPrime := range primes {
		if !isPrime {
			continue
		}

		for max%number == 0 {
			max /= number
			maxPrime = number
		}
	}

	return maxPrime
}

func createPrimesArray(length int) []bool {
	primes := make([]bool, length)

	for i := 2; i < len(primes); i++ {
		primes[i] = true
	}

	return primes
}
