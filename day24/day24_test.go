package day24

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
	{"example.txt", 10, 2208},
	{"input.txt", 360, 3924},
}

func TestExample(t *testing.T) {
	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			t.Parallel()

			input, err := aoc2020.ReadStringFile(tc.input)
			if err != nil {
				t.Errorf(err.Error())
			}

			blackHexes := GetBlackHexes(input)
			assert.Equal(t, tc.answer1, len(blackHexes))

			numBlack := BlackHexesAfterDays(blackHexes, 100)
			assert.Equal(t, tc.answer2, numBlack)
		})
	}
}
