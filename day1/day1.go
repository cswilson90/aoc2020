package day1

import "fmt"

// Finds two numbers from the list that add up to the total and multiplies them
func productOfTwoThatSum(numbers []int, total int) (int, error) {
	needed := make(map[int]int)

	for _, v := range numbers {
		other, ok := needed[v]
		if ok {
			return other * v, nil
		}
		needed[total - v] = v
	}

	return 0, fmt.Errorf("No two numbers sum to %v", total)
}

// Finds three numbers from the list that add up to the total and multiplies them
func productOfThreeThatSum(numbers []int, total int) (int, error) {
	for i, v := range numbers {
		product, err := productOfTwoThatSum(numbers[i+1:], total - v)
		if err == nil {
			return v * product, nil
		}
	}

	return 0, fmt.Errorf("No three numbers sum to %v", total)
}
