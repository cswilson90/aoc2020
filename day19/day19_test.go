package day19

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
	{"example.txt", 3, 12},
	{"input.txt", 144, 260},
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

			rules := ParseRules(input[0])

			answer := MatchingMessages(input[1], rules)
			assert.Equal(t, tc.answer1, answer)

			answer = MatchingMessages2(input[1], rules)
			assert.Equal(t, tc.answer2, answer)
		})
	}
}
