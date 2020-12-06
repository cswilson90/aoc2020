package day6

import (
	"testing"

	"github.com/cswilson90/aoc2020"
)

func TestDay6Part1(t *testing.T) {
	answerGroups, err := aoc2020.ReadStringRecords("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	sumOfCounts := SumAnswerCounts(answerGroups)
	t.Errorf("Day 6 Part 1 Answer: %v", sumOfCounts)
}

func TestDay6Part2(t *testing.T) {
	answerGroups, err := aoc2020.ReadStringRecords("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	sumOfCounts := SumAllAnswerCounts(answerGroups)
	t.Errorf("Day 6 Part 2 Answer: %v", sumOfCounts)
}
