package day4

import (
	"testing"

	"github.com/cswilson90/aoc2020"
)

func TestDay4Part1(t *testing.T) {
	passportStrings, err := aoc2020.ReadStringRecords("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	passports, err := ParsePassports(passportStrings)
	if err != nil {
		t.Errorf(err.Error())
	}

	numValid := NumberOfValidPassports(passports)
	t.Errorf("Day 4 Part 1 Answer: %v", numValid)
}

func TestDay4Part2(t *testing.T) {
	passportStrings, err := aoc2020.ReadStringRecords("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	passports, err := ParsePassports(passportStrings)
	if err != nil {
		t.Errorf(err.Error())
	}

	numValid := NumberOfValidPassportsStrict(passports)
	t.Errorf("Day 4 Part 2 Answer: %v", numValid)
}
