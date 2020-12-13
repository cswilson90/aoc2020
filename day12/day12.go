package day12

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

var SplitInstruction = regexp.MustCompile(`^([NESWLRF])(\d+)$`)

type Ship interface {
	manhattanDistance() int
	move(int, int)
	moveForward(int)
	rotate(int)
}

// Represents a ship moved by it's own direction
// Implements Ship
type DirectionShip struct {
	XPos      int
	YPos      int
	Direction int
}

// Represents a ship moved by a waypoint
// Implements Ship
type WaypointShip struct {
	ShipX     int
	ShipY     int
	WaypointX int
	WaypointY int
}

type Instruction struct {
	Action string
	Value  int
}

// Applies instructions to a ship and returns the manhattan distance of the ship
// after the instructions have been applied
func ManhattanDistance(instructions []string, ship Ship) (int, error) {
	for i, v := range instructions {
		instruction, err := parseInstruction(v)
		if err != nil {
			return 0, fmt.Errorf("Error parsing instruction %v: "+err.Error(), i)
		}
		err = applyInstruction(instruction, ship)
		if err != nil {
			return 0, fmt.Errorf("Error applying instruction %v: "+err.Error(), i)
		}
	}

	return ship.manhattanDistance(), nil
}

// Parses an instruction stirng into an instruction object
func parseInstruction(instructionString string) (*Instruction, error) {
	matches := SplitInstruction.FindStringSubmatch(instructionString)
	if matches == nil {
		return nil, fmt.Errorf("%v does not match instruction regex", instructionString)
	}

	value, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil, fmt.Errorf("%v non int value", instructionString)
	}

	return &Instruction{
		Action: matches[1],
		Value:  value,
	}, nil
}

// Applies a single instruction to a ship
func applyInstruction(instruction *Instruction, ship Ship) error {
	switch instruction.Action {
	case "N":
		ship.move(0, instruction.Value)
	case "S":
		ship.move(0, -instruction.Value)
	case "E":
		ship.move(instruction.Value, 0)
	case "W":
		ship.move(-instruction.Value, 0)
	case "L":
		ship.rotate(-instruction.Value)
	case "R":
		ship.rotate(instruction.Value)
	case "F":
		ship.moveForward(instruction.Value)
	default:
		return fmt.Errorf("Unknown instruction action %v", instruction.Action)
	}

	return nil
}

func (ship *DirectionShip) move(x, y int) {
	ship.XPos += x
	ship.YPos += y
}

func (ship *DirectionShip) moveForward(distance int) {
	switch ship.Direction {
	case 0:
		ship.move(0, distance)
	case 90:
		ship.move(distance, 0)
	case 180:
		ship.move(0, -distance)
	case 270:
		ship.move(-distance, 0)
	default:
		log.Fatalf("Unsupported direction: %v", ship.Direction)
	}
}

func (ship *DirectionShip) rotate(angle int) {
	value := ship.Direction + angle
	if value < 0 {
		value = 360 + value
	}
	ship.Direction = value % 360
}

func (ship *DirectionShip) manhattanDistance() int {
	return abs(ship.XPos) + abs(ship.YPos)
}

func (ship *WaypointShip) move(x, y int) {
	ship.WaypointX += x
	ship.WaypointY += y
}

func (ship *WaypointShip) moveForward(steps int) {
	for i := 0; i < steps; i++ {
		ship.ShipX += ship.WaypointX
		ship.ShipY += ship.WaypointY
	}
}

func (ship *WaypointShip) rotate(angle int) {
	if angle < 0 {
		angle = 360 + angle
	}

	oldX := ship.WaypointX
	oldY := ship.WaypointY

	switch angle {
	case 90:
		ship.WaypointX = oldY
		ship.WaypointY = -oldX
	case 180:
		ship.WaypointX = -oldX
		ship.WaypointY = -oldY
	case 270:
		ship.WaypointX = -oldY
		ship.WaypointY = oldX
	case 360:
	default:
		log.Fatalf("Unsupported angle: %v", angle)
	}
}

func (ship *WaypointShip) manhattanDistance() int {
	return abs(ship.ShipX) + abs(ship.ShipY)
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
