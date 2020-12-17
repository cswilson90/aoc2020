package day17

// Stores the points of all active cubes
type CubeMap map[Point]bool

type Point interface {
	getAdjacent() []Point
}

type Point3D struct {
	x, y, z int
}

type Point4D struct {
	x, y, z, w int
}

// Runs the boot process for the input in 3D and returns the number of active cubes after boot is complete.
func RunBootProcess3D(input []string) int {
	cubeMap := IntialiseCubeMap3D(input)
	return runBootProcess(cubeMap)
}

// Runs the boot process for the input in 4D and returns the number of active cubes after boot is complete.
func RunBootProcess4D(input []string) int {
	cubeMap := IntialiseCubeMap4D(input)
	return runBootProcess(cubeMap)
}

func runBootProcess(cubeMap CubeMap) int {
	for i := 1; i <= 6; i++ {
		cubeMap = runCycle(cubeMap)
	}

	return len(cubeMap)
}

// Run a single cycle on a CubeMap
func runCycle(cubeMap CubeMap) CubeMap {
	newMap := make(CubeMap)
	// Caches non active points we've already calculated
	seenNonActive := make(map[Point]bool)

	for point := range cubeMap {
		for _, adjacentPoint := range point.getAdjacent() {
			// Skip points we've already calculated
			if newMap[adjacentPoint] || seenNonActive[adjacentPoint] {
				continue
			}

			// Work out if cube is active in next step
			activeAdjacent := cubeMap.activeAdjacentCubes(adjacentPoint)
			if cubeMap[adjacentPoint] && (activeAdjacent == 2 || activeAdjacent == 3) {
				newMap[adjacentPoint] = true
			} else if !cubeMap[adjacentPoint] && activeAdjacent == 3 {
				newMap[adjacentPoint] = true
			} else {
				seenNonActive[adjacentPoint] = false
			}
		}
	}

	return newMap
}

func (p Point3D) getAdjacent() []Point {
	adjacent := make([]Point, 0)
	for i := p.x - 1; i <= p.x+1; i++ {
		for j := p.y - 1; j <= p.y+1; j++ {
			for k := p.z - 1; k <= p.z+1; k++ {
				point := Point3D{i, j, k}
				if p != point {
					adjacent = append(adjacent, point)
				}
			}
		}
	}
	return adjacent
}

func (p Point4D) getAdjacent() []Point {
	adjacent := make([]Point, 0)
	for i := p.x - 1; i <= p.x+1; i++ {
		for j := p.y - 1; j <= p.y+1; j++ {
			for k := p.z - 1; k <= p.z+1; k++ {
				for l := p.w - 1; l <= p.w+1; l++ {
					point := Point4D{i, j, k, l}
					if p != point {
						adjacent = append(adjacent, point)
					}
				}
			}
		}
	}
	return adjacent
}

func (m CubeMap) activeAdjacentCubes(point Point) int {
	allAdjacent := point.getAdjacent()
	active := 0
	for _, point := range allAdjacent {
		if m[point] {
			active++
		}
	}
	return active
}

func IntialiseCubeMap3D(input []string) CubeMap {
	cubeMap := make(CubeMap)

	for x, rowString := range input {
		for y, value := range rowString {
			if value == '#' {
				cubeMap[Point3D{x, y, 0}] = true
			}
		}
	}

	return cubeMap
}

func IntialiseCubeMap4D(input []string) CubeMap {
	cubeMap := make(CubeMap)

	for x, rowString := range input {
		for y, value := range rowString {
			if value == '#' {
				cubeMap[Point4D{x, y, 0, 0}] = true
			}
		}
	}

	return cubeMap
}
