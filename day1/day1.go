package day1

import (
	"fmt"

	"github.com/cswilson90/aoc2020"
)

// Finds two numbers from the list that add up to the total and multiplies them
func productOfTwoThatSum(numbers []int, total int) (int, error) {
	num1, num2, ok := aoc2020.TwoThatSum(numbers, total)

	if ok {
		return num1 * num2, nil
	}

	return 0, fmt.Errorf("No two numbers sum to %v", total)
}

// Finds three numbers from the list that add up to the total and multiplies them
func productOfThreeThatSum(numbers []int, total int) (int, error) {
	for i, v := range numbers {
		product, err := productOfTwoThatSum(numbers[i+1:], total-v)
		if err == nil {
			return v * product, nil
		}
	}

	return 0, fmt.Errorf("No three numbers sum to %v", total)
}
