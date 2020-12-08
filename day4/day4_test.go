package day4

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

func TestDay4(t *testing.T) {
	passportStrings, err := aoc2020.ReadStringRecords("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	passports, err := ParsePassports(passportStrings)
	if err != nil {
		t.Errorf(err.Error())
	}

	numValid := NumberOfValidPassports(passports)
	assert.Equal(t, 192, numValid, "Part 1 correct")

	numValid = NumberOfValidPassportsStrict(passports)
	assert.Equal(t, 101, numValid, "Part 2 correct")
}
