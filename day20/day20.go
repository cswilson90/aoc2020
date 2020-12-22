package day20

import (
	"log"
	"math"
	"regexp"
	"strconv"
)

const (
	Top    = 0
	Right  = 1
	Bottom = 2
	Left   = 3
)

var (
	TileIdSplit = regexp.MustCompile(`^Tile (\d+):$`)
	SeaMonster  = SeaMonsterTemplate{
		Point{1, 0},
		Point{2, 1},
		Point{2, 4},
		Point{1, 5},
		Point{1, 6},
		Point{2, 7},
		Point{2, 10},
		Point{1, 11},
		Point{1, 12},
		Point{2, 13},
		Point{2, 16},
		Point{1, 17},
		Point{1, 18},
		Point{1, 19},
		Point{0, 18},
	}
)

type (
	BorderMap          map[string][]int
	Picture            [][]rune
	PictureTiles       [][]*Tile
	SeaMonsterTemplate []Point
	Tiles              map[int]*Tile
)

type TileSet struct {
	tiles   Tiles
	borders BorderMap
}

type Tile struct {
	Id      int
	Pixels  [][]rune
	Borders []string
}

type Point struct {
	x, y int
}

func ParseTiles(tilesInput [][]string) *TileSet {
	tiles := make(Tiles)
	borderMap := make(BorderMap)

	for i, tile := range tilesInput {
		idMatches := TileIdSplit.FindStringSubmatch(tile[0])
		if idMatches == nil {
			log.Fatalf("Error parsing title '%v' from tile %v", tile[0], i)
		}
		tileId, err := strconv.Atoi(idMatches[1])
		if err != nil {
			log.Fatalf(err.Error())
		}
		image := tile[1:]

		pixels := make([][]rune, len(image))
		for r, imageRow := range image {
			pixels[r] = []rune(imageRow)
		}
		borders := getBorders(pixels)

		for _, border := range borders {
			_, ok := borderMap[border]
			if ok {
				borderMap[border] = append(borderMap[border], tileId)
			} else {
				borderMap[border] = []int{tileId}
			}
		}

		tiles[tileId] = &Tile{tileId, pixels, borders}
	}

	return &TileSet{tiles, borderMap}
}

func FindAnswers(tileSet *TileSet) (int, int) {
	pictureTiles := BuildTiles(tileSet)

	maxIndex := len(pictureTiles) - 1
	corners := pictureTiles[0][0].Id
	corners *= pictureTiles[0][maxIndex].Id
	corners *= pictureTiles[maxIndex][0].Id
	corners *= pictureTiles[maxIndex][maxIndex].Id

	picture := BuildPicture(pictureTiles)
	chopiness := findChopiness(picture)

	return corners, chopiness
}

func BuildTiles(tileSet *TileSet) PictureTiles {
	picSize := int(math.Sqrt(float64(len(tileSet.tiles))))
	picture := make(PictureTiles, picSize)
	for i := range picture {
		picture[i] = make([]*Tile, picSize)
	}

	picture, ok := findTileForPlace(picture, 0, 0, tileSet, make(map[int]bool))
	if !ok {
		log.Fatalf("Failed to build picture")
	}
	return picture
}

func BuildPicture(pictureTiles PictureTiles) Picture {
	tileSize := len(pictureTiles[0][0].Pixels)
	numTiles := len(pictureTiles)
	// Picture size is all tiles minus the stripped borders
	picSize := numTiles * (tileSize - 2)

	picture := make(Picture, picSize)
	for i := range picture {
		picture[i] = make([]rune, picSize)
	}

	for i, tileRow := range pictureTiles {
		for j, tile := range tileRow {
			for k := 1; k < tileSize-1; k++ {
				for l := 1; l < tileSize-1; l++ {
					modifierX := (tileSize - 2) * i
					modifierY := (tileSize - 2) * j
					picture[modifierX+k-1][modifierY+l-1] = tile.Pixels[k][l]
				}
			}
		}
	}

	return picture
}

func findChopiness(picture Picture) int {
	chopiness := 0

	for _, picRow := range picture {
		for _, value := range picRow {
			if value == '#' {
				chopiness++
			}
		}
	}

	for _, picVar := range pictureVariations(picture) {
		monsterPoints := make(map[Point]bool)
		for i, picRow := range picVar {
			for j := range picRow {
				if monsterPoints[Point{i, j}] {
					continue
				}

				newMonsterPoints, ok := findMonsterPoints(picVar, i, j)
				if ok {
					for _, point := range newMonsterPoints {
						monsterPoints[point] = true
					}
				}
			}
		}

		if len(monsterPoints) > 0 {
			chopiness -= len(monsterPoints)
			break
		}
	}

	return chopiness
}

