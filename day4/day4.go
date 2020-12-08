package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var fieldSplit = regexp.MustCompile(`^(.*):(.*)$`)

// Parsed a list of passport strings into a list of passports
func ParsePassports(passportList [][]string) ([]map[string]string, error) {
	passports := make([]map[string]string, 0)

	for _, passportStrings := range passportList {
		passportString := strings.Join(passportStrings, " ")
		fields := strings.Split(passportString, " ")

		passport := make(map[string]string)
		for j, field := range fields {
			if field == "" {
				continue
			}

			matches := fieldSplit.FindStringSubmatch(field)
			if matches == nil {
				return nil, fmt.Errorf("Error parsing passport '%v', field '%v' couldn't split", passportString, j)
			}
			passport[matches[1]] = matches[2]
		}
		passports = append(passports, passport)
	}

	return passports, nil
}

// Counts the number of valid passports in the list of passports given
func NumberOfValidPassports(passports []map[string]string) int {
	numValid := 0

	for _, passport := range passports {
		if hasAllFields(passport) {
			numValid++
		}
	}

	return numValid
}

// Counts the number of valid passports in the list of passports given
func NumberOfValidPassportsStrict(passports []map[string]string) int {
	numValid := 0

	for _, passport := range passports {
		if allFieldsValid(passport) {
			numValid++
		}
	}

	return numValid
}

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func hasAllFields(passport map[string]string) bool {
	for _, field := range requiredFields {
		_, ok := passport[field]
		if !ok {
			return false
		}
	}
	return true
}

func allFieldsValid(passport map[string]string) bool {
	for _, field := range requiredFields {
		value, ok := passport[field]
		if !ok {
			return false
		}
		if !fieldValid(field, value) {
			return false
		}
	}
	return true
}

var validHairColour = regexp.MustCompile(`^#[0-9a-f]{6}$`)
var validPID = regexp.MustCompile(`^\d{9}$`)

func fieldValid(name string, value string) bool {
	switch name {
	case "byr":
		return validIntRange(value, 1920, 2002)
	case "iyr":
		return validIntRange(value, 2010, 2020)
	case "eyr":
		return validIntRange(value, 2020, 2030)
	case "hgt":
		return validHeight(value)
	case "hcl":
		return validHairColour.MatchString(value)
	case "ecl":
		return validEyeColour(value)
	case "pid":
		return validPID.MatchString(value)
	default:
		return true
	}
}

func validIntRange(value string, lowerBound int, upperBound int) bool {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		return false
	}

	return intValue >= lowerBound && intValue <= upperBound
}

var heightSplit = regexp.MustCompile(`^(\d{2,3})(cm|in)$`)

func validHeight(value string) bool {
	matches := heightSplit.FindStringSubmatch(value)
	if matches == nil {
		return false
	}

	return (matches[2] == "cm" && validIntRange(matches[1], 150, 193)) ||
		(matches[2] == "in" && validIntRange(matches[1], 59, 76))
}

var validEyeColours = map[string]bool{
	"amb": true,
	"blu": true,
	"brn": true,
	"gry": true,
	"grn": true,
	"hzl": true,
	"oth": true,
}

func validEyeColour(value string) bool {
	_, ok := validEyeColours[value]
	return ok
}
