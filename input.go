package aoc2020

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadIntFile(filename string) ([]int, error) {
	stringList, err := ReadStringFile(filename)
	if err != nil {
		return nil, err
	}

	nums := make([]int, 0)
	for _, v := range stringList {
		if v == "" {
			continue
		}

		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}

	return nums, nil
}

func ReadStringRecords(filename string) ([]string, error) {
	records, err := readStringFile(filename, "\n\n")
	if err != nil {
		return nil, err
	}

	// Remove triling new line from last record if there is one
	lastRecordLength := len(records[len(records)-1])
	if records[len(records)-1][lastRecordLength-1] == '\n' {
		records[len(records)-1] = records[len(records)-1][:lastRecordLength-1]
	}

	return records, nil
}

func ReadStringFile(filename string) ([]string, error) {
	return readStringFile(filename, "\n")
}

func readStringFile(filename string, separator string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	allStrings := strings.Split(string(data), separator)

	// Remove empty last line if there is one
	if allStrings[len(allStrings)-1] == "" {
		allStrings = allStrings[:len(allStrings)-1]
	}

	return allStrings, nil
}
