package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
)

type oper = struct {
	prec   int
	rAssoc bool
}

// Adapted from https://rosettacode.org/wiki/Parsing/Shunting-yard_algorithm#Go
func parseInfix(e string, opa map[string]oper) (rpn string) {
	var stack []string // holds operators and left parenthesis
	for i := range e {
		tok := e[i:i+1]
		switch tok {
		case " ":
			continue
		case "(":
			stack = append(stack, tok) // push "(" to stack
		case ")":
			var op string
			for {
				// pop item ("(" or operator) from stack
				if stack == nil || len(stack) == 0 {
					panic("mismatched parentheses")
				}
				op, stack = stack[len(stack)-1], stack[:len(stack)-1]
				if op == "(" {
					break // discard "("
				}
				rpn += " " + op // add operator to result
			}
		default:
			if o1, isOp := opa[tok]; isOp {
				// token is an operator
				for len(stack) > 0 {
					// consider top item on stack
					op := stack[len(stack)-1]
					if o2, isOp := opa[op]; !isOp || o1.prec > o2.prec ||
						o1.prec == o2.prec && o1.rAssoc {
						break
					}
					// top item is an operator that needs to come off
					stack = stack[:len(stack)-1] // pop it
					rpn += " " + op              // add it to result
				}
				// push operator (the new one) to stack
				stack = append(stack, tok)
			} else { // token is an operand
				if rpn > "" {
					rpn += " "
				}
				rpn += tok // add operand to result
			}
		}
	}
	// drain stack to result
	for len(stack) > 0 {
		rpn += " " + stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	return
}

func evaluateRPN(r string) int {
	stack := make([]int, 0)
	for i := range r {
		tok := r[i : i+1]
		switch tok {
		case " ":
			continue
		case "+":
			var v1, v2 int
			v1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			v2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, v1+v2)
		case "*":
			var v1, v2 int
			v1, stack = stack[len(stack)-1], stack[:len(stack)-1]
			v2, stack = stack[len(stack)-1], stack[:len(stack)-1]
			stack = append(stack, v1*v2)
		default:
			v := utils.AtoiOrPanic(tok)
			stack = append(stack, v)
		}
	}
	if len(stack) != 1 {
		panic("evaluation failed")
	}
	return stack[0]
}

func main() {
	expressions := make([]string, 0)
	err := utils.OpenAndReadAll("input18.txt", func(s string) error {
		expressions = append(expressions, s)
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part One\n")
	sum := 0
	for i := range expressions {
		rpn := parseInfix(expressions[i], map[string]oper{
			"*": {2, false},
			"+": {2, false},
		})
		v := evaluateRPN(rpn)
		sum += v
	}
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Part Two\n")
	sum = 0
	for i := range expressions {
		rpn := parseInfix(expressions[i], map[string]oper{
			"*": {2, false},
			"+": {3, false},
		})
		v := evaluateRPN(rpn)
		sum += v
	}
	fmt.Printf("Sum: %d\n", sum)
}

