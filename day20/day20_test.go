package day20

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
	{"example.txt", 20899048083289, 273},
	{"input.txt", 66020135789767, 1537},
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

			tiles := ParseTiles(input)

			answer1, answer2 := FindAnswers(tiles)
			assert.Equal(t, tc.answer1, answer1)

			assert.Equal(t, tc.answer2, answer2)
		})
	}
}
