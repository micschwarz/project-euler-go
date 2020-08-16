package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("Result: %v", problem4())
}

// problem4 gives a solution for project euler problem #4.
// See: https://projecteuler.net/problem=4.
// It returns the largest palindrome made from the product of two 3-digit numbers.
func problem4() int {
	biggest := 0

	for factor1 := 100; factor1 <= 999; factor1++ {
		for factor2 := 100; factor2 <= 999; factor2++ {
			if product := factor1 * factor2; isPalindromic(product) && product > biggest {
				biggest = product
			}
		}
	}

	return biggest
}

// isPalindromic tests if a number is a palindrome
func isPalindromic(num int) bool {
	number := strconv.Itoa(num)
	digits := []rune(number)

	// Look if x-th digit from left equals to x-th digit from right
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		if digits[i] != digits[j] {
			return false
		}
	}

	return true
}
