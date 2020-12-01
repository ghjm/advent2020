package main

import (
    "fmt"
    "github.com/ghjm/advent2020/pkg/utils"
    "strconv"
    "strings"
)

func search(data map[int]struct{}, visited map[int]struct{}, target int, num int) []int {
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
        subvisited := make(map[int]struct{})
        for v := range visited {
            subvisited[v] = struct{}{}
        }
        for v := range data {
            _, ok := visited[v]
            if ok {
                continue
            }
            subvisited[v] = struct{}{}
            results := search(data, subvisited, target-v, num-1)
            if results != nil {
                results = append([]int{v}, results...)
                return results
            }
        }
    }
    return nil
}

func slowsearch(data []int, target int, num int) []int {
    if num <= 0 {
        return nil
    }
    for i, v := range data {
        if num == 1 {
            if v == target {
                return []int{v}
            }
        } else {
            results := slowsearch(data[i+1:], target-v, num-1)
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
    dataMap := make(map[int]struct{})
    dataList := make([]int, 0)
    useSlowSearch := false
    err := utils.OpenAndReadAll("input1.txt", func(s string) error {
        value, err := strconv.Atoi(s)
        if err != nil {
            return err
        }
        dataList = append(dataList, value)
        _, ok := dataMap[value]
        if ok {
            fmt.Printf("Found duplicate value. Using slow search.\n")
            useSlowSearch = true
        }
        if !useSlowSearch {
            dataMap[value] = struct{}{}
        }
        return nil
    })
    if err != nil {
        panic(err)
    }

    fmt.Printf("Part One\n")
    var results []int
    if useSlowSearch {
        results = slowsearch(dataList, 2020, 2)
    } else {
        results = search(dataMap, nil, 2020, 2)
    }
    printResults(results)

    fmt.Printf("Part Two\n")
    if useSlowSearch {
        results = slowsearch(dataList, 2020, 3)
    } else {
        results = search(dataMap, nil, 2020, 3)
    }
    printResults(results)
}
