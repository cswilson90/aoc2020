package day25

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var tests = []struct {
	input1  int
	input2  int
	answer1 int
}{
	{5764801, 17807724, 14897079},
	{5099500, 7648211, 11288669},
}

func TestExample(t *testing.T) {
	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			t.Parallel()

			encryptionKey := GetEncryptionKey(tc.input1, tc.input2)
			assert.Equal(t, tc.answer1, encryptionKey)
		})
	}
}
