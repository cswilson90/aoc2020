package day15

// Plays the elf memory game with the given input for the given number of turns
// Returns the number said on the last turn
func PlayMemoryGame(startNums []int, turns int) int {
	if turns <= len(startNums) {
		return startNums[turns-1]
	}

	seenLast := make(map[int]int)
	for i, v := range startNums {
		seenLast[v] = i + 1
	}

	number := startNums[len(startNums)-1]
	for turn := len(startNums) + 1; turn <= turns; turn++ {
		newNum := 0
		lastSeen, ok := seenLast[number]
		if ok {
			newNum = turn - 1 - lastSeen
		}
		seenLast[number] = turn - 1
		number = newNum
	}

	return number
}
