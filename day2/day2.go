package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var policyMatch = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

func parsePolicy(policy string) (int, int, string, string, error) {
	matches := policyMatch.FindStringSubmatch(policy)
	if matches == nil {
		return 0, 0, "", "", fmt.Errorf("String '%v' did not match password policy regex", policy)
	}
	lowerBoundString, _ := strconv.Atoi(matches[1])
	lowerBound := int(lowerBoundString)
	upperBoundString, _ := strconv.Atoi(matches[2])
	upperBound := int(upperBoundString)

	matchingLetter := matches[3]
	password := matches[4]
	return lowerBound, upperBound, matchingLetter, password, nil
}

// Given a list of passwords and policies returns the number of passwords which pass their policy
func SledPasswordsMatchingPolicy(passwords []string) (int, error) {
	totalMatched := 0
	for _, policyString := range passwords {
		lowerBound, upperBound, matchingLetter, password, err := parsePolicy(policyString)
		if err != nil {
			return 0, err
		}

		occurences := strings.Count(password, matchingLetter)
		if occurences >= lowerBound && occurences <= upperBound {
			totalMatched += 1
		}
	}

	return totalMatched, nil
}

// Given a list of passwords and policies returns the number of passwords which pass their policy
func TobogganPasswordsMatchingPolicy(passwords []string) (int, error) {
	totalMatched := 0
	for _, policyString := range passwords {
		lowerBound, upperBound, matchingLetter, password, err := parsePolicy(policyString)
		if err != nil {
			return 0, err
		}

		positionMatches := 0
		for i, v := range password {
			if ((i+1) == lowerBound || (i+1) == upperBound) && string([]rune{v}) == matchingLetter {
				positionMatches++
			}
		}

		if positionMatches == 1 {
			totalMatched++
		}
	}

	return totalMatched, nil
}
