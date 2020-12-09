package aoc2020

func TwoThatSum(numbers []int, total int) (int, int, bool) {
	needed := make(map[int]int)

	for _, v := range numbers {
		other, ok := needed[v]
		if ok {
			return other, v, true
		}
		needed[total-v] = v
	}

	return 0, 0, false
}
