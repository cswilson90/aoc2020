package day10

import (
	"fmt"
	"sort"
)

// Returns the multiple of the number of 1 diff and 2 diff adapters
func GetJoltDifference(adapters []int) (int, error) {
	sort.Ints(adapters)

	differences := map[int]int{1: 0, 2: 0, 3: 0}

	previous := 0
	for _, v := range adapters {
		diff := v - previous
		if diff < 0 || diff > 3 {
			return 0, fmt.Errorf("Diff between %v and %v unexpected: %v", previous, v, diff)
		}
		differences[diff]++
		previous = v
	}

	// Add on final 3 diff for adapter to device
	differences[3]++

	return differences[1] * differences[3], nil
}

// Returns the number of distinct adapter combinations that work
func DistinctAdapterArrangements(adapters []int) int {
	// Sort adapters and add start (0) and end device to them
	adapters = append(adapters, 0)
	sort.Ints(adapters)
	adapters = append(adapters, adapters[len(adapters)-1]+3)

	// Stores number of paths from given index to end device
	indexPaths := make([]int, len(adapters))

	indexPaths[len(adapters)-1] = 1
	for i := len(adapters) - 2; i >= 0; i-- {
		for j := i + 1; j < len(adapters); j++ {
			if adapters[j]-adapters[i] > 3 {
				break
			}
			indexPaths[i] += indexPaths[j]
		}
	}

	return indexPaths[0]
}
