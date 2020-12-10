package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
	"sort"
	"strconv"
)

func main() {
	adapters := make([]int, 0)
	err := utils.OpenAndReadAll("input10.txt", func(s string) error {
		j, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		adapters = append(adapters, j)
		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One\n")
	sort.Ints(adapters)
	adapters = append([]int{0}, adapters...)
	adapters = append(adapters, adapters[len(adapters)-1]+3)
	diffs := make(map[int]int)
	for i := 0; i < len(adapters)-1; i++ {
		diff := adapters[i+1] - adapters[i]
		diffs[diff] += 1
	}
	result := diffs[1] * diffs[3]
	fmt.Printf("Result: %d\n", result)

	fmt.Printf("Part Two\n")
	arrTable := make([]int, len(adapters))
	arrTable[len(adapters)-1] = 1
	for startAdapter := len(adapters)-2; startAdapter >= 0; startAdapter-- {
		startJolts := adapters[startAdapter]
		nextAdapter := startAdapter + 1
		total := 0
		for nextAdapter < len(adapters) && 1 <= adapters[nextAdapter]-startJolts && adapters[nextAdapter]-startJolts <= 3 {
			total += arrTable[nextAdapter]
			nextAdapter++
		}
		arrTable[startAdapter] = total
	}
	fmt.Printf("Total: %d\n", arrTable[0])
}

