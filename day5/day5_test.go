package day5

import (
	"testing"

	"github.com/cswilson90/aoc2020"
)

func TestDay5Part1(t *testing.T) {
	seatStrings, err := aoc2020.ReadStringFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	maxID, err := MaxSeatID(seatStrings)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Errorf("Day 5 Part 1 Answer: %v", maxID)
}

func TestDay5Part2(t *testing.T) {
	seatStrings, err := aoc2020.ReadStringFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	mySeat, err := FindMySeat(seatStrings)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Errorf("Day 5 Part 2 Answer: %v", mySeat)
}
