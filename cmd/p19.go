package main

import (
	"fmt"
	"github.com/ghjm/advent2020/pkg/utils"
	"regexp"
	"strings"
)

func ruleRegex(rules map[string][][]string, rule string) string {
	ruleData, ok := rules[rule]
	if !ok {
		panic(fmt.Sprintf("no rule named %s", rule))
	}
	cyclic := false
	for _, opt := range ruleData {
		for _, tok := range opt {
			if tok == rule {
				cyclic = true
			}
		}
	}
	if cyclic {
		if len(ruleData) == 2 {
			if len(ruleData[0]) == 1 && len(ruleData[1]) == 2 && ruleData[0][0] == ruleData[1][0] &&
				    ruleData[1][1] == rule {
				return fmt.Sprintf("(%s)+", ruleRegex(rules, ruleData[0][0]))
			} else if len(ruleData[0]) == 2 && len(ruleData[1]) == 3 && ruleData[0][0] == ruleData[1][0] &&
				ruleData[0][1] == ruleData[1][2] && ruleData[1][1] == rule {
				// this is an awful hack
				leftSide := ruleRegex(rules, ruleData[0][0])
				rightSide := ruleRegex(rules, ruleData[0][1])
				matches := make([]string, 0)
				for i := 1; i <= 10; i++ {
					leftRepeats := ""
					rightRepeats := ""
					for j := 0; j < i; j++ {
						leftRepeats += leftSide
						rightRepeats += rightSide
					}
					matches = append(matches, leftRepeats+rightRepeats)
				}
				return fmt.Sprintf("(%s)", strings.Join(matches,"|"))
			}
		}
		panic("cyclic")
	}
	ruleOptions := make([]string, 0)
	for _, opt := range ruleData {
		substr := ""
		for _, tok := range opt {
			if strings.HasPrefix(tok, `"`) {
				substr += tok[1:len(tok)-1]
			} else {
				substr += ruleRegex(rules, tok)
			}
		}
		ruleOptions = append(ruleOptions, substr)
	}
	if len(ruleOptions) == 0 {
		panic("no rule options")
	} else if len(ruleOptions) == 1 {
		return ruleOptions[0]
	} else {
		return fmt.Sprintf("(%s)", strings.Join(ruleOptions, "|"))
	}
}

func main() {
	rules := make(map[string][][]string, 0)
	messages := make([]string, 0)
	var state = 0
	err := utils.OpenAndReadAll("input19.txt", func(s string) error {
		if s == "" {
			state++
		} else if state == 0 {
			rv := strings.Split(s, ": ")
			if len(rv) != 2 {
				return fmt.Errorf("bad rule: %s", s)
			}
			rule := make([][]string, 0)
			for _, optStr := range strings.Split(rv[1], "|") {
				optStr = strings.TrimSpace(optStr)
				ruleArr := make([]string, 0)
				for _, ruleStr := range strings.Split(optStr, " ") {
					ruleStr = strings.TrimSpace(ruleStr)
					ruleArr = append(ruleArr, ruleStr)
				}
				rule = append(rule, ruleArr)
			}
			rules[rv[0]] = rule
		} else if state == 1 {
			messages = append(messages, s)
		}
		return nil
	})
	if err != nil {
		panic(err)
	}
	fmt.Printf("Part One\n")
	re := regexp.MustCompile(fmt.Sprintf("^%s$", ruleRegex(rules, "0")))
	matchCount := 0
	for _, msg := range messages {
		if re.MatchString(msg) {
			matchCount++
		}
	}
	fmt.Printf("Matches: %d\n", matchCount)
	fmt.Printf("Part Two\n")
	rules["8"] = [][]string{{"42"}, {"42", "8"}}
	rules["11"] = [][]string{{"42", "31"}, {"42", "11", "31"}}
	re = regexp.MustCompile(fmt.Sprintf("^%s$", ruleRegex(rules, "0")))
	matchCount = 0
	for _, msg := range messages {
		if re.MatchString(msg) {
			matchCount++
		}
	}
	fmt.Printf("Matches: %d\n", matchCount)
}

