package day6

// Given a list of group answers returns the sum of all yes answered questions per group
func SumAnswerCounts(answerGroups [][]string) int {
	sumCounts := 0

	for _, answerGroup := range answerGroups {
		sumCounts += len(getGroupAnswers(answerGroup))
	}

	return sumCounts
}

// Given a list of group answers returns the sum of questions answered yes by all people in each group
func SumAllAnswerCounts(answerGroups [][]string) int {
	sumCounts := 0

	for _, answerGroup := range answerGroups {
		for _, count := range getGroupAnswers(answerGroup) {
			// Increment if everyone in the answer group answered yes
			if len(answerGroup) == count {
				sumCounts++
			}
		}
	}

	return sumCounts
}

func getGroupAnswers(answerGroup []string) map[rune]int {
	groupAnswers := make(map[rune]int)

	for _, answerSet := range answerGroup {
		for _, letter := range answerSet {
			if letter >= 'a' && letter <= 'z' {
				groupAnswers[letter]++
			}
		}
	}

	return groupAnswers
}
