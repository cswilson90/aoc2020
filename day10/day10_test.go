package day10

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

func TestExample(t *testing.T) {
	test(t, "example.txt", 220, 19208)
}

func TestDay10(t *testing.T) {
	test(t, "input.txt", 2210, 7086739046912)
}

func test(t *testing.T, inputFile string, part1 int, part2 int) {
	numbers, err := aoc2020.ReadIntFile(inputFile)
	if err != nil {
		t.Errorf(err.Error())
	}

	answer, err := GetJoltDifference(numbers)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, part1, answer, "Part 1 Incorrect")

	answer = DistinctAdapterArrangements(numbers)
	assert.Equal(t, part2, answer, "Part 2 Incorrect")
}
