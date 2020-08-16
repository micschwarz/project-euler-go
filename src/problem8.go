package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Result: %v", problem8(13))
}

// problem8 gives a solution for project euler problem #8.
// See: https://projecteuler.net/problem=8.
// It returns the biggest product of n adjacent digits of a number defined in problem8_data.txt.
func problem8(adjacentDigits int) int {
	number := getData()
	maxProduct := 0

	for position := 0; position < len(number)-adjacentDigits; position++ {
		product := 1
		for offset := 0; offset < adjacentDigits; offset++ {
			product *= getDigitAt(position, offset, number)
		}

		if product > maxProduct {
			maxProduct = product
		}
	}

	return maxProduct
}

// getData reads data needed for the problem.
// It returns the content of the file problem8_data.txt.
func getData() string {
	fileContent, err := ioutil.ReadFile("./problem8_data.txt")
	if err != nil {
		panic(err)
	}
	return strings.Replace(string(fileContent), "\r\n", "", -1)
}

// getDigitAt gets the digit at a position (+offset) of a given string.
// It returns the digit at the given position.
func getDigitAt(pos int, offset int, number string) int {
	digit, err := strconv.Atoi(string(number[pos+offset]))

	if err != nil {
		panic(err)
	}

	return digit
}
