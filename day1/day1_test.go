package day1

import (
	"testing"

	"github.com/cswilson90/aoc2020"
)

func TestDay1Problem1(t *testing.T) {
	nums, err := aoc2020.ReadIntFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := productOfTwoThatSum(nums, 2020)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Errorf("Problem 1 result is %v", result)
}

func TestDay1Problem2(t *testing.T) {
	nums, err := aoc2020.ReadIntFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := productOfThreeThatSum(nums, 2020)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Errorf("Problem 2 result is %v", result)
}
