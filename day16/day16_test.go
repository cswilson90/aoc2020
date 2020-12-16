package day16

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/cswilson90/aoc2020"
)

var DepartureMatch = regexp.MustCompile(`^departure`)

var tests = []struct {
	input   string
	answer1 int
	answer2 int
}{
	{"example.txt", 71, 1},
	{"input.txt", 21956, 3709435214239},
}

func TestExample(t *testing.T) {
	for i, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("Test %v", i), func(t *testing.T) {
			t.Parallel()

			records, err := aoc2020.ReadStringRecords(tc.input)
			if err != nil {
				t.Errorf(err.Error())
			}

			rules := ParseRules(records[0])
			if records[1][0] != "your ticket:" {
				t.Errorf("Couldn't find your ticket")
			}
			if records[2][0] != "nearby tickets:" {
				t.Errorf("Couldn't find nearby tickets")
			}

			tickets := records[2][1:]

			assert.Equal(t, tc.answer1, TicketScanningError(tickets, rules))

			fieldMapping := FindFieldsInTickets(tickets, rules)

			myTicket := ParseTicket(records[1][1])
			product := 1
			for field, index := range fieldMapping {
				if DepartureMatch.MatchString(field) {
					product *= myTicket[index]
				}
			}

			assert.Equal(t, tc.answer2, product)
		})
	}
}
