package day5

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Takes a list of seat strings and returns the corresponding IDs
func SeatStringsToIDs(seatStrings []string) ([]int, error) {
	seatIDs := make([]int, len(seatStrings))
	for i, seatString := range seatStrings {
		if len(seatString) != 10 {
			return nil, fmt.Errorf("String %v not 10 chars long", seatString)
		}

		// Replace F/L with 0 and B/R with 1 and convert from binary string
		seatString = strings.Replace(seatString, "F", "0", -1)
		seatString = strings.Replace(seatString, "B", "1", -1)
		seatString = strings.Replace(seatString, "L", "0", -1)
		seatString = strings.Replace(seatString, "R", "1", -1)

		intValue, err := strconv.ParseInt(seatString, 2, 0)
		if err != nil {
			return nil, fmt.Errorf("Failed to convert seat %v: %v", seatString, err.Error())
		}

		seatIDs[i] = int(intValue)
	}

	return seatIDs, nil
}

// Finds the seat with the max ID from the given list of seat strings
func MaxSeatID(seatStrings []string) (int, error) {
	seatIDs, err := SeatStringsToIDs(seatStrings)
	if err != nil {
		return 0, err
	}

	maxID := 0
	for _, v := range seatIDs {
		if v > maxID {
			maxID = v
		}
	}

	return maxID, nil
}

// Finds the missing seat in the list of seat Strings ignoring non existent seats at front and back
func FindMySeat(seatStrings []string) (int, error) {
	seatIDs, err := SeatStringsToIDs(seatStrings)
	if err != nil {
		return 0, err
	}

	sort.Ints(seatIDs)
	lastSeat := seatIDs[0]
	for _, v := range seatIDs {
		if v == lastSeat+2 {
			return lastSeat + 1, nil
		}
		lastSeat = v
	}

	return 0, fmt.Errorf("Empty seat not found")
}
