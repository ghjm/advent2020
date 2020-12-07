package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
)

func main() {
	err := utils.OpenAndReadAll("input.txt", func(s string) error {
		fmt.Printf("%s\n", s)
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part One\n")
	fmt.Printf("Part Two\n")
}

