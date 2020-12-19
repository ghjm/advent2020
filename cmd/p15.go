package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
	"strings"
)

func playGame(startingNums []int, limit int) int {
	history1 := make(map[int]int)
	history2 := make(map[int]int)
	var lastSpoken int
	for i := 0; i < limit; i++ {
		var curSpoken int
		if i < len(startingNums) {
			curSpoken = startingNums[i]
		} else {
			ls1, ok1 := history1[lastSpoken]
			ls2, ok2 := history2[lastSpoken]
			if !ok1 || !ok2 {
				curSpoken = 0
			} else {
				curSpoken = ls1 - ls2
			}
		}
		fmt.Printf("%d\n", curSpoken)
		cs1, ok1 := history1[curSpoken]
		if ok1 {
			history2[curSpoken] = cs1
		}
		history1[curSpoken] = i
		lastSpoken = curSpoken
	}
	return lastSpoken
}

func main() {
	startingNums := make([]int, 0)
	err := utils.OpenAndReadAll("input15.txt", func(s string) error {
		if len(startingNums) > 0 {
			panic("Too much input")
		}
		nums := strings.Split(s,",")
		for _, num := range nums {
			startingNums = append(startingNums, utils.AtoiOrPanic(num))
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part One\n")
	fmt.Printf("Number spoken: %d\n", playGame(startingNums, 2020))
	fmt.Printf("Part Two\n")
	fmt.Printf("Number spoken: %d\n", playGame(startingNums, 30000000))
}

