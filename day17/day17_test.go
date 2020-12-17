package day17

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
	{"example.txt", 112, 848},
	{"input.txt", 242, 2292},
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

			activeCubes := RunBootProcess3D(input)
			assert.Equal(t, tc.answer1, activeCubes)

			activeCubes = RunBootProcess4D(input)
			assert.Equal(t, tc.answer2, activeCubes)
		})
	}
}
