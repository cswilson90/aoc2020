package day5

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

func TestDay5(t *testing.T) {
	seatStrings, err := aoc2020.ReadStringFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	maxID, err := MaxSeatID(seatStrings)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, 955, maxID, "Part 1 incorrect")

	mySeat, err := FindMySeat(seatStrings)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, 569, mySeat, "Part 2 incorrect")
}
