package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
	"math"
	"strconv"
)

func check(v int, list []int) bool {
	if v == 143 {
		fmt.Printf("here we go\n")
	}
	for i := 0; i < len(list); i++ {
		for j := i+1; j < len(list); j++ {
			if list[i]+list[j] == v {
				return true
			}
		}
	}
	return false
}

func main() {
	numbers := make([]int, 0)
	err := utils.OpenAndReadAll("input9.txt", func(s string) error {
		v, err := strconv.Atoi(s)
		if err != nil {
			return err
		}
		numbers = append(numbers, v)
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part One\n")
	var failure int
	for i := 25; i < len(numbers); i++ {
		if !check(numbers[i], numbers[i-25:i]) {
			failure = numbers[i]
			fmt.Printf("First failure: %d\n", failure)
			break
		}
	}
	fmt.Printf("Part Two\n")
	for i := range numbers {
		sum := 0
		j := 0
		found := false
		min := math.MaxInt32
		max := -math.MaxInt32
		for {
			sum += numbers[i+j]
			if sum == failure {
				for k := i; k <= i+j; k++ {
					if numbers[k] > max {
						max = numbers[k]
					}
					if numbers[k] < min {
						min = numbers[k]
					}
				}
				found = true
				break
			} else if sum > failure {
				break
			}
			j++
		}
		if found {
			fmt.Printf("Weakness: %d\n", min+max)
			break
		}
	}
}

