package main

import (
	"bufio"
	"fmt"
	"os"
)

func countUniqueAnswers(answers []string) int {
	found := make(map[rune]int)
	for _, a := range answers {
		for _, r := range a {
			found[r]++
		}
	}
	return len(found)
}

func countAllSharedAnswers(answers []string) int {
	found := make(map[rune]int)
	for _, a := range answers {
		for _, r := range a {
			found[r]++
		}
	}

	var count int
	for _, i := range found {
		if i == len(answers) {
			count++
		}
	}
	return count
}

// Task1 ...
func Task1(input [][]string) int {
	var result int
	for _, answers := range input {
		result += countUniqueAnswers(answers)
	}
	return result
}

// Task2 ...
func Task2(input [][]string) int {
	var result int
	for _, answers := range input {
		result += countAllSharedAnswers(answers)
	}
	return result
}

func main() {
	input := make([][]string, 0, 100)
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	work := make([]string, 0, 10)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			input = append(input, work)
			work = make([]string, 0, 10)
			continue
		}
		work = append(work, scanner.Text())
	}
	file.Close()
	if len(work) > 0 {
		input = append(input, work)
	}

	result := Task1(input)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(input)
	fmt.Printf("Task 2: %d\n", result)
}
