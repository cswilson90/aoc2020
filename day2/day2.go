package day2

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type PasswordPolicy struct {
	LowerBound     int
	UpperBound     int
	MatchingLetter string
	Password       string
}

var policyMatch = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

func parsePasswordPolicy(policy string) (*PasswordPolicy, error) {
	matches := policyMatch.FindStringSubmatch(policy)
	if matches == nil {
		return nil, fmt.Errorf("String '%v' did not match password policy regex", policy)
	}

	lowerBound, _ := strconv.Atoi(matches[1])
	upperBound, _ := strconv.Atoi(matches[2])

	return &PasswordPolicy{
		LowerBound:     lowerBound,
		UpperBound:     upperBound,
		MatchingLetter: matches[3],
		Password:       matches[4],
	}, nil
}

// Given a list of passwords and policies returns the number of passwords which pass their policy
func SledPasswordsMatchingPolicy(passwords []string) (int, error) {
	totalMatched := 0
	for _, policyString := range passwords {
		pp, err := parsePasswordPolicy(policyString)
		if err != nil {
			return 0, err
		}

		occurences := strings.Count(pp.Password, pp.MatchingLetter)
		if occurences >= pp.LowerBound && occurences <= pp.UpperBound {
			totalMatched += 1
		}
	}

	return totalMatched, nil
}

// Given a list of passwords and policies returns the number of passwords which pass their policy
func TobogganPasswordsMatchingPolicy(passwords []string) (int, error) {
	totalMatched := 0
	for _, policyString := range passwords {
		pp, err := parsePasswordPolicy(policyString)
		if err != nil {
			return 0, err
		}

		positionMatches := 0
		for i, v := range pp.Password {
			if ((i+1) == pp.LowerBound || (i+1) == pp.UpperBound) && string([]rune{v}) == pp.MatchingLetter {
				positionMatches++
			}
		}

		if positionMatches == 1 {
			totalMatched++
		}
	}

	return totalMatched, nil
}
