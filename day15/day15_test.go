package day15

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input   []int
	answer1 int
	answer2 int
}{
	{[]int{0, 3, 6}, 436, 175594},
	{[]int{1, 3, 2}, 1, 2578},
	{[]int{2, 1, 3}, 10, 3544142},
	{[]int{1, 2, 3}, 27, 261214},
	{[]int{19, 0, 5, 1, 10, 13}, 1015, 201},
}

func TestExample(t *testing.T) {
	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			t.Parallel()

			assert.Equal(t, tc.answer1, PlayMemoryGame(tc.input, 2020))
			assert.Equal(t, tc.answer2, PlayMemoryGame(tc.input, 30000000))
		})
	}
}
