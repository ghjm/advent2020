package main

import (
    "fmt"
    "github.com/ghjm/advent2020/pkg/utils"
)

func main() {
    var (
        min int
        max int
        char uint8
        password string
        charCount int
        validPartOne int
        validPartTwo int
    )
    err := utils.OpenAndReadAll("input2.txt", func(s string) error {
        _, err := fmt.Sscanf(s, "%d-%d %c: %s", &min, &max, &char, &password)
        if err != nil {
            return err
        }

        // Part One
        charCount = 0
        for _, c := range password {
            if c == int32(char) {
                charCount++
            }
        }
        if charCount >= min && charCount <= max {
            validPartOne++
        }

        // Part Two
        if len(password) >= max && (password[min-1] == char) != (password[max-1] == char) {
            validPartTwo++
        }

        return nil
    })
    if err != nil {
        panic(err)
    }

    fmt.Printf("Part One\n")
    fmt.Printf("Valid count: %d\n", validPartOne)
    fmt.Printf("Part Two\n")
    fmt.Printf("Valid count: %d\n", validPartTwo)
}
