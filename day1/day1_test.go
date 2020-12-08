package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

func TestDay1(t *testing.T) {
	nums, err := aoc2020.ReadIntFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	assert := assert.New(t)

	result, err := productOfTwoThatSum(nums, 2020)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(955584, result, "Part 1 incorrect result")

	result, err = productOfThreeThatSum(nums, 2020)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(287503934, result, "Part 2 incorrect result")
}
