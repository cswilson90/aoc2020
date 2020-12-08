package day6

// Given a list of group answers returns the sum of all yes answered questions per group
func SumAnswerCounts(answerGroups [][]string) int {
	sumCounts := 0

	for _, answerGroup := range answerGroups {
		groupAnswers := make(map[rune]bool)
		for _, answerSet := range answerGroup {
			for _, letter := range answerSet {
				if letter >= 'a' && letter <= 'z' {
					groupAnswers[letter] = true
				}
			}
		}
		for _ = range groupAnswers {
			sumCounts++
		}
	}

	return sumCounts
}

// Given a list of group answers returns the sum of questions answered yes by all people in each group
func SumAllAnswerCounts(answerGroups [][]string) int {
	sumCounts := 0

	for _, answerGroup := range answerGroups {
		numPeople := 1
		groupAnswers := make(map[rune]int)
		for _, answerSet := range answerGroup {
			for _, letter := range answerSet {
				if letter >= 'a' && letter <= 'z' {
					_, ok := groupAnswers[letter]
					if !ok {
						groupAnswers[letter] = 1
					} else {
						groupAnswers[letter]++
					}
				} else if letter == '\n' {
					numPeople++
				}
			}
		}
		for _, count := range groupAnswers {
			if numPeople == count {
				sumCounts++
			}
		}
	}

	return sumCounts
}
