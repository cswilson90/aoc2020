package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

func TestExample(t *testing.T) {
	test(t, "example.txt", 25, 286)
}

func TestDay(t *testing.T) {
	test(t, "input.txt", 508, 30761)
}

func test(t *testing.T, inputFile string, part1 int, part2 int) {
	instructions, err := aoc2020.ReadStringFile(inputFile)
	if err != nil {
		t.Errorf(err.Error())
	}

	directionShip := &DirectionShip{
		Direction: 90,
	}
	answer, err := ManhattanDistance(instructions, directionShip)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, part1, answer, "Part 1 Incorrect")

	waypointShip := &WaypointShip{
		WaypointX: 10,
		WaypointY: 1,
	}
	answer, err = ManhattanDistance(instructions, waypointShip)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, part2, answer, "Part 2 Incorrect")
}