func findMonsterPoints(picture Picture, x, y int) ([]Point, bool) {
	monsterPoints := make([]Point, 0)
	for _, point := range SeaMonster {
		picX := x + point.x
		picY := y + point.y

		if picX >= len(picture) || picY >= len(picture) {
			return nil, false
		}

		if picture[picX][picY] == '#' {
			monsterPoints = append(monsterPoints, Point{picX, picY})
		} else {
			return nil, false
		}
	}

	return monsterPoints, true
}

func findTileForPlace(picture PictureTiles, x, y int, tileSet *TileSet, used map[int]bool) (PictureTiles, bool) {
	var tileAbove, tileLeft *Tile
	if x > 0 {
		tileAbove = picture[x-1][y]
	}
	if y > 0 {
		tileLeft = picture[x][y-1]
	}

	nextX := x + 1
	nextY := y
	if nextX >= len(picture) {
		nextX = 0
		nextY = y + 1
	}

	for _, tile := range tileSet.tiles {
		if used[tile.Id] {
			continue
		}

		for _, tileVar := range tileVariations(tile) {
			// If theres a tile above check the top border matches otherwise check
			// the top border matches no other tile
			if tileAbove != nil {
				if tileAbove.Borders[Bottom] != tileVar.Borders[Top] {
					continue
				}
			} else if anotherMatchingBorder(tileVar.Borders[Top], tileSet.borders) {
				continue
			}

			// If theres a tile to the left check the left border matches otherwise check
			// the left border matches no other tile
			if tileLeft != nil {
				if tileLeft.Borders[Right] != tileVar.Borders[Left] {
					continue
				}
			} else if anotherMatchingBorder(tileVar.Borders[Left], tileSet.borders) {
				continue
			}

			picture[x][y] = tileVar
			used[tile.Id] = true

			// If we've reached the end then the picture is complete
			if nextY >= len(picture) {
				return picture, true
			}

			// Try and find the tile for the next position
			varPicture, ok := findTileForPlace(picture, nextX, nextY, tileSet, used)
			if ok {
				// Picture complete, return it
				return varPicture, true
			}

			// Picture not complete keep trying
			used[tile.Id] = false
			picture[x][y] = nil
		}
	}

	return picture, false
}

func anotherMatchingBorder(border string, borders BorderMap) bool {
	return len(borders[border])+len(borders[reverse(border)]) > 1
}

func tileVariations(tile *Tile) []*Tile {
	variations := make([]*Tile, 0)
	for i := 0; i <= 3; i++ {
		for _, flipVertical := range []bool{true, false} {
			for _, flipHorizontal := range []bool{true, false} {
				variations = append(variations, tile.transform(i, flipVertical, flipHorizontal))
			}
		}
	}
	return variations
}

func pictureVariations(picture Picture) []Picture {
	variations := make([]Picture, 0)
	for i := 0; i <= 3; i++ {
		for _, flipVertical := range []bool{true, false} {
			for _, flipHorizontal := range []bool{true, false} {
				variations = append(variations, Picture(transformPixels(picture, i, flipVertical, flipHorizontal)))
			}
		}
	}
	return variations
}

func (t *Tile) transform(rotation int, flipVert, flipHori bool) *Tile {
	newPixels := transformPixels(t.Pixels, rotation, flipVert, flipHori)
	return &Tile{
		Id:      t.Id,
		Pixels:  newPixels,
		Borders: getBorders(newPixels),
	}
}

func transformPixels(pixels [][]rune, rotation int, flipVert, flipHori bool) [][]rune {
	size := len(pixels)
	newPixels := make([][]rune, size)

	// RotatePixels
	for i := 0; i < size; i++ {
		newPixels[i] = make([]rune, size)
		for j := 0; j < size; j++ {
			switch rotation {
			case 1:
				newPixels[i][j] = pixels[j][size-1-i]
			case 2:
				newPixels[i][j] = pixels[size-1-i][size-1-j]
			case 3:
				newPixels[i][j] = pixels[size-1-j][i]
			default:
				newPixels[i][j] = pixels[i][j]
			}
		}
	}

	if flipHori {
		for i := 0; i < size; i++ {
			for j := 0; j < size/2; j++ {
				flipJ := size - 1 - j
				oldValue := newPixels[i][j]
				newPixels[i][j] = newPixels[i][flipJ]
				newPixels[i][flipJ] = oldValue
			}
		}
	}

	if flipVert {
		for i := 0; i < size/2; i++ {
			for j := 0; j < size; j++ {
				flipI := size - 1 - i
				oldValue := newPixels[i][j]
				newPixels[i][j] = newPixels[flipI][j]
				newPixels[flipI][j] = oldValue
			}
		}
	}

	return newPixels
}

// Returns the borders of the given pixel grid
func getBorders(pixels [][]rune) []string {
	borders := make([]string, 4)

	for r, row := range pixels {
		if r == 0 {
			borders[Top] = string(row)
		} else if r == len(pixels)-1 {
			borders[Bottom] = string(row)
		}

		borders[Right] += string([]rune{row[len(row)-1]})
		borders[Left] += string([]rune{row[0]})
	}

	return borders
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
