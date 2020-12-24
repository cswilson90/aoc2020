package day23

import "fmt"

type CupMap map[int]*Cup

type Cup struct {
	value int
	next  *Cup
}

func PlayCrabGame(cups []int) string {
	cupMap := buildCupList(cups, len(cups))
	playCrabGame(cupMap, cups[0], 100)

	nextCup := cupMap[1].next
	cupsString := ""
	for nextCup.value != 1 {
		cupsString += fmt.Sprintf("%v", nextCup.value)
		nextCup = nextCup.next
	}

	return cupsString
}

func PlayCrabGameExpanded(cups []int) int {
	cupMap := buildCupList(cups, 1000000)
	playCrabGame(cupMap, cups[0], 10000000)

	cup1 := cupMap[1]
	return cup1.next.value * cup1.next.next.value
}

func playCrabGame(cupMap CupMap, first int, moves int) {
	currentCup := cupMap[first]

	for move := 1; move <= moves; move++ {
		firstToMove := currentCup.next
		lastToMove := firstToMove.next.next

		destinationValue := getDestinationValue(currentCup.value, 1, len(cupMap))
		for inThreeToMove(destinationValue, firstToMove) {
			destinationValue = getDestinationValue(destinationValue, 1, len(cupMap))
		}
		destinationCup := cupMap[destinationValue]

		currentCup.next = lastToMove.next
		lastToMove.next = destinationCup.next
		destinationCup.next = firstToMove

		currentCup = currentCup.next
	}
}

// Build a circularly linked list of cups
func buildCupList(initialNums []int, size int) CupMap {
	cupMap := make(CupMap)
	var firstCup, lastCup *Cup

	for i := 0; i < size; i++ {
		value := i + 1
		if i < len(initialNums) {
			value = initialNums[i]
		}

		newCup := &Cup{value: value}
		cupMap[value] = newCup

		if firstCup == nil {
			firstCup = newCup
		} else {
			lastCup.next = newCup
		}

		lastCup = newCup
	}

	lastCup.next = firstCup
	return cupMap
}

func inThreeToMove(value int, first *Cup) bool {
	next := first
	for i := 1; i <= 3; i++ {
		if next.value == value {
			return true
		}
		next = next.next
	}
	return false
}

func getDestinationValue(current, min, max int) int {
	destinationValue := current - 1
	if destinationValue < min {
		destinationValue = max
	}
	return destinationValue
}
