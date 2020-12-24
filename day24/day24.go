package day24

import "log"

type BlackHexes map[Hex]bool

// Stores a hex coordinate where x + y + z = 0
type Hex struct {
	x, y, z int
}

// Returns a map of all the black hexes after initialising from the instructions.
func GetBlackHexes(instructions []string) BlackHexes {
	blackHexes := make(BlackHexes)

	for _, instruction := range instructions {
		hex := Hex{0, 0, 0}
		instructionRunes := []rune(instruction)
		nextRune := 0
		for nextRune < len(instructionRunes) {
			firstRune := instructionRunes[nextRune]
			switch firstRune {
			case 'e', 'w':
				hex = moveHex(hex, string([]rune{firstRune}))
			case 'n', 's':
				nextRune++
				hex = moveHex(hex, string([]rune{firstRune, instructionRunes[nextRune]}))
			default:
				log.Fatalf("Unrecognised instruction: %v", firstRune)
			}
			nextRune++
		}

		if blackHexes[hex] {
			delete(blackHexes, hex)
		} else {
			blackHexes[hex] = true
		}
	}

	return blackHexes
}

// Returns the number of black hexes after iterating the floor for a number of days.
func BlackHexesAfterDays(blackHexes BlackHexes, numDays int) int {
	for i := 1; i <= numDays; i++ {
		newBlackHexes := make(BlackHexes)
		activeNeighbours := make(map[Hex]int)

		for hex := range blackHexes {
			for _, adjacent := range adjacentHexes(hex) {
				activeNeighbours[adjacent]++
			}
		}

		for hex, blackNeighbours := range activeNeighbours {
			if (blackHexes[hex] && blackNeighbours <= 2) || (!blackHexes[hex] && blackNeighbours == 2) {
				newBlackHexes[hex] = true
			}
		}

		blackHexes = newBlackHexes
	}

	return len(blackHexes)
}

func moveHex(hex Hex, direction string) Hex {
	x, y, z := hex.x, hex.y, hex.z
	switch direction {
	case "e":
		x++
		y--
	case "w":
		x--
		y++
	case "ne":
		x++
		z--
	case "sw":
		x--
		z++
	case "nw":
		y++
		z--
	case "se":
		y--
		z++
	default:
		log.Fatalf("Unrecognised direction: %v", direction)
	}
	return Hex{x, y, z}
}

func adjacentHexes(hex Hex) []Hex {
	adjacent := make([]Hex, 0)
	for _, direction := range []string{"e", "w", "ne", "sw", "nw", "se"} {
		adjacent = append(adjacent, moveHex(hex, direction))
	}
	return adjacent
}
