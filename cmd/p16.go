package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
	"regexp"
	"strings"
)

func main() {
	state := 0
	myTicket := make([]int, 0)
	nearbyTickets := make([][]int, 0)
	fields := make(map[string]map[int]struct{})
	fieldsRe := regexp.MustCompile(`^([a-z ]+): (\d+)-(\d+) or (\d+)-(\d+)$`)
	err := utils.OpenAndReadAll("input16.txt", func(s string) error {
		if s == "" {
			return nil
		} else if s == "your ticket:" {
			state = 1
		} else if s == "nearby tickets:" {
			state = 2
		} else if state == 0 {
			m := fieldsRe.FindStringSubmatch(s)
			if m == nil {
				panic("Regex failed to match")
			}
			fieldHash := make(map[int]struct{})
			for i := utils.AtoiOrPanic(m[2]); i <= utils.AtoiOrPanic(m[3]); i++ {
				fieldHash[i] = struct{}{}
			}
			for i := utils.AtoiOrPanic(m[4]); i <= utils.AtoiOrPanic(m[5]); i++ {
				fieldHash[i] = struct{}{}
			}
			fields[m[1]] = fieldHash
		} else if state == 1 {
			if len(myTicket) > 0 {
				panic("More than one my ticket")
			}
			for _, v := range strings.Split(s, ",") {
				myTicket = append(myTicket, utils.AtoiOrPanic(v))
			}
		} else if state == 2 {
			curTicket := make([]int, 0)
			for _, v := range strings.Split(s, ",") {
				curTicket = append(curTicket, utils.AtoiOrPanic(v))
			}
			nearbyTickets = append(nearbyTickets, curTicket)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part One\n")
	tser := 0
	validTickets := make([][]int, 0)
	for _, ticket := range nearbyTickets {
		ticketValid := true
		for _, value := range ticket {
			valid := false
			for _, field := range fields {
				_, ok := field[value]
				if ok {
					valid = true
					break
				}
			}
			if !valid {
				ticketValid = false
				tser += value
			}
		}
		if ticketValid {
			validTickets = append(validTickets, ticket)
		}
	}
	fmt.Printf("Error Rate: %d\n", tser)
	fmt.Printf("Part Two\n")
	fieldPositions := make([]map[string]struct{}, len(myTicket))
	for i := range myTicket {
		fMap := make(map[string]struct{})
		for f := range fields {
			fMap[f] = struct{}{}
		}
		fieldPositions[i] = fMap
	}
	for _, ticket := range validTickets {
		for pos, value := range ticket {
			for fieldName := range fieldPositions[pos] {
				_, ok := fields[fieldName][value]
				if !ok {
					delete(fieldPositions[pos], fieldName)
				}
			}
		}
	}
	done := false
	for !done {
		done = true
		for pos := range myTicket {
			if len(fieldPositions[pos]) == 1 {
				var fieldName string
				for k := range fieldPositions[pos] {
					fieldName = k
				}
				for delpos := range myTicket {
					if delpos != pos {
						_, ok := fieldPositions[delpos][fieldName]
						if ok {
							delete(fieldPositions[delpos], fieldName)
							done = false
						}
					}
				}
			}
		}
	}
	product := 1
	for pos := range myTicket {
		if len(fieldPositions[pos]) != 1 {
			panic("Satisfaction error")
		}
		var fieldName string
		for k := range fieldPositions[pos] {
			fieldName = k
		}
		if strings.HasPrefix(fieldName, "departure") {
			product *= myTicket[pos]
		}
	}
	fmt.Printf("Product: %d\n", product)
}

