package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
	"regexp"
	"strconv"
)

type instruction struct {
	instr string
	param int
}

func runcode(code []instruction) (ip int, accum int, terminated bool) {
	accum = 0
	ip = 0
	visited := make(map[int]struct{})
	for {
		_, ok := visited[ip]
		if ok {
			return ip, accum, false
		}
		visited[ip] = struct{}{}
		instr := code[ip]
		if instr.instr == "acc" {
			accum += instr.param
			ip += 1
		} else if instr.instr == "jmp" {
			ip += instr.param
		} else if instr.instr == "nop" {
			ip += 1
		}
		if ip >= len(code) {
			return ip, accum, true
		}
	}
}

func main() {
	re := regexp.MustCompile(`^(acc|jmp|nop) ([+-]\d+)$`)
	code := make([]instruction, 0)
	err := utils.OpenAndReadAll("input8.txt", func(s string) error {
		m := re.FindStringSubmatch(s)
		if m == nil || len(m) != 3 {
			return fmt.Errorf("Regex failed to match on line %s\n", s)
		}
		paramInt, _ := strconv.Atoi(m[2])
		code = append(code, instruction{
			instr: m[1],
			param: paramInt,
		})
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part One\n")
	_, accum, _ := runcode(code)
	fmt.Printf("Accum: %d\n", accum)
	fmt.Printf("Part Two\n")
	for i := range code {
		if code[i].instr != "acc" {
			newcode := make([]instruction, len(code))
			for j := range code {
				if i == j {
					if code[j].instr == "jmp" {
						newcode[j].instr = "nop"
					} else {
						newcode[j].instr = "jmp"
					}
					newcode[j].param = code[i].param
				} else {
					newcode[j] = code[j]
				}
			}
			_, accum, term := runcode(newcode)
			if term {
				fmt.Printf("Accum: %d\n", accum)
				break
			}
		}
	}
}

