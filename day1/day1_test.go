package day1

import (
	"io/ioutil"
	"strconv"
	"strings"
	"testing"
)

func readInput(t *testing.T) []int {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		t.Errorf(err.Error())
	}

	stringList := strings.Split(string(data), "\n")
	nums := make([]int, 0)
	for _, v := range stringList {
		if v == "" {
			continue
		}

		num, err := strconv.Atoi(v)
		if err != nil {
			t.Errorf(err.Error())
		}
		nums = append(nums, num)
	}

	return nums
}

func TestDay1Problem1(t *testing.T) {
	nums := readInput(t)

	result, err := productOfTwoThatSum(nums, 2020)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Errorf("Problem 1 result is %v", result)
}

func TestDay1Problem2(t *testing.T) {
	nums := readInput(t)

	result, err := productOfThreeThatSum(nums, 2020)
	if err != nil {
		t.Errorf(err.Error())
	}
	t.Errorf("Problem 2 result is %v", result)
}
