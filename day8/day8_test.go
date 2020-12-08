package day8

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

func TestExample(t *testing.T) {
	instructions := []string{
		"nop +0",
		"acc +1",
		"jmp +4",
		"acc +3",
		"jmp -3",
		"acc -99",
		"acc +1",
		"jmp -4",
		"acc +6",
	}
	runInstructions(t, instructions, 5, 8)
}

func TestDay8(t *testing.T) {
	instructions, err := aoc2020.ReadStringFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}
	runInstructions(t, instructions, 1475, 1270)
}

func runInstructions(t *testing.T, instructions []string, part1, part2 int) {
	machine, err := ParseMachineInstructions(instructions)
	if err != nil {
		t.Errorf(err.Error())
	}

	result, err := machine.FindInfiniteLoopValue()
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, part1, result, "Incorrect infinite loop value")

	result2, err := machine.FixInfiniteLoop()
	if err != nil {
		t.Errorf(err.Error())
	}

	assert.Equal(t, part2, result2, "Incorrect fixed infinite loop value")
}
