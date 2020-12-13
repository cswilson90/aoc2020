package day13

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

func TestExample(t *testing.T) {
	test(t, "example.txt", 295, 1068781)
}

func TestDay10(t *testing.T) {
	test(t, "input.txt", 2092, 702970661767766)
}

func test(t *testing.T, inputFile string, part1 int, part2 int) {
	notesString, err := aoc2020.ReadStringFile(inputFile)
	if err != nil {
		t.Errorf(err.Error())
	}

	notes, err := ParseBusNotes(notesString)
	if err != nil {
		t.Errorf(err.Error())
	}

	departure := NextBusDeparture(notes)
	answer := (departure.Time - notes.StartTime) * departure.BusNo
	assert.Equal(t, part1, answer, "Part 1 Incorrect")

	answer, err = GetFirstTimeOfPattern(notesString[1])
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, part2, answer, "Part 2 Incorrect")
}
