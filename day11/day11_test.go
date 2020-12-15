package day11

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

func TestExample(t *testing.T) {
	test(t, "example.txt", 37, 26)
}

func TestDay(t *testing.T) {
	test(t, "input.txt", 2277, 2066)
}

func test(t *testing.T, inputFile string, part1 int, part2 int) {
	seatStrings, err := aoc2020.ReadStringFile(inputFile)
	if err != nil {
		t.Errorf(err.Error())
	}

	seatMap, err := ParseSeatStrings(seatStrings)
	if err != nil {
		t.Errorf(err.Error())
	}

	answer := EquilibriumOccupation(seatMap, OccupiedNeighbours, 4)
	assert.Equal(t, part1, answer, "Part 1 Incorrect")

	answer = EquilibriumOccupation(seatMap, OccupiedVisibleNeighbours, 5)
	assert.Equal(t, part2, answer, "Part 2 Incorrect")
}
