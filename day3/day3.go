package day3

import "fmt"

// Generates a 2D bool map of tree locations (true is a tree).
func GenerateTreeMap(treeStrings []string) ([][]bool, error) {
	treeMap := make([][]bool, len(treeStrings))
	rowLength := len(treeStrings[0])

	for i, treeRow := range treeStrings {
		if len(treeRow) != rowLength {
			return nil, fmt.Errorf("Row %v of tree map input has different length to first row", i)
		}

		treeMap[i] = make([]bool, rowLength)
		for j, v := range treeRow {
			if v == '.' {
				treeMap[i][j] = false
			} else if v == '#' {
				treeMap[i][j] = true
			} else {
				return nil, fmt.Errorf("Unexpected character '%v' on line %v col %v of tree map input", v, i, j)
			}
		}
	}

	return treeMap, nil
}

// Caluclates the number of tree whilst moving down the tree maps in the given steps.
func CalculateTreeHits(treeMap [][]bool, rightStep int, downStep int) int {
	treeHits := 0
	x, y := 0, 0

	mapHeight := len(treeMap)
	if mapHeight < 1 {
		return 0
	}

	mapWidth := len(treeMap[0])

	for y < len(treeMap) {
		if treeMap[y][x] {
			treeHits++
		}

		x += rightStep
		x = x % (mapWidth)

		y += downStep
	}

	return treeHits
}
