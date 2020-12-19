package day19

import (
	"log"
	"regexp"
	"strings"
)

var RuleSplit = regexp.MustCompile(`^(\d+): (.*)$`)
var StringRuleMatch = regexp.MustCompile(`^"(\w)"$`)

type Rule interface {
	buildRegexString(Rules) string
}

type CompositeRule struct {
	id       string
	SubRules [][]string
}

type (
	Rules      map[string]Rule
	StringRule string
)

func ParseRules(ruleStrings []string) Rules {
	rules := make(Rules)

	for _, ruleString := range ruleStrings {
		ruleMatches := RuleSplit.FindStringSubmatch(ruleString)
		if ruleMatches == nil {
			log.Fatalf("Rule '%v' doesn't match regex", ruleString)
		}

		ruleId := ruleMatches[1]
		var rule Rule

		stringRuleMatches := StringRuleMatch.FindStringSubmatch(ruleMatches[2])
		if stringRuleMatches != nil {
			rule = StringRule(stringRuleMatches[1])
		} else {
			subRules := strings.Split(ruleMatches[2], " | ")
			newRule := &CompositeRule{ruleId, make([][]string, 0)}
			for _, subRule := range subRules {
				ruleIDs := strings.Split(subRule, " ")
				newRule.SubRules = append(newRule.SubRules, ruleIDs)
			}
			rule = newRule
		}

		rules[ruleId] = rule
	}

	return rules
}

// Get number of matching messages for part 1
func MatchingMessages(messages []string, rules Rules) int {
	rule0, ok := rules["0"]
	if !ok {
		log.Fatalf("Found no rule 0")
	}

	regexString := "^" + rule0.buildRegexString(rules) + "$"
	messageMatch, err := regexp.Compile(regexString)
	if err != nil {
		log.Fatalf(err.Error())
	}

	matches := 0
	for _, message := range messages {
		if messageMatch.MatchString(message) {
			matches++
		}
	}

	return matches
}

// Get number of matching messages for part 2
// Uses fact rules are <42>+<42>{n}<31>{n}
// Counts matches of rule 42 at front and 31 at end of message
func MatchingMessages2(messages []string, rules Rules) int {
	regex42, err := regexp.Compile("^" + rules["42"].buildRegexString(rules))
	if err != nil {
		log.Fatalf(err.Error())
	}

	regex31, err := regexp.Compile(rules["31"].buildRegexString(rules) + "$")
	if err != nil {
		log.Fatalf(err.Error())
	}

	matches := 0
	for _, message := range messages {
		num42 := 0
		num31 := 0

		for {
			oldMessage := message
			message = regex42.ReplaceAllString(message, "")
			if oldMessage == message {
				break
			}
			num42++
		}

		for {
			oldMessage := message
			message = regex31.ReplaceAllString(message, "")
			if oldMessage == message {
				break
			}
			num31++
		}

		// Matches if no message left, both rules matched and there are enough matches of each
		if len(message) == 0 && num42 > 1 && num31 > 0 && num42 > num31 {
			matches++
		}
	}

	return matches
}

func (r *CompositeRule) buildRegexString(rules Rules) string {
	orStrings := make([]string, len(r.SubRules))

	for i, subRule := range r.SubRules {
		ruleString := ""
		for _, ruleID := range subRule {
			ruleString += rules[ruleID].buildRegexString(rules)
		}
		orStrings[i] = ruleString
	}

	if len(orStrings) == 1 {
		return orStrings[0]
	}

	return "(" + strings.Join(orStrings, "|") + ")"
}

func (r StringRule) buildRegexString(rules Rules) string {
	return string(r)
}
