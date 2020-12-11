package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
)

func countAdjDirect(seatMap [][]rune, i int, j int) int {
	adjacent := 0
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			ci := i + di
			cj := j + dj
			if ci < 0 || cj < 0 || ci >= len(seatMap) || cj >= len(seatMap[i]) {
				continue
			}
			if seatMap[ci][cj] == '#' {
				adjacent++
			}
		}
	}
	return adjacent
}

func countAdjLineOfSight(seatMap [][]rune, i int, j int) int {
	adjacent := 0
	for di := -1; di <= 1; di++ {
		for dj := -1; dj <= 1; dj++ {
			if di == 0 && dj == 0 {
				continue
			}
			ci := i
			cj := j
			for {
				ci = ci + di
				cj = cj + dj
				if ci < 0 || cj < 0 || ci >= len(seatMap) || cj >= len(seatMap[i]) {
					break
				}
				if seatMap[ci][cj] == '#' {
					adjacent++
					break
				} else if seatMap[ci][cj] == 'L' {
					break
				}
			}
		}
	}
	return adjacent
}

func simulate(seatMap [][]rune, tolerance int, adjFunc func([][]rune, int, int) int) [][]rune {
	oldMap := seatMap
	var newMap [][]rune
	for {
		newMap = make([][]rune, len(oldMap))
		for i := range oldMap {
			row := make([]rune, len(oldMap[i]))
			for j := range oldMap[i] {
				row[j] = oldMap[i][j]
			}
			newMap[i] = row
		}
		changed := false
		for i := range oldMap {
			for j := range oldMap[i] {
				if oldMap[i][j] == '.' {
					continue
				}
				adjacent := adjFunc(oldMap, i, j)
				if oldMap[i][j] == 'L' && adjacent == 0 {
					newMap[i][j] = '#'
					changed = true
				} else if oldMap[i][j] == '#' && adjacent >= tolerance {
					newMap[i][j] = 'L'
					changed = true
				}
			}
		}
		if !changed {
			break
		}
		oldMap = newMap
	}
	return newMap
}

func countOccupied(seatMap [][]rune) int {
	occupied := 0
	for i := range seatMap {
		for j := range seatMap[i] {
			if seatMap[i][j] == '#' {
				occupied++
			}
		}
	}
	return occupied
}

func main() {
	seatMap := make([][]rune, 0)
	err := utils.OpenAndReadAll("input11.txt", func(s string) error {
		row := make([]rune, len(s))
		for i := range s {
			row[i] = rune(s[i])
		}
		seatMap = append(seatMap, row)
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part One\n")
	occupied := countOccupied(simulate(seatMap, 4, countAdjDirect))
	fmt.Printf("Occupied: %d\n", occupied)
	fmt.Printf("Part Two\n")
	occupied = countOccupied(simulate(seatMap, 5, countAdjLineOfSight))
	fmt.Printf("Occupied: %d\n", occupied)
}

