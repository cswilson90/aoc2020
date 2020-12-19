package day18

import (
	"log"
	"strconv"
	"strings"
)

type EquationState struct {
	acc      int
	operator string
}

type ParsingState struct {
	newTokens   []string
	accumulated []string
}

// Evaluates all equations and returns the sum of the results
func EvaluateAllEquations(equationStrings []string, complexRules bool) int {
	sum := 0

	for _, equationString := range equationStrings {
		sum += EvaluateEquation(equationString, complexRules)
	}

	return sum
}

// Evaluates a single equation
func EvaluateEquation(equationString string, complexRules bool) int {
	// Add spaces around brackets so we can split on space to tokenise
	equationString = strings.ReplaceAll(equationString, "(", "( ")
	equationString = strings.ReplaceAll(equationString, ")", " )")
	tokens := strings.Split(equationString, " ")

	if complexRules {
		tokens = insertAdditionBrackets(tokens)
	}

	state := &EquationState{}
	stateStack := make([]*EquationState, 0)

	for _, value := range tokens {
		switch value {
		case "+", "*":
			state.operator = value
		case "(":
			stateStack = append(stateStack, state)
			state = &EquationState{}
		case ")":
			bracketValue := state.acc
			state = stateStack[len(stateStack)-1]
			stateStack = stateStack[:len(stateStack)-1]
			state.evaluateInt(bracketValue)
		default:
			state.evaluateString(value)
		}
	}

	return state.acc
}

// Inserts brackets around "+" operations so they have higher precedence
func insertAdditionBrackets(equationTokens []string) []string {
	state := newParsingState()
	stateStack := make([]*ParsingState, 0)

	for _, token := range equationTokens {
		switch token {
		case "+":
			state.accumulated = append(state.accumulated, token)
		case "*":
			state.newTokens = append(state.newTokens, wrapBrackets(state.accumulated)...)
			state.newTokens = append(state.newTokens, "*")
			state.accumulated = state.accumulated[:0]
		case "(":
			stateStack = append(stateStack, state)
			state = newParsingState()
		case ")":
			state.newTokens = append(state.newTokens, wrapBrackets(state.accumulated)...)

			oldState := state
			state = stateStack[len(stateStack)-1]
			stateStack = stateStack[:len(stateStack)-1]

			state.accumulated = append(state.accumulated, wrapBrackets(oldState.newTokens)...)
		default:
			state.accumulated = append(state.accumulated, token)
		}
	}

	return append(state.newTokens, wrapBrackets(state.accumulated)...)
}

func newParsingState() *ParsingState {
	return &ParsingState{
		newTokens:   make([]string, 0),
		accumulated: make([]string, 0),
	}
}

func wrapBrackets(tokens []string) []string {
	newTokens := []string{"("}
	newTokens = append(newTokens, tokens...)
	return append(newTokens, ")")
}

func (s *EquationState) evaluateString(value string) {
	number, err := strconv.Atoi(value)
	if err != nil {
		log.Fatalf(err.Error())
	}
	s.evaluateInt(number)
}

func (s *EquationState) evaluateInt(number int) {
	switch s.operator {
	case "+":
		s.acc = s.acc + number
	case "*":
		s.acc = s.acc * number
	default:
		s.acc = number
	}
}
