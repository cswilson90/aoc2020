package day22

import (
	"log"
	"strconv"
)

type Hand []int

type HandScores struct {
	score1, score2 int
}

// Plays a game of combat and returns the winner's score
func PlayCombat(player1Hand []string, player2Hand []string) int {
	hand1 := parseHand(player1Hand)
	hand2 := parseHand(player2Hand)

	var winner Hand
	for winner == nil {
		num1 := hand1[0]
		num2 := hand2[0]

		hand1 = hand1[1:]
		hand2 = hand2[1:]

		if num1 > num2 {
			hand1 = append(hand1, []int{num1, num2}...)
		} else {
			hand2 = append(hand2, []int{num2, num1}...)
		}

		if len(hand1) == 0 {
			winner = hand2
		} else if len(hand2) == 0 {
			winner = hand1
		}
	}

	return scoreHand(winner)
}

// Plays a game of recursive combat and returns the winner's score
func PlayRecursiveCombat(player1Hand []string, player2Hand []string) int {
	hand1 := parseHand(player1Hand)
	hand2 := parseHand(player2Hand)

	_, score := playRecursiveCombat(hand1, hand2)
	return score
}

// Playes a game of recurisve comabt and returns the winner and their score
func playRecursiveCombat(hand1, hand2 Hand) (int, int) {
	var winnerNum int
	seenScores := make(map[HandScores]bool)

	for winnerNum == 0 {
		handScores := handScores(hand1, hand2)
		if seenScores[handScores] {
			winnerNum = 1
			continue
		} else {
			seenScores[handScores] = true
		}

		num1 := hand1[0]
		num2 := hand2[0]

		hand1 = hand1[1:]
		hand2 = hand2[1:]

		roundWinner := 2
		if num1 <= len(hand1) && num2 <= len(hand2) {
			hand1Copy := make(Hand, num1)
			copy(hand1Copy, hand1)

			hand2Copy := make(Hand, num2)
			copy(hand2Copy, hand2)

			roundWinner, _ = playRecursiveCombat(hand1Copy, hand2Copy)
		} else if num1 > num2 {
			roundWinner = 1
		}

		if roundWinner == 1 {
			hand1 = append(hand1, []int{num1, num2}...)
		} else {
			hand2 = append(hand2, []int{num2, num1}...)
		}

		if len(hand1) == 0 {
			winnerNum = 2
		} else if len(hand2) == 0 {
			winnerNum = 1
		}
	}

	if winnerNum == 1 {
		return 1, scoreHand(hand1)
	}

	return 2, scoreHand(hand2)
}

func handScores(hand1, hand2 Hand) HandScores {
	return HandScores{scoreHand(hand1), scoreHand(hand2)}
}

func scoreHand(hand Hand) int {
	handSize := len(hand)
	score := 0
	for i, v := range hand {
		score += (v * (handSize - i))
	}
	return score
}

func parseHand(startHand []string) Hand {
	hand := make(Hand, 0)

	for i, v := range startHand {
		// Skip first row as it's player number
		if i == 0 {
			continue
		}

		value, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf(err.Error())
		}
		hand = append(hand, value)
	}

	return hand
}
