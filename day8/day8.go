package day8

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type MachineState struct {
	Accumulator      int
	Code             []*Operation
	NextInstruction  int
	DoneInstructions []int
}

type Operation struct {
	Operator string
	Operand  int
}

var validOperators = map[string]bool{
	"acc": true,
	"jmp": true,
	"nop": true,
}

var (
	EndOfCodeError          = errors.New("Code has finished executing")
	InstructionPointerError = errors.New("Instruction pointer no pointing at valid code line")
	InfiniteLoopError       = errors.New("Hit infinite loop in code")
)

func ParseMachineInstructions(instructions []string) (*MachineState, error) {
	instrucs := &MachineState{
		Code: make([]*Operation, 0),
	}

	for i, instruction := range instructions {
		parts := strings.Split(instruction, " ")
		if len(parts) != 2 {
			return nil, fmt.Errorf("Line %v: Error splitting instruction '%v', got %v parts, expected 2",
				i, instruction, len(parts))
		}

		if !validOperator(parts[0]) {
			return nil, fmt.Errorf("Line %v: Invalid operator '%v'", i, parts[0])
		}

		operand, err := strconv.Atoi(parts[1])
		if err != nil {
			return nil, fmt.Errorf("Line %v: Error parsing operand '%v': "+err.Error(), i, parts[1])
		}

		operation := &Operation{
			Operator: parts[0],
			Operand:  operand,
		}
		instrucs.Code = append(instrucs.Code, operation)
	}

	return instrucs, nil
}

func (m *MachineState) FindInfiniteLoopValue() (int, error) {
	err := m.ExecuteProgram()
	if err != nil {
		if err == InfiniteLoopError {
			return m.Accumulator, nil
		}
		return 0, err
	}
	return 0, fmt.Errorf("Program executed without an infinite loop")
}

func (m *MachineState) FixInfiniteLoop() (int, error) {
	err := m.ExecuteProgram()
	if err != nil {
		if err != InfiniteLoopError {
			return 0, err
		}
	} else {
		return 0, fmt.Errorf("Program executed without an infinite loop")
	}

	executedInstructions := m.DoneInstructions
	for _, v := range executedInstructions {
		oldInstruc := m.Code[v]
		if oldInstruc.Operator != "jmp" && oldInstruc.Operator != "nop" {
			continue
		}

		newInstruc := switchJmpNop(oldInstruc)
		m.Code[v] = newInstruc

		err := m.ExecuteProgram()
		if err != nil {
			if err != InfiniteLoopError {
				return 0, err
			}
		} else {
			return m.Accumulator, nil
		}

		m.Code[v] = oldInstruc
	}

	return 0, fmt.Errorf("Found no fix for infinte loop")
}

func (m *MachineState) ExecuteProgram() error {
	m.Reset()

	seenInstrucs := make(map[int]bool)
	seenInstrucs[m.NextInstruction] = true

	for {
		err := m.ExecuteNext()
		if err != nil {
			return err
		}

		if m.NextInstruction < 0 {
			return nil
		}

		_, ok := seenInstrucs[m.NextInstruction]
		if ok {
			return InfiniteLoopError
		} else {
			seenInstrucs[m.NextInstruction] = true
		}
	}

	return nil
}

func (m *MachineState) Reset() {
	m.Accumulator = 0
	m.NextInstruction = 0
	m.DoneInstructions = make([]int, 0)
}

func (m *MachineState) ExecuteNext() error {
	if m.NextInstruction < 0 {
		return EndOfCodeError
	}

	if m.NextInstruction >= len(m.Code) {
		return InstructionPointerError
	}

	m.DoneInstructions = append(m.DoneInstructions, m.NextInstruction)
	nextInstruction := m.Code[m.NextInstruction]
	switch nextInstruction.Operator {
	case "acc":
		m.Accumulator += nextInstruction.Operand
		m.NextInstruction++
	case "nop":
		m.NextInstruction++
	case "jmp":
		m.NextInstruction += nextInstruction.Operand
	}

	if m.NextInstruction >= len(m.Code) {
		m.NextInstruction = -1
	}

	return nil
}

func switchJmpNop(op *Operation) *Operation {
	operator := "jmp"
	if op.Operator == "jmp" {
		operator = "nop"
	}
	return &Operation{
		Operator: operator,
		Operand:  op.Operand,
	}
}

func validOperator(opCode string) bool {
	_, ok := validOperators[opCode]
	return ok
}
