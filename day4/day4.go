package day4

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Passport struct {
	Fields map[string]string
}

var (
	fieldSplit = regexp.MustCompile(`^(.*):(.*)$`)

	requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

	heightSplit = regexp.MustCompile(`^(\d{2,3})(cm|in)$`)

	validHairColour = regexp.MustCompile(`^#[0-9a-f]{6}$`)
	validPID        = regexp.MustCompile(`^\d{9}$`)
	validEyeColours = map[string]struct{}{
		"amb": struct{}{},
		"blu": struct{}{},
		"brn": struct{}{},
		"gry": struct{}{},
		"grn": struct{}{},
		"hzl": struct{}{},
		"oth": struct{}{},
	}
)

// Parsed a list of passport strings into a list of passports
func ParsePassports(passportList [][]string) ([]*Passport, error) {
	passports := make([]*Passport, 0)

	for _, passportStrings := range passportList {
		passportString := strings.Join(passportStrings, " ")
		fields := strings.Split(passportString, " ")

		passport := &Passport{
			Fields: make(map[string]string),
		}

		for j, field := range fields {
			if field == "" {
				continue
			}

			matches := fieldSplit.FindStringSubmatch(field)
			if matches == nil {
				return nil, fmt.Errorf("Error parsing passport '%v', field '%v' couldn't split", passportString, j)
			}
			passport.Fields[matches[1]] = matches[2]
		}
		passports = append(passports, passport)
	}

	return passports, nil
}

// Counts the number of valid passports in the list of passports given
func NumberOfValidPassports(passports []*Passport) int {
	numValid := 0

	for _, passport := range passports {
		if passport.hasAllFields() {
			numValid++
		}
	}

	return numValid
}

// Counts the number of valid passports in the list of passports given
func NumberOfValidPassportsStrict(passports []*Passport) int {
	numValid := 0

	for _, passport := range passports {
		if passport.allFieldsValid() {
			numValid++
		}
	}

	return numValid
}

func (p *Passport) hasAllFields() bool {
	for _, field := range requiredFields {
		_, ok := p.Fields[field]
		if !ok {
			return false
		}
	}
	return true
}

func (p *Passport) allFieldsValid() bool {
	for _, field := range requiredFields {
		value, ok := p.Fields[field]
		if !ok {
			return false
		}
		if !fieldValid(field, value) {
			return false
		}
	}
	return true
}

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

func validHeight(value string) bool {
	matches := heightSplit.FindStringSubmatch(value)
	if matches == nil {
		return false
	}

	return (matches[2] == "cm" && validIntRange(matches[1], 150, 193)) ||
		(matches[2] == "in" && validIntRange(matches[1], 59, 76))
}

func validEyeColour(value string) bool {
	_, ok := validEyeColours[value]
	return ok
}
