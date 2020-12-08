package day7

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type (
	BagColourMap map[string]*Bag
)

type BagLink struct {
	Colour string
	Number int
}

type Bag struct {
	Colour   string
	Contains []*BagLink
}

var (
	splitBagString   = regexp.MustCompile(`^(\w+ \w+) bags contain (.*)$`)
	splitLinksString = regexp.MustCompile(`^(\d+) (\w+ \w+) bags?[.]?$`)
)

// Counts the number of distinct bag colours that can contain a shiny gold bag
func CountColoursContainingGold(bagMap BagColourMap) int {
	colourCount := 0
	for _, bag := range bagMap {
		if containsShinyGold(bag, bagMap) {
			colourCount++
		}
	}

	return colourCount
}

// Counts the number of bags you must have inside a shiny gold bag
func InsideShinyGoldBagCount(bagMap BagColourMap) (int, error) {
	goldBag, ok := bagMap["shiny gold"]
	if !ok {
		return 0, fmt.Errorf("no shiny gold bag found in bag colour map")
	}

	return sumContainedBags(goldBag, bagMap), nil
}

func containsShinyGold(bag *Bag, bagMap BagColourMap) bool {
	for _, linkedBag := range bag.Contains {
		if linkedBag.Colour == "shiny gold" {
			return true
		}

		if containsShinyGold(bagMap[linkedBag.Colour], bagMap) {
			return true
		}
	}

	return false
}

func sumContainedBags(bag *Bag, bagMap BagColourMap) int {
	bagSum := 0

	for _, bagLink := range bag.Contains {
		bagSum += bagLink.Number
		bagSum += (bagLink.Number * sumContainedBags(bagMap[bagLink.Colour], bagMap))
	}

	return bagSum
}

func ParseBagStrings(bagStrings []string) (BagColourMap, error) {
	bagMap := make(BagColourMap)

	for _, v := range bagStrings {
		bag, err := parseBagString(v)
		if err != nil {
			return nil, err
		}
		bagMap[bag.Colour] = bag
	}

	return bagMap, nil
}

func parseBagString(bagString string) (*Bag, error) {
	matches := splitBagString.FindStringSubmatch(bagString)
	if matches == nil {
		return nil, fmt.Errorf("Could not parse bag string '%v'", bagString)
	}

	bag := &Bag{
		Colour:   matches[1],
		Contains: make([]*BagLink, 0),
	}

	links := matches[2]
	if links != "no other bags." {
		linkStrings := strings.Split(links, ", ")
		for _, v := range linkStrings {
			matches := splitLinksString.FindStringSubmatch(v)
			if matches == nil {
				return nil, fmt.Errorf("Could not parse link string '%v'", v)
			}

			numBags, err := strconv.Atoi(matches[1])
			if err != nil {
				return nil, fmt.Errorf("Error converting number '%v' in '%v':"+err.Error(), matches[1], v)
			}

			bagLink := &BagLink{
				Colour: matches[2],
				Number: numBags,
			}
			bag.Contains = append(bag.Contains, bagLink)
		}
	}

	return bag, nil
}
