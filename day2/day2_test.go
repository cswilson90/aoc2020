package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

func TestDay2(t *testing.T) {
	strings, err := aoc2020.ReadStringFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	assert := assert.New(t)
	answer, err := SledPasswordsMatchingPolicy(strings)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(645, answer, "Part 1 incorrect")

	answer, err = TobogganPasswordsMatchingPolicy(strings)
	if err != nil {
		t.Errorf(err.Error())
	}
	assert.Equal(737, answer, "Part 2 incorrect")
}
