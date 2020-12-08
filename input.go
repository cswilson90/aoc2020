package aoc2020

import (
	"bufio"
	"io"
	"os"
	"strconv"
)

func ReadIntFile(filename string) ([]int, error) {
	stringList, err := ReadStringFile(filename)
	if err != nil {
		return nil, err
	}

	nums := make([]int, len(stringList))
	for i, v := range stringList {
		num, err := strconv.Atoi(v)
		if err != nil {
			return nil, err
		}
		nums[i] = num
	}

	return nums, nil
}

func ReadStringRecords(filename string) ([][]string, error) {
	lines, err := ReadStringFile(filename)
	if err != nil {
		return nil, err
	}

	records := make([][]string, 0)
	curRecord := make([]string, 0)

	for _, line := range lines {
		if line == "" {
			records = append(records, curRecord)
			curRecord = make([]string, 0)
		} else {
			curRecord = append(curRecord, line)
		}
	}
	if len(curRecord) > 0 {
		records = append(records, curRecord)
	}

	return records, nil
}

func ReadStringFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	reader := bufio.NewReader(file)

	allStrings := make([]string, 0)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return allStrings, nil
		} else if err != nil {
			return nil, err
		}

		allStrings = append(allStrings, line[:len(line)-1])
	}

	return nil, nil
}
