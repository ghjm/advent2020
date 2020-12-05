package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
	"strconv"
	"strings"
)

func main() {
	var highestID int64
	knownSeatIDs := make(map[int64]struct{})
	err := utils.OpenAndReadAll("input5.txt", func(s string) error {
		if len(s) != 10 {
			panic("Wrong length")
		}
		row := s[:7]
		col := s[7:]
		row = strings.Replace(row, "F", "0", -1)
		row = strings.Replace(row, "B", "1", -1)
		rowInt, err := strconv.ParseInt(row, 2, 32)
		if err != nil {
			panic(err)
		}
		col = strings.Replace(col, "L", "0", -1)
		col = strings.Replace(col, "R", "1", -1)
		colInt, err := strconv.ParseInt(col, 2, 32)
		if err != nil {
			panic(err)
		}
		seatID := rowInt*8 + colInt
		knownSeatIDs[seatID] = struct{}{}
		if seatID > highestID {
			highestID = seatID
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part One\n")
	fmt.Printf("Highest seat ID: %d\n", highestID)
	fmt.Printf("Part Two\n")
	for i := int64(0); i < 128*8; i++ {
		_, ok := knownSeatIDs[i]
		if ok {
			continue
		}
		_, ok1 := knownSeatIDs[i-8]
		_, ok2 := knownSeatIDs[i+8]
		if ok1 && ok2 {
			fmt.Printf("My seat ID: %d\n", i)
			break
		}
	}
}

