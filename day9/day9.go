package day9

import (
	"fmt"

	"github.com/cswilson90/aoc2020"
)

// Given a list of xmas data with the given length of preamble finds the first
// value in the data which isn't valid
func FirstNonValidXmas(xmasData []int, preamble int) (int, error) {
	for i := preamble; i < len(xmasData); i++ {
		_, _, ok := aoc2020.TwoThatSum(xmasData[i-preamble:i], xmasData[i])
		if !ok {
			return xmasData[i], nil
		}
	}

	return 0, fmt.Errorf("All XMAS data is valid")
}

// Given a list of xmas data with the given length of preamble finds the encryption weakness of the data
func XmasEncryptionWeakness(xmasData []int, preamble int) (int, error) {
	target, err := FirstNonValidXmas(xmasData, preamble)
	if err != nil {
		return 0, err
	}

	// Find start and end of range which sums to the target
	start, end := 0, 0
	runningTotal := xmasData[0]

	for runningTotal != target {
		if runningTotal > target {
			runningTotal -= xmasData[start]
			start++
		} else if runningTotal < target {
			end++
			if end >= len(xmasData) {
				return 0, fmt.Errorf("Failed to calculate encryption weakness")
			}
			runningTotal += xmasData[end]
		}
	}

	// Find smallest and largest number in range
	smallest := xmasData[start]
	largest := xmasData[start]
	for _, v := range xmasData[start+1 : end+1] {
		if v < smallest {
			smallest = v
		}
		if v > largest {
			largest = v
		}
	}

	return smallest + largest, nil
}
