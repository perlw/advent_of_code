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

func findVariations(input []int) int {
	fmt.Printf("%+v", input)

	if len(input) <= 2 {
		return 1
	}
	if len(input) == 3 {
		if input[2]-input[0] <= 3 {
			return 2
		}
		return 1
	}

	return -1

	/*
		var count int
		for i := 1; i < len(input)-1; i++ {
			if input[i+1]-input[i-1] <= 3 {
				sub := append([]int{}, input[:i]...)
				sub = append(sub, input[i+1:]...)

				count += findVariations(sub, step+1)
			}
		}
		fmt.Printf(" +%d\n", count+step)
		return count + step
	*/
}

// Task2 ...
func Task2(input []int) int {
	variations := make([]int, 0, 10)

	input = append([]int{0}, input...)
	sort.Ints(input)
	input = append(input, input[len(input)-1]+3)

	fmt.Printf("%+v\n", input)

	var marker int
	for i := 0; i < len(input)-1; i++ {
		jolt := input[i+1] - input[i]
		if jolt == 3 {
			variations = append(variations, findVariations(input[marker:i+1]))
			marker = i + 1
		}
	}

	return -1
	/*
		fmt.Printf("%#v\n", variations)
		result := 1
		for _, i := range variations {
			result *= i
		}
		fmt.Printf("RES: %d\n", result)

		return result
	*/
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
