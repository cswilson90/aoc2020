package day9

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

func TestExample(t *testing.T) {
	test(t, "example.txt", 5, 127, 62)
}

func TestDay9(t *testing.T) {
	test(t, "input.txt", 25, 36845998, 4830226)
}

func test(t *testing.T, inputFile string, preamble int, part1 int, part2 int) {
	numbers, err := aoc2020.ReadIntFile(inputFile)
	if err != nil {
		t.Errorf(err.Error())
	}

	answer, err := FirstNonValidXmas(numbers, preamble)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, part1, answer, "Part 1 Incorrect")

	answer, err = XmasEncryptionWeakness(numbers, preamble)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(t, part2, answer, "Part 2 Incorrect")
}
