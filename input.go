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

func ReadStringFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	return strings.Split(string(data), "\n"), nil
}
