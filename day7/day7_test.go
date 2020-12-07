package day7

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

func TestBagParsing(t *testing.T) {
	testStrings := []string{
		"mirrored beige bags contain 4 pale gold bags, 1 pale aqua bag.",
		"pale maroon bags contain 2 dotted orange bags.",
		"dim tan bags contain no other bags.",
	}

	bagMap, err := ParseBagStrings(testStrings)
	if err != nil {
		t.Errorf(err.Error())
	}

	beigeBag, ok := bagMap["mirrored beige"]
	if !ok {
		t.Errorf("beige Bag nt found in parsed output")
	}

	assert := assert.New(t)
	assert.Equal(beigeBag.Colour, "mirrored beige", "Beige bag has incorrect colour")
	assert.Equal(len(beigeBag.Contains), 2, "Beige bag has incorrect length of links")
	assert.Equal(beigeBag.Contains[0].Colour, "pale gold", "Beige bag has correct first link colour")
	assert.Equal(beigeBag.Contains[0].Number, 4, "Beige bag has correct first link number")

	tanBag, ok := bagMap["dim tan"]
	if !ok {
		t.Errorf("Tan bag missing from parsed output")
	}
	assert.Equal(len(tanBag.Contains), 0, "Tan bag has incorrect length of links")
}

func TestDay7(t *testing.T) {
	bagStrings, err := aoc2020.ReadStringFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	bagMap, err := ParseBagStrings(bagStrings)
    if err != nil {
		t.Errorf(err.Error())
    }

	colourCount := CountColoursContainingGold(bagMap)
	t.Errorf("Day 7 Part 1 answer: %v", colourCount)

	insideGoldCount, err := InsideShinyGoldBagCount(bagMap)
	if err != nil {
		t.Errorf(err.Error())
	}

	t.Errorf("Day 7 Part 2 answer: %v", insideGoldCount)
}
