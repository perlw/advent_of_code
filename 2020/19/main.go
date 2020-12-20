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
	R     rune
	Match [][]int
}

// Task1 ...
func Task1(rules map[int]Rule, messages []string) int {
	var result int
	return result
}

// Task2 ...
func Task2() int {
	var result int
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
			Match: make([][]int, 0, 10),
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
			rule.Match = append(rule.Match, match)
		}
		rules[i] = rule
	}

	for scanner.Scan() {
		messages = append(messages, scanner.Text())
	}
	file.Close()

	result := Task1(rules, messages)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2()
	fmt.Printf("Task 2: %d\n", result)
}
