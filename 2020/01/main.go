package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Task1 ...
func Task1(input []int) int {
	sort.Ints(input)

	for i := len(input) - 1; i >= 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if input[i]+input[j] == 2020 {
				return input[i] * input[j]
			}
		}
	}

	return -1
}

// Task2 ...
func Task2(input []int) int {
	sort.Ints(input)

	for i := len(input) - 1; i >= 0; i-- {
		icurrent := input[i]
		for j := i - 1; j >= 0; j-- {
			jcurrent := icurrent + input[j]
			for k := j - 1; k >= 0; k-- {
				if jcurrent+input[k] == 2020 {
					return input[i] * input[j] * input[k]
				}
			}
		}
	}

	return -1
}

func main() {
	input := make([]int, 0, 100)

	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		input = append(input, val)
	}
	file.Close()

	input1 := make([]int, len(input))
	copy(input1, input)
	result := Task1(input1)
	fmt.Printf("Task 1: %d\n", result)

	input2 := make([]int, len(input))
	copy(input2, input)
	result = Task2(input2)
	fmt.Printf("Task 2: %d\n", result)
}
