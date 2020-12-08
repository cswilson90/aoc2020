package day3

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

func TestDay3(t *testing.T) {
	treeStrings, err := aoc2020.ReadStringFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	treeMap, err := GenerateTreeMap(treeStrings)
	if err != nil {
		t.Errorf(err.Error())
	}

	treeHits := CalculateTreeHits(treeMap, 3, 1)
	assert.Equal(t, 272, treeHits, "Part 1 incorrect")

	treeHits = CalculateTreeHits(treeMap, 1, 1)
	treeHits *= CalculateTreeHits(treeMap, 3, 1)
	treeHits *= CalculateTreeHits(treeMap, 5, 1)
	treeHits *= CalculateTreeHits(treeMap, 7, 1)
	treeHits *= CalculateTreeHits(treeMap, 1, 2)

	assert.Equal(t, 3898725600, treeHits, "Part 2 incorrect")
}
