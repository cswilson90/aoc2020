package day21

import (
	"log"
	"regexp"
	"sort"
	"strings"
)

var SplitFood = regexp.MustCompile(`^(.*) \(contains (.*)\)$`)

type IngredientSet map[string]bool

type FoodInfo struct {
	IngredientCount       map[string]int
	AllergenPossibilities map[string]IngredientSet
}

func ParseFoods(foodStrings []string) *FoodInfo {
	ingredientCount := make(map[string]int)
	allergenPossibilities := make(map[string]IngredientSet)

	for _, foodString := range foodStrings {
		matches := SplitFood.FindStringSubmatch(foodString)
		if matches == nil {
			log.Fatalf("Could not split food string '%v'", foodString)
		}

		ingredients := strings.Split(matches[1], " ")
		allergens := strings.Split(matches[2], ", ")

		for _, v := range ingredients {
			ingredientCount[v]++
		}

		for _, allergen := range allergens {
			allergenPossibilities[allergen] = intersection(allergenPossibilities[allergen], ingredients)
		}
	}

	return &FoodInfo{ingredientCount, allergenPossibilities}
}

func CountNonAllergens(foodStrings []string) int {
	foodInfo := ParseFoods(foodStrings)

	// Go through all ingredients that may have an allergen and delete from list
	ingredients := foodInfo.IngredientCount
	for _, ingredientSet := range foodInfo.AllergenPossibilities {
		for ingredient := range ingredientSet {
			delete(ingredients, ingredient)
		}
	}

	sum := 0
	for _, v := range ingredients {
		sum += v
	}

	return sum
}

func DangerousIngredients(foodStrings []string) string {
	foodInfo := ParseFoods(foodStrings)

	allergens := make(map[string]string)
	foundIngredients := make(IngredientSet)

	// Iterate through allergens finding ones with only on possible ingredient until
	// we've found all combinations
	for len(allergens) < len(foodInfo.AllergenPossibilities) {
		for allergen, ingredients := range foodInfo.AllergenPossibilities {
			if allergens[allergen] != "" {
				continue
			}
			left := singleLeft(ingredients, foundIngredients)

			if left != "" {
				allergens[allergen] = left
				foundIngredients[left] = true
			}
		}
	}

	// Alphabetically order by allergen
	allergenList := make([]string, 0)
	for a, _ := range allergens {
		allergenList = append(allergenList, a)
	}
	sort.Strings(allergenList)

	dangerous := make([]string, len(allergenList))
	for i, v := range allergenList {
		dangerous[i] = allergens[v]
	}
	return strings.Join(dangerous, ",")
}

func intersection(ingredientSet IngredientSet, newIngredients []string) IngredientSet {
	newSet := make(IngredientSet)

	for _, ingredient := range newIngredients {
		if ingredientSet == nil || ingredientSet[ingredient] {
			newSet[ingredient] = true
		}
	}

	return newSet
}

func singleLeft(ingredientSet IngredientSet, found IngredientSet) string {
	left := ""
	for ingredient := range ingredientSet {
		if !found[ingredient] {
			if left != "" {
				return ""
			}
			left = ingredient
		}
	}
	return left
}
