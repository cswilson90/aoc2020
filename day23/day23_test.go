package day23

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input   []int
	answer1 string
	answer2 int
}{
	{[]int{3, 8, 9, 1, 2, 5, 4, 6, 7}, "67384529", 149245887792},
	{[]int{5, 6, 2, 8, 9, 3, 1, 4, 7}, "38925764", 131152940564},
}

func TestExample(t *testing.T) {
	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			t.Parallel()

			answer1 := PlayCrabGame(tc.input)
			assert.Equal(t, tc.answer1, answer1)

			answer2 := PlayCrabGameExpanded(tc.input)
			assert.Equal(t, tc.answer2, answer2)
		})
	}
}
