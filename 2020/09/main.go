package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func checkNumber(slice []int, number int) [][2]int {
	result := make([][2]int, 0, 25)
	for i := 0; i < len(slice); i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i]+slice[j] == number {
				result = append(result, [2]int{slice[i], slice[j]})
			}
		}
	}
	return result
}

// Task1 ...
func Task1(input []int, preamble int) int {
	for i := preamble; i < len(input); i++ {
		if len(checkNumber(input[i-preamble:i], input[i])) == 0 {
			return input[i]
		}
	}
	return -1
}

// Task2 ...
func Task2(input []int, target int) int {
	for i := 0; i < len(input); i++ {
		result, smallest, largest := input[i], input[i], input[i]
		for j := i + 1; j < len(input); j++ {
			result += input[j]
			if input[j] < smallest {
				smallest = input[j]
			}
			if input[j] > largest {
				largest = input[j]
			}
			if result == target {
				return smallest + largest
			} else if result > target {
				break
			}
		}
	}
	return -1
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	input := make([]int, 0, 100)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		input = append(input, val)
	}
	file.Close()

	result := Task1(input, 25)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(input, result)
	fmt.Printf("Task 2: %d\n", result)
}
