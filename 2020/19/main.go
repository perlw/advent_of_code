package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Rule ...
type Rule struct {
	R rune
	M [][]int
}

var ruleCache = make(map[int][]string)

func resetCache() {
	ruleCache = make(map[int][]string)
}

func generateValidMessages(rules map[int]Rule, index int, debug bool) []string {
	if debug {
		fmt.Printf("----\n")
	}
	if c, ok := ruleCache[index]; ok {
		return c
	}
	result := make([]string, 0, 10)
	for _, m := range rules[index].M {
		current := make([]string, 0, 10)

		for _, i := range m {
			if rules[i].R > 0 {
				alt := string(rules[i].R)
				if debug {
					fmt.Printf("adding %s to %+v\n", alt, current)
				}
				if len(current) == 0 {
					current = append(current, alt)
				} else {
					for i := 0; i < len(current); i++ {
						current[i] += alt
					}
				}
				if debug {
					fmt.Printf("now %+v\n", current)
				}
			} else {
				tmp := make([]string, 0, 10)
				parts := generateValidMessages(rules, i, debug)
				if debug {
					fmt.Printf("current %+v parts %+v\n", current, parts)
				}
				for _, p := range parts {
					if len(current) == 0 {
						tmp = append(tmp, p)
					} else {
						for _, c := range current {
							tmp = append(tmp, c+p)
						}
					}
				}
				current = tmp

				if debug {
					fmt.Printf("new current %+v\n", current)
				}
			}
		}

		for _, c := range current {
			result = append(result, c)
		}
		ruleCache[index] = result
	}
	ruleCache[index] = result
	return result
}

// Task1 ...
func Task1(rules map[int]Rule, messages []string, debug bool) int {
	valid := generateValidMessages(rules, 0, debug)
	fmt.Printf("Valid rules: %d\n", len(valid))

	var result int
	for _, m := range messages {
		var found bool
		for _, v := range valid {
			if m == v {
				found = true
				fmt.Printf("%s found!\n", m)
				result++
				break
			}
		}
		if !found {
			fmt.Printf("%s missing...\n", m)
		}
	}

	return result
}

// Task2 ...
func Task2(rules map[int]Rule, messages []string, debug bool) int {
	rules[8] = Rule{
		M: [][]int{
			{42},
			{42, 8},
		},
	}
	rules[11] = Rule{
		M: [][]int{
			{42, 31},
			{42, 11, 31},
		},
	}

	valid := generateValidMessages(rules, 0, debug)
	fmt.Printf("Valid rules: %d\n", len(valid))

	var result int
	for _, m := range messages {
		var found bool
		for _, v := range valid {
			if m == v {
				found = true
				fmt.Printf("%s found!\n", m)
				result++
				break
			}
		}
		if !found {
			fmt.Printf("%s missing...\n", m)
		}
	}

	return result
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	rules := make(map[int]Rule)
	messages := make([]string, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, ": ")
		i, _ := strconv.Atoi(parts[0])
		matchParts := strings.Split(parts[1], " | ")

		rule := Rule{
			M: make([][]int, 0, 10),
		}
		if matchParts[0][0] == '"' {
			rule.R = rune(matchParts[0][1])
		}
		for _, m := range matchParts {
			match := make([]int, 0, 2)
			for _, str := range strings.Split(m, " ") {
				v, _ := strconv.Atoi(str)
				match = append(match, v)
			}
			rule.M = append(rule.M, match)
		}
		rules[i] = rule
	}

	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}
	file.Close()

	resetCache()
	result := Task1(rules, messages, false)
	fmt.Printf("Task 1: %d\n", result)

	resetCache()
	result = Task2(rules, messages, false)
	fmt.Printf("Task 2: %d\n", result)
}
