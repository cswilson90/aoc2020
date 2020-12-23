package day22

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

var tests = []struct {
	input   string
	answer1 int
	answer2 int
}{
	{"example.txt", 306, 291},
	{"input.txt", 32413, 31596},
}

func TestExample(t *testing.T) {
	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			t.Parallel()

			input, err := aoc2020.ReadStringRecords(tc.input)
			if err != nil {
				t.Errorf(err.Error())
			}

			answer1 := PlayCombat(input[0], input[1])
			assert.Equal(t, tc.answer1, answer1)

			answer2 := PlayRecursiveCombat(input[0], input[1])
			assert.Equal(t, tc.answer2, answer2)
		})
	}
}
