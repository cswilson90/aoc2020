package day11

import "fmt"

type (
	CountNeighbours func(*SeatMap, int, int) int
)

type SeatMap struct {
	seats    [][]rune
	occupied int
}

// Parses the input list of seat string into a seat map
func ParseSeatStrings(seatStrings []string) (*SeatMap, error) {
	seatMap := &SeatMap{
		seats: make([][]rune, len(seatStrings)),
	}
	rowLength := len(seatStrings[0])

	for i, seatString := range seatStrings {
		if len(seatString) != rowLength {
			return nil, fmt.Errorf("Row %v different length to first row", i)
		}

		seatMap.seats[i] = make([]rune, rowLength)
		for j, char := range seatString {
			seatMap.seats[i][j] = char
		}
	}

	return seatMap, nil
}

// Calulates the number of seats occupied once the given seat map reaches a state of equilibrium.
// Iterates the map with the given CountNeighbours function and neighbout limit
func EquilibriumOccupation(seatMap *SeatMap, count CountNeighbours, neighbourLimit int) int {
	nextStep := seatMap
	reachedEquilibrium := false

	for !reachedEquilibrium {
		nextStep, reachedEquilibrium = stepSeatMap(nextStep, count, neighbourLimit)
	}

	return nextStep.occupied
}

// Does a single step on a seatMap and returns whether any changes were made
func stepSeatMap(seatMap *SeatMap, count CountNeighbours, neighbourLimit int) (*SeatMap, bool) {
	newMap := &SeatMap{
		seats:    make([][]rune, len(seatMap.seats)),
		occupied: 0,
	}
	changed := false

	for i := 0; i < len(seatMap.seats); i++ {
		newMap.seats[i] = make([]rune, len(seatMap.seats[i]))
		for j := 0; j < len(seatMap.seats[i]); j++ {
			if seatMap.seats[i][j] == '.' {
				newMap.seats[i][j] = '.'
				continue
			}

			occupiedNeighbours := count(seatMap, i, j)
			if seatMap.seats[i][j] == 'L' {
				if occupiedNeighbours == 0 {
					newMap.seats[i][j] = '#'
					changed = true
					newMap.occupied++
				} else {
					newMap.seats[i][j] = 'L'
				}
			} else if seatMap.seats[i][j] == '#' {
				if occupiedNeighbours < neighbourLimit {
					newMap.seats[i][j] = '#'
					newMap.occupied++
				} else {
					newMap.seats[i][j] = 'L'
					changed = true
				}
			}
		}
	}

	return newMap, !changed
}

// Calculates the number of occupied seats adjacent to the one with the given x, y coordinates
// Implements CountNeighbours
func OccupiedNeighbours(seatMap *SeatMap, x, y int) int {
	occupied := 0

	for i := x - 1; i <= x+1; i++ {
		if i < 0 || i >= len(seatMap.seats) {
			continue
		}
		for j := y - 1; j <= y+1; j++ {
			if j < 0 || j >= len(seatMap.seats[i]) || (i == x && j == y) {
				continue
			}
			if seatMap.seats[i][j] == '#' {
				occupied++
			}
		}
	}

	return occupied
}

// Calculates the number of occupied seats visible from the one with the given x, y coordinates
// Implements CountNeighbours
func OccupiedVisibleNeighbours(seatMap *SeatMap, x, y int) int {
	occupied := 0

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if nextSeatOccupied(seatMap, x, y, i, j) {
				occupied++
			}
		}
	}

	return occupied
}

// Returns the number of occupied visible seats from the given seat in the direction
// specified by the step values. An empty seat blocks the view of occupied seats
func nextSeatOccupied(seatMap *SeatMap, x, y, stepX, stepY int) bool {
	i := x + stepX
	j := y + stepY

	for i >= 0 && i < len(seatMap.seats) && j >= 0 && j < len(seatMap.seats[i]) {
		if seatMap.seats[i][j] == 'L' {
			return false
		} else if seatMap.seats[i][j] == '#' {
			return true
		}
		i += stepX
		j += stepY
	}
	return false
}
