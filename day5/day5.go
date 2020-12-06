package day5

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Takes a list of seat strings and returns the corresponding IDs
func SeatStringssToIDs(seatStrings []string) ([]int, error) {
	seatIDs := make([]int, len(seatStrings))
	for i, seatString := range seatStrings {
		if len(seatString) != 10 {
			return nil, fmt.Errorf("String %v not 10 chars long", seatString)
		}

		seatRow, err := binaryStringToNumber(seatString[:7], "F", "B")
		if err != nil {
			return nil, err
		}
		seatColumn, err := binaryStringToNumber(seatString[7:], "L", "R")
		if err != nil {
			return nil, err
		}

		seatIDs[i] = (seatRow * 8) + seatColumn
	}

	return seatIDs, nil
}

// Finds the seat with the max ID from the given list of seat strings
func MaxSeatID(seatStrings []string) (int, error) {
	seatIDs, err := SeatStringssToIDs(seatStrings)
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
	seatIDs, err := SeatStringssToIDs(seatStrings)
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

// Converts a binary string encoded with the given letters to an int
// e.g. ("AABB", "A", "B") would give 3
func binaryStringToNumber(binString, zeroChar, oneChar string) (int, error) {
	binString = strings.Replace(binString, zeroChar, "0", -1)
	binString = strings.Replace(binString, oneChar, "1", -1)

	intValue, err := strconv.ParseInt(binString, 2, 0)
	if err != nil {
		return 0, fmt.Errorf("Failed to convert %v: %v", binString, err.Error())
	}

	return int(intValue), nil
}
