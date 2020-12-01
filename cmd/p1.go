package main

import (
    "fmt"
    "github.com/ghjm/advent2020/pkg/utils"
    "strconv"
    "strings"
)

func search(data map[int]struct{}, target int, num int) []int {
    if num <= 0 {
        return nil
    } else if num == 1 {
        _, ok := data[target]
        if ok {
            return []int{target}
        } else {
            return nil
        }
    } else {
        for v := range data {
            results := search(data, target-v, num-1)
            if results != nil {
                results = append([]int{v}, results...)
                return results
            }
        }
    }
    return nil
}

func printResults(results []int) {
    if results == nil || len(results) == 0 {
        fmt.Printf("Not found!\n")
    } else {
        resultStrings := make([]string, 0)
        resultProduct := 1
        for _, v := range results {
            resultStrings = append(resultStrings, strconv.Itoa(v))
            resultProduct *= v
        }
        fmt.Printf("Items: %s. Product: %d.\n", strings.Join(resultStrings, ", "), resultProduct)
    }
}

func main() {
    data := make(map[int]struct{}, 0)
    err := utils.OpenAndReadAll("input1.txt", func(s string) error {
        value, err := strconv.Atoi(s)
        if err != nil {
            return err
        }
        data[value] = struct{}{}
        return nil
    })
    if err != nil {
        panic(err)
    }

    fmt.Printf("Part One\n")
    printResults(search(data, 2020, 2))

    fmt.Printf("Part Two\n")
    printResults(search(data, 2020, 3))
}
