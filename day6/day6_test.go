package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

func TestDay6Part1(t *testing.T) {
	answerGroups, err := aoc2020.ReadStringRecords("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	sumOfCounts := SumAnswerCounts(answerGroups)
	assert.Equal(t, 6504, sumOfCounts, "Part 1 incorrect")

	sumOfCounts = SumAllAnswerCounts(answerGroups)
	assert.Equal(t, 3351, sumOfCounts, "Part 1 incorrect")
}
