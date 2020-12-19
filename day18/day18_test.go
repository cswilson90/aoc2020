package day18

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
	{"example.txt", 26457, 694173},
	{"input.txt", 13976444272545, 88500956630893},
}

func TestExample(t *testing.T) {
	assert.Equal(t, 231, EvaluateEquation("1 + 2 * 3 + 4 * 5 + 6", true))
	assert.Equal(t, 51, EvaluateEquation("1 + (2 * 3) + (4 * (5 + 6))", true))
	assert.Equal(t, 46, EvaluateEquation("2 * 3 + (4 * 5)", true))
	assert.Equal(t, 1445, EvaluateEquation("5 + (8 * 3 + 9 + 3 * 4 * 3)", true))
	assert.Equal(t, 669060, EvaluateEquation("5 * 9 * (7 * 3 * 3 + 9 * 3 + (8 + 6 * 4))", true))
	assert.Equal(t, 23340, EvaluateEquation("((2 + 4 * 9) * (6 + 9 * 8 + 6) + 6) + 2 + 4 * 2", true))

	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			t.Parallel()

			input, err := aoc2020.ReadStringFile(tc.input)
			if err != nil {
				t.Errorf(err.Error())
			}

			answer := EvaluateAllEquations(input, false)
			assert.Equal(t, tc.answer1, answer)

			answer = EvaluateAllEquations(input, true)
			assert.Equal(t, tc.answer2, answer)
		})
	}
}
