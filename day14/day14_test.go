package day14

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

func TestExample(t *testing.T) {
	instructions, err := aoc2020.ReadStringFile("example1.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	bitMaskV1 := &BitMaskV1{}
	answer, err := RunDockingProgram(instructions, bitMaskV1)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, 165, int(answer), "Part 1 Incorrect")

	instructions, err = aoc2020.ReadStringFile("example2.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	bitMaskV2 := &BitMaskV2{}
	answer, err = RunDockingProgram(instructions, bitMaskV2)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, 208, int(answer), "Part 2 Incorrect")
}

func TestDay10(t *testing.T) {
	instructions, err := aoc2020.ReadStringFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	bitMaskV1 := &BitMaskV1{}
	answer, err := RunDockingProgram(instructions, bitMaskV1)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, 11179633149677, int(answer), "Part 1 Incorrect")

	bitMaskV2 := &BitMaskV2{}
	answer, err = RunDockingProgram(instructions, bitMaskV2)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, 4822600194774, int(answer), "Part 2 Incorrect")
}
