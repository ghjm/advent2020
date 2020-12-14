package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	instrs := make([]string, 0)
	err := utils.OpenAndReadAll("input14.txt", func(s string) error {
		instrs = append(instrs, s)
		return nil
	})
	if err != nil {
		panic(err)
	}
	maskRe := regexp.MustCompile(`^mask = ([01X]+)$`)
	memRe := regexp.MustCompile(`^mem\[(\d+)\] = (\d+)$`)
	fmt.Printf("Part One\n")
	var andMask int64
	var orMask int64
	mem := make(map[int64]int64)
	for _, s := range instrs {
		m := maskRe.FindStringSubmatch(s)
		if m != nil {
			mask := m[1]
			andString := strings.ReplaceAll(mask, "X", "1")
			andMask, err = strconv.ParseInt(andString, 2, 64)
			if err != nil {
				panic(err)
			}
			orString := strings.ReplaceAll(mask, "X", "0")
			orMask, err = strconv.ParseInt(orString, 2, 64)
			if err != nil {
				panic(err)
			}
			continue
		}
		m = memRe.FindStringSubmatch(s)
		if m != nil {
			address := utils.AtoiOrPanic64(m[1])
			value := utils.AtoiOrPanic64(m[2])
			value &= andMask
			value |= orMask
			mem[address] = value
			continue
		}
		panic(fmt.Sprintf("Line %s did not match", s))
	}
	var sum int64
	for _, v := range mem {
		sum += v
	}
	fmt.Printf("Sum: %d\n", sum)

	fmt.Printf("Part Two\n")
	orMask = 0
	var floatMask []int64
	mem = make(map[int64]int64)
	for _, s := range instrs {
		m := maskRe.FindStringSubmatch(s)
		if m != nil {
			mask := m[1]
			orString := strings.ReplaceAll(mask, "X", "0")
			orMask, err = strconv.ParseInt(orString, 2, 64)
			bitVal := int64(1)
			floatMask = make([]int64, 0)
			for p := len(mask)-1; p >= 0; p-- {
				if mask[p] == 'X' {
					floatMask = append(floatMask, bitVal)
				}
				bitVal *= 2
			}
			if err != nil {
				panic(err)
			}
			continue
		}
		m = memRe.FindStringSubmatch(s)
		if m != nil {
			address := utils.AtoiOrPanic64(m[1])
			value := utils.AtoiOrPanic64(m[2])
			address |= orMask
			if len(floatMask) == 0 {
				mem[address] = value
			} else {
				var maskSum int64
				for i := range floatMask {
					maskSum += floatMask[i]
				}
				for i := int64(0); i < 1<<len(floatMask); i++ {
					curAddress := address &^ maskSum
					bits := strconv.FormatInt(i, 2)
					for len(bits) < len(floatMask) {
						bits = "0" + bits
					}
					for j := 0; j < len(bits); j++ {
						if bits[j] == '1' {
							curAddress |= floatMask[j]
						}
					}
					mem[curAddress] = value
				}
			}
			continue
		}
		panic(fmt.Sprintf("Line %s did not match", s))
	}
	sum = 0
	for _, v := range mem {
		sum += v
	}
	fmt.Printf("Sum: %d\n", sum)

}

