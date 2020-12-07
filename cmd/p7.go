package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
	"regexp"
	"strconv"
)

func main() {
	mainRE := regexp.MustCompile(`^(.*) bags contain (.*)\.$`)
	subRE := regexp.MustCompile(`(\d+) ([a-z ]*) bag`)
	bags := make(map[string]map[string]int)
	err := utils.OpenAndReadAll("input7.txt", func(s string) error {
		mainMatch := mainRE.FindStringSubmatch(s)
		if mainMatch == nil || len(mainMatch) != 3 {
			return fmt.Errorf("Line failed to match: %s\n", s)
		}
		mainColor := mainMatch[1]
		_, ok := bags[mainColor]
		if !ok {
			bags[mainColor] = make(map[string]int)
		}
		subMatches := subRE.FindAllStringSubmatch(mainMatch[2], -1)
		for _, subBag := range subMatches {
			if len(subBag) != 3 {
				return fmt.Errorf("Line failed to match: %s\n", s)
			}
			qty, _ := strconv.Atoi(subBag[1])
			subColor := subBag[2]
			bags[mainColor][subColor] = qty
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part One\n")
	seen := make(map[string]struct{})
	open := []string{"shiny gold"}
	solutions := make(map[string]struct{})
	for len(open) > 0 {
		target := open[0]
		open = open[1:]
		_, ok := seen[target]
		if !ok {
			seen[target] = struct{}{}
			for mainColor, subMap := range bags {
				for subColor := range subMap {
					_, ok := solutions[subColor]
					if ok || subColor == target {
						open = append(open, mainColor)
						solutions[mainColor] = struct{}{}
						break
					}
				}
			}
		}
	}
	fmt.Printf("Colors: %d\n", len(solutions))
	fmt.Printf("Part Two\n")
	needed := map[string]int{"shiny gold": 1}
	total := make(map[string]int)
	for len(needed) > 0 {
		ok := false
		var neededItem string
		var neededQty int
		for k, v := range needed {
			neededItem = k
			neededQty = v
			ok = true
			delete(needed, k)
			break
		}
		if !ok {
			panic("failed to get item from needed list")
		}
		subMap := bags[neededItem]
		for subColor, subQty := range subMap {
			subQty *= neededQty
			needed[subColor] += subQty
			total[subColor] += subQty
		}
	}
	sum := 0
	for _, v := range total {
		sum += v
	}
	fmt.Printf("Bags: %d\n", sum)
}

