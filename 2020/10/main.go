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
	var ones int
	var threes int

	input = append([]int{0}, input...)
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)

	for i := 0; i < len(input)-1; i++ {
		jolt := input[i+1] - input[i]
		if jolt == 1 {
			ones++
		} else if jolt == 3 {
			threes++
		}
	}

	return ones * threes
}

func tri(n int) int {
	if n < 3 {
		return 0
	}
	if n == 3 {
		return 1
	}
	return tri(n-1) + tri(n-2) + tri(n-3)
}

// Task2 ...
func Task2(input []int) int {

	input = append([]int{0}, input...)
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)

	var marker int
	samples := make([][]int, 0, 100)
	var maxLen int
	for i := 0; i < len(input)-1; i++ {
		jolt := input[i+1] - input[i]
		if jolt == 3 {
			end := i
			sample := input[marker:end]
			if end == len(input)-1 {
				sample = input[marker:]
			}
			if len(sample) > maxLen {
				maxLen = len(sample)
			}
			samples = append(samples, sample)
			marker = i + 1
		}
	}

	triCached := make([]int, 0, 10)
	for i := 0; i < maxLen+1; i++ {
		triCached = append(triCached, tri(i+3))
	}

	result := 1
	for _, i := range samples {
		result *= triCached[len(i)]
	}
	return result
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

	result := Task1(input)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(input)
	fmt.Printf("Task 2: %d\n", result)
}
