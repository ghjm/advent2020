package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/ghjm/advent2020/pkg/utils"
	"regexp"
	"sort"
	"strings"
)

type foodT = struct {
	ingredients []string
	allergens []string
}

func main() {
	foods := make([]foodT, 0)
	re := regexp.MustCompile(`^([a-z ]+)\(contains ([a-z ,]+)\)$`)
	err := utils.OpenAndReadAll("input21.txt", func(s string) error {
		m := re.FindStringSubmatch(s)
		if m == nil {
			panic("regex didn't match")
		}
		food := foodT{
			ingredients: strings.Split(strings.TrimSpace(m[1]), " "),
			allergens:   strings.Split(strings.TrimSpace(m[2]), ", "),
		}
		foods = append(foods, food)
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part One\n")
	allergenSets := make(map[string][]mapset.Set)
	for _, food := range foods {
		ingSet := mapset.NewSet()
		for _, ing := range food.ingredients {
			ingSet.Add(ing)
		}
		for _, allergen := range food.allergens {
			_, ok := allergenSets[allergen]
			if !ok {
				allergenSets[allergen] = make([]mapset.Set, 0)
			}
			allergenSets[allergen] = append(allergenSets[allergen], ingSet)
		}
	}
	possibleSets := make(map[string]mapset.Set)
	for aName, aSets := range allergenSets {
		var interSet mapset.Set
		for _, aSet := range aSets {
			if interSet == nil {
				interSet = mapset.NewSet()
				interSet = interSet.Union(aSet)
			} else {
				interSet = interSet.Intersect(aSet)
			}
		}
		possibleSets[aName] = interSet
	}
	everyPossible := mapset.NewSet()
	for _, pSet := range possibleSets {
		everyPossible = everyPossible.Union(pSet)
	}
	sum := 0
	for _, food := range foods {
		for _, ing := range food.ingredients {
			if !everyPossible.Contains(ing) {
				sum++
			}
		}
	}
	fmt.Printf("Sum: %d\n", sum)

	fmt.Printf("Part Two\n")
	ingredientAllergens := make(map[string]string)
	foundAny := true
	for foundAny {
		foundAny = false
		for aName, pSet := range possibleSets {
			if pSet.Cardinality() == 1 {
				foundAny = true
				ing := pSet.Pop().(string)
				ingredientAllergens[aName] = ing
				delete(possibleSets, aName)
				for _, ppSet := range possibleSets {
					ppSet.Remove(ing)
				}
			}
		}
	}
	keys := make([]string, 0)
	for allergen := range ingredientAllergens {
		keys = append(keys, allergen)
	}
	sort.Strings(keys)
	ingredients := make([]string, 0)
	for _, key := range keys {
		ingredients = append(ingredients, ingredientAllergens[key])
	}
	fmt.Printf("Canonical List: %s", strings.Join(ingredients, ","))
}

