package day15

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExample(t *testing.T) {
	tests := []struct {
		input   []int
		answer1 int
		answer2 int
	}{
		{[]int{0, 3, 6}, 436, 175594},
		// Commented out for speed as part 2 is quite slow
		//{ []int{1, 3, 2}, 1, 2578 },
		//{ []int{2, 1, 3}, 10, 3544142 },
		//{ []int{1, 2, 3}, 27, 261214 },
	}

	for _, test := range tests {
		answer := PlayMemoryGame(test.input, 2020)
		assert.Equal(t, test.answer1, answer)

		answer = PlayMemoryGame(test.input, 30000000)
		assert.Equal(t, test.answer2, answer)
	}
}

func TestDay(t *testing.T) {
	input := []int{19, 0, 5, 1, 10, 13}

	answer := PlayMemoryGame(input, 2020)
	assert.Equal(t, 1015, answer, "Part 1 Incorrect")

	answer = PlayMemoryGame(input, 30000000)
	assert.Equal(t, 201, answer, "Part 2 Incorrect")
}
