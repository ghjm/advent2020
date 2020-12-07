package main

import (
	"fmt"
	mapset "github.com/deckarep/golang-set"
	"github.com/ghjm/advent2020/pkg/utils"
)

func alphabetSet() mapset.Set {
	result := mapset.NewSet()
	for c := 'a'; c <= 'z'; c++ {
		result.Add(c)
	}
	return result
}

func main() {
	accumAny := mapset.NewSet()
	accumAll := alphabetSet()
	groupsAny := make([]mapset.Set, 0)
	groupsAll := make([]mapset.Set, 0)
	err := utils.OpenAndReadAll("input6.txt", func(s string) error {
		if s == "" {
				groupsAny = append(groupsAny, accumAny)
				groupsAll = append(groupsAll, accumAll)
				accumAny = mapset.NewSet()
				accumAll = alphabetSet()
		} else {
			thisSet := mapset.NewSet()
			for _, c := range s {
				thisSet.Add(c)
			}
			accumAny = accumAny.Union(thisSet)
			accumAll = accumAll.Intersect(thisSet)
		}
		return nil
	})
	groupsAny = append(groupsAny, accumAny)
	groupsAll = append(groupsAll, accumAll)
	if err != nil {
		panic(err)
	}
	sumAny := 0
	for _, g := range groupsAny {
		sumAny += g.Cardinality()
	}
	sumAll := 0
	for _, g := range groupsAll {
		sumAll += g.Cardinality()
	}
	fmt.Printf("Part One\n")
	fmt.Printf("Sum: %d\n", sumAny)
	fmt.Printf("Part Two\n")
	fmt.Printf("Sum: %d\n", sumAll)
}

