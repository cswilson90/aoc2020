package day21

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

var tests = []struct {
	input   string
	answer1 int
	answer2 string
}{
	{"example.txt", 5, "mxmxvkd,sqjhc,fvjkl"},
	{"input.txt", 2307, "cljf,frtfg,vvfjj,qmrps,hvnkk,qnvx,cpxmpc,qsjszn"},
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

			answer1 := CountNonAllergens(input)
			assert.Equal(t, tc.answer1, answer1)

			answer2 := DangerousIngredients(input)
			assert.Equal(t, tc.answer2, answer2)
		})
	}
}
