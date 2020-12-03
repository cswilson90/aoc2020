package day3

import (
	"testing"

	"github.com/cswilson90/aoc2020"
)

func TestDay3Part1(t *testing.T) {
	treeStrings, err := aoc2020.ReadStringFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	treeMap, err := GenerateTreeMap(treeStrings)
	if err != nil {
		t.Errorf(err.Error())
	}

	treeHits := CalculateTreeHits(treeMap, 3, 1)
	t.Errorf("Day 3 Part 1 Answer: %v", treeHits)
}

func TestDay3Part2(t *testing.T) {
	treeStrings, err := aoc2020.ReadStringFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	treeMap, err := GenerateTreeMap(treeStrings)
	if err != nil {
		t.Errorf(err.Error())
	}

	treeHits := CalculateTreeHits(treeMap, 1, 1)
	treeHits *= CalculateTreeHits(treeMap, 3, 1)
	treeHits *= CalculateTreeHits(treeMap, 5, 1)
	treeHits *= CalculateTreeHits(treeMap, 7, 1)
	treeHits *= CalculateTreeHits(treeMap, 1, 2)

	t.Errorf("Day 3 Part 1 Answer: %v", treeHits)
}
