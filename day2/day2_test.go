package day2

import (
	"testing"

	"github.com/cswilson90/aoc2020"
)

func TestDay2Part1(t *testing.T) {
	strings, err := aoc2020.ReadStringFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	answer, err := SledPasswordsMatchingPolicy(strings)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Errorf("Day 2 Part 1 answer: %v", answer)
}

func TestDay2Part2(t *testing.T) {
	strings, err := aoc2020.ReadStringFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	answer, err := TobogganPasswordsMatchingPolicy(strings)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Errorf("Day 2 Part 2 answer: %v", answer)
}
