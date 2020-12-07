package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Rule ...
type Rule struct {
	Color    string
	Contains map[string]int
}

func findContainingColors(rules []Rule, target string) []string {
	found := make(map[string]int)
	for _, r := range rules {
		if _, ok := r.Contains[target]; ok {
			found[r.Color]++

			for _, i := range findContainingColors(rules, r.Color) {
				found[i]++
			}
		}
	}

	result := make([]string, 0, len(found))
	for i := range found {
		result = append(result, i)
	}
	return result
}

func findColorsContained(rules []Rule, target string) map[string]int {
	result := make(map[string]int)

	for _, r := range rules {
		if r.Color == target {
			for color, count := range r.Contains {
				result[color] += count

				found := findColorsContained(rules, color)
				for foundColor, foundCount := range found {
					result[foundColor] += foundCount * count
				}
			}

			break
		}
	}

	return result
}

// Task1 ...
func Task1(input []Rule) int {
	return len(findContainingColors(input, "shiny gold"))
}

// Task2 ...
func Task2(input []Rule) int {
	var result int
	for _, count := range findColorsContained(input, "shiny gold") {
		result += count
	}
	return result
}

func main() {
	input := make([]Rule, 0, 100)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		var color string

		line := scanner.Text()
		parts := strings.Split(line, " bags contain ")
		color = parts[0]

		parts = strings.Split(parts[1], ", ")
		rule := Rule{
			Color:    color,
			Contains: make(map[string]int),
		}
		for _, p := range parts {
			var c int
			var a, b string
			fmt.Sscanf(p, "%d %s %s", &c, &a, &b)
			rule.Contains[a+" "+b] = c
		}
		input = append(input, rule)
	}
	file.Close()

	result := Task1(input)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(input)
	fmt.Printf("Task 2: %d\n", result)
}
