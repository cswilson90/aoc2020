package day16

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

var (
	RangeSplit = regexp.MustCompile(`^(\d+)-(\d+)$`)
	RuleSplit  = regexp.MustCompile(`^(.*): (.*) or (.*)$`)
)

type (
	Rules   map[string]*Rule
	Ticket  []int
	Tickets []Ticket
)

type Rule struct {
	field  string
	ranges []*Range
}

type Range struct {
	lowerBound, upperBound int
}

func ParseRules(ruleStrings []string) Rules {
	rules := make(Rules)

	for _, ruleString := range ruleStrings {
		rule := parseRule(ruleString)
		rules[rule.field] = rule
	}

	return rules
}

// Calculates the ticket scanning error for part 1
func TicketScanningError(tickets []string, rules Rules) int {
	sum := 0
	for _, ticketString := range tickets {
		ticket := ParseTicket(ticketString)
		for _, v := range ticket {
			if len(rules.ruleMatches(v)) == 0 {
				sum += v
			}
		}
	}

	return sum
}

func FindFieldsInTickets(ticketStrings []string, rules Rules) map[string]int {
	fieldMapping := make(map[string]int)
	tickets := make(Tickets, 0)

	// Get list of valid tickets
	for _, ticketString := range ticketStrings {
		ticket := ParseTicket(ticketString)
		valid := true
		for _, v := range ticket {
			if len(rules.ruleMatches(v)) == 0 {
				valid = false
				break
			}
		}
		if valid {
			tickets = append(tickets, ticket)
		}
	}

	// Stores a list of possible ticket field numbers for each field name
	possibleFields := make(map[string]map[int]struct{})
	for ruleField, _ := range rules {
		possibleFields[ruleField] = make(map[int]struct{}, 0)
	}

	for ticketField := 0; ticketField < len(tickets[0]); ticketField++ {
		matchingFields := make([]string, 0)

		for fieldName, _ := range possibleFields {
			allMatch := true

			for ticketNum := 0; ticketNum < len(tickets); ticketNum++ {
				value := tickets[ticketNum][ticketField]
				if !rules[fieldName].matches(value) {
					allMatch = false
					break
				}
			}

			if allMatch {
				matchingFields = append(matchingFields, fieldName)
			}
		}

		// If only one field matches then we know that field matches the ticket field index
		// We can then store this and skip the rule for other indexes
		if len(matchingFields) == 1 {
			fieldMapping[matchingFields[0]] = ticketField
			delete(possibleFields, matchingFields[0])
		} else {
			for _, field := range matchingFields {
				possibleFields[field][ticketField] = struct{}{}
			}
		}
	}

	// While we still have fields we haven't figured out look for the field with only
	// one option, once found remove that option for other fields
	for len(possibleFields) > 0 {
		deleteIndex := -1
		for fieldName, possibleIndexes := range possibleFields {
			if len(possibleIndexes) == 1 {
				fieldIndex := 0
				for k := range possibleIndexes {
					fieldIndex = k
				}

				fieldMapping[fieldName] = fieldIndex
				delete(possibleFields, fieldName)
				deleteIndex = fieldIndex
				break
			}
		}

		if deleteIndex == -1 {
			log.Fatalf("Found no index to delete from field list")
		}

		for fieldName := range possibleFields {
			delete(possibleFields[fieldName], deleteIndex)
		}
	}

	return fieldMapping
}

func ParseTicket(ticketString string) Ticket {
	stringValues := strings.Split(ticketString, ",")
	ticket := make(Ticket, len(stringValues))
	for i, v := range stringValues {
		value, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalf("Ticket %v: error parsing value %v", ticketString, v)
		}

		ticket[i] = value
	}

	return ticket
}

func (r Rules) ruleMatches(value int) []*Rule {
	matched := make([]*Rule, 0)

	for _, rule := range r {
		if rule.matches(value) {
			matched = append(matched, rule)
		}
	}

	return matched
}

func parseRule(ruleString string) *Rule {
	matches := RuleSplit.FindStringSubmatch(ruleString)
	if matches == nil {
		log.Fatalf("Rule %v does not match regex", ruleString)
	}

	ranges := []*Range{parseRange(matches[2]), parseRange(matches[3])}

	return &Rule{matches[1], ranges}
}

func (r *Rule) matches(num int) bool {
	for _, v := range r.ranges {
		if v.contains(num) {
			return true
		}
	}

	return false
}

func parseRange(rangeString string) *Range {
	matches := RangeSplit.FindStringSubmatch(rangeString)
	if matches == nil {
		log.Fatalf("Range %v does not match regex", rangeString)
	}

	lowerBound, err := strconv.Atoi(matches[1])
	if err != nil {
		log.Fatalf("Range %v: lower bound doesn't convert to int", rangeString)
	}
	upperBound, err := strconv.Atoi(matches[2])
	if err != nil {
		log.Fatalf("Range %v: upper bound doesn't convert to int", rangeString)
	}

	return &Range{lowerBound, upperBound}
}

func (r *Range) contains(num int) bool {
	return num >= r.lowerBound && num <= r.upperBound
}
