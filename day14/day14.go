package day14

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	MaskInstruction   = regexp.MustCompile(`^mask = ([X01]+$)`)
	MemoryInstruction = regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
)

type BitMask interface {
	// Applies the bitmap to the instruction and stores the result in the memory
	apply(*Instruction, *Memory)

	// Parses a bit mask string and sets up state so future calls to apply()
	// will use the mask
	parse(string) error
}

type BitMaskV1 struct {
	andMask uint64
	orMask  uint64
}

type BitMaskV2 struct {
	orMask        uint64
	floatingMasks []uint64
}

type Memory struct {
	memory map[uint64]uint64
}

type Instruction struct {
	Address uint64
	Value   uint64
}

// Runs the docking code instructions using a bitmask
// The bitmask can either be BitMaskV1 or BitMaskV2
func RunDockingProgram(instructions []string, bitMask BitMask) (uint64, error) {
	memory := &Memory{
		memory: make(map[uint64]uint64),
	}

	for i, instructionString := range instructions {
		maskMatches := MaskInstruction.FindStringSubmatch(instructionString)
		if maskMatches != nil {
			err := bitMask.parse(maskMatches[1])
			if err != nil {
				return 0, fmt.Errorf("Error parsing bitmask on line %v:"+err.Error(), i+1)
			}
			continue
		}

		instruc, err := parseInstruction(instructionString)
		if err != nil {
			return 0, fmt.Errorf("Error parsing instruction on line %v:"+err.Error(), i+1)
		}
		bitMask.apply(instruc, memory)
	}

	sum := uint64(0)
	for _, v := range memory.memory {
		sum += v
	}
	return sum, nil
}

func parseInstruction(instruction string) (*Instruction, error) {
	memMatches := MemoryInstruction.FindStringSubmatch(instruction)
	if memMatches == nil {
		return nil, fmt.Errorf("Failed to match instruction")
	}

	memAddress, err := strconv.ParseUint(memMatches[1], 10, 64)
	if err != nil {
		return nil, err
	}
	value, err := strconv.ParseUint(memMatches[2], 10, 64)
	if err != nil {
		return nil, err
	}

	return &Instruction{
		Address: memAddress,
		Value:   value,
	}, nil
}

func (b *BitMaskV1) apply(instruction *Instruction, memory *Memory) {
	memory.memory[instruction.Address] = (instruction.Value & b.andMask) | b.orMask
}

func (b *BitMaskV1) parse(bitMask string) error {
	var err error

	andMaskString := strings.ReplaceAll(bitMask, "X", "1")
	b.andMask, err = strconv.ParseUint(andMaskString, 2, 36)
	if err != nil {
		return err
	}

	b.orMask, err = parseOrMask(bitMask)
	if err != nil {
		return err
	}

	return nil
}

func (b *BitMaskV2) apply(instruction *Instruction, memory *Memory) {
	newAddress := instruction.Address | b.orMask

	addresses := applyFloatingMasks(b.floatingMasks, newAddress)
	for _, address := range addresses {
		memory.memory[address] = instruction.Value
	}
}

// Calculates all possible transformations of a value by applying the list
// of floating masks
func applyFloatingMasks(floatingMasks []uint64, value uint64) []uint64 {
	if len(floatingMasks) == 0 {
		return []uint64{value}
	}

	// Get the two variations of the value setting the first mask's bit to 0 and 1
	oneValue := value | floatingMasks[0]
	zeroValue := value & ^floatingMasks[0]

	// If there's only the one floating mask we can return the two values
	if len(floatingMasks) == 1 {
		return []uint64{oneValue, zeroValue}
	}

	allResults := make([]uint64, 0)

	// Recursively apply all remaining maps and collect all the possible values
	nextMasks := floatingMasks[1:]
	allResults = append(allResults, applyFloatingMasks(nextMasks, oneValue)...)
	allResults = append(allResults, applyFloatingMasks(nextMasks, zeroValue)...)

	return allResults
}

func (b *BitMaskV2) parse(bitMask string) error {
	var err error
	b.orMask, err = parseOrMask(bitMask)
	if err != nil {
		return err
	}

	b.floatingMasks = make([]uint64, 0)
	for i, v := range bitMask {
		if v == 'X' {
			mask := uint64(1) << (len(bitMask) - i - 1)
			b.floatingMasks = append(b.floatingMasks, mask)
		}
	}

	return nil
}

func parseOrMask(bitMask string) (uint64, error) {
	orMaskString := strings.ReplaceAll(bitMask, "X", "0")
	orMask, err := strconv.ParseUint(orMaskString, 2, 36)
	if err != nil {
		return 0, err
	}

	return orMask, nil
}
