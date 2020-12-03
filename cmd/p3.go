package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
)

func countTrees(forest [][]int32, right int, down int) int {
	row := 0
	col := 0
	trees := 0
	width := len(forest[0])
	for {
		col += right
		row += down
		if row >= len(forest) {
			break
		}
		colmod := col % width
		if forest[row][colmod] == '#' {
			trees++
		}
	}
	return trees
}

func main() {
	forest := make([][]int32, 0)
	width := 0
	err := utils.OpenAndReadAll("input3.txt", func(s string) error {
		if width == 0 {
			width = len(s)
		} else {
			if len(s) != width {
				panic("Map width not consistent")
			}
		}
		forest = append(forest, []int32(s))
		return nil
	})
	if err != nil {
		panic(err)
	}

	fmt.Printf("Part One\n")
	fmt.Printf("Trees: %d\n", countTrees(forest, 3, 1))

	fmt.Printf("Part Two\n")
	product := countTrees(forest, 1, 1) *
		countTrees(forest, 3, 1) *
		countTrees(forest, 5, 1) *
		countTrees(forest, 7, 1) *
		countTrees(forest, 1, 2)
	fmt.Printf("Trees: %d\n", product)

}

