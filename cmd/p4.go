package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
	"strconv"
	"strings"
	"unicode"
)

func checkYear(year string, min int, max int) bool {
	if !checkNumber(year, 4) {
		return false
	}
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		panic(err)
	}
	return yearInt >= min && yearInt <= max
}

func checkHeight(height string) bool {
	heightNumPart := height[0:len(height)-2]
	heightInt, err := strconv.Atoi(heightNumPart)
	if err != nil {
		return false
	}
	if strings.HasSuffix(height, "cm") {
		return heightInt >= 150 && heightInt <= 193
	} else if strings.HasSuffix(height, "in") {
		return heightInt >= 59 && heightInt <= 76
	} else {
		return false
	}
}

func checkHairColor(color string) bool {
	if len(color) != 7 {
		return false
	}
	if color[0] != '#' {
		return false
	}
	for i := 1; i < len(color); i++ {
		if !unicode.Is(unicode.Hex_Digit, rune(color[i])) {
			return false
		}
	}
	return true
}

var validEyeColors = map[string]struct{} {
	"amb": {},
	"blu": {},
	"brn": {},
	"gry": {},
	"grn": {},
	"hzl": {},
	"oth": {},
}

func checkEyeColor(color string) bool {
	_, ok := validEyeColors[color]
	return ok
}

func checkNumber(number string, length int) bool {
	if len(number) != length {
		return false
	}
	for _, c := range number {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

func main() {
	requiredFields := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	passport := make(map[string]string)
	validCount := 0
	fullyValidCount := 0
	checkPassport := func() {
		valid := true
		for _, rf := range requiredFields {
			_, ok := passport[rf]
			if !ok {
				valid = false
				break
			}
		}
		if valid {
			validCount++
			fullyValid :=
				checkYear(passport["byr"], 1920, 2002) &&
				checkYear(passport["iyr"], 2010, 2020) &&
				checkYear(passport["eyr"], 2020, 2030) &&
				checkHeight(passport["hgt"]) &&
				checkHairColor(passport["hcl"]) &&
				checkEyeColor(passport["ecl"]) &&
				checkNumber(passport["pid"], 9)
			if fullyValid {
				fullyValidCount++
			}
		}
		passport = make(map[string]string)
	}
	err := utils.OpenAndReadAll("input4.txt", func(s string) error {
		if s == "" {
			checkPassport()
		} else {
			values := strings.Split(s, " ")
			for _, v := range values {
				comps := strings.Split(v, ":")
				if len(comps) != 2 {
					panic("Passport format error")
				}
				passport[comps[0]] = comps[1]
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	checkPassport()
	fmt.Printf("Part One\n")
	fmt.Printf("Valid passports: %d\n", validCount)
	fmt.Printf("Part Two\n")
	fmt.Printf("Fully valid passports: %d\n", fullyValidCount)
}

