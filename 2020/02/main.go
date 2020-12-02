package main

import (
	"bufio"
	"fmt"
	"os"
)

// Input ...
type Input struct {
	ruleMin  int
	ruleMax  int
	ruleRune rune
	password string
}

// Task1 ...
func Task1(input []Input) int {
	var result int

	for _, i := range input {
		var count int
		for _, r := range i.password {
			if r == i.ruleRune {
				count++
			}
		}
		if count >= i.ruleMin && count <= i.ruleMax {
			result++
		}
	}

	return result
}

// Task2 ...
func Task2(input []Input) int {
	var result int

	for _, i := range input {
		a := rune(i.password[i.ruleMin-1])
		b := rune(i.password[i.ruleMax-1])
		if (a == i.ruleRune) != (b == i.ruleRune) {
			result++
		}
	}

	return result
}

func main() {
	input := make([]Input, 0, 100)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var i Input
		fmt.Sscanf(
			scanner.Text(), "%d-%d %c: %s",
			&i.ruleMin, &i.ruleMax, &i.ruleRune, &i.password,
		)
		input = append(input, i)
	}
	file.Close()

	input1 := make([]Input, len(input))
	copy(input1, input)
	result := Task1(input1)
	fmt.Printf("Task 1: %d\n", result)

	input2 := make([]Input, len(input))
	copy(input2, input)
	result = Task2(input2)
	fmt.Printf("Task 2: %d\n", result)
}
