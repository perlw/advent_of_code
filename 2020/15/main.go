package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func play(input []int, endAt int, debug bool) int {
	seen := make(map[int][2]int)
	for i := 0; i < len(input); i++ {
		seen[input[i]] = [2]int{i + 1, 0}
	}
	last := input[len(input)-1]
	for i := len(input); i < endAt; i++ {
		if debug {
			fmt.Printf("Turn %d: [%d]", i+1, last)
		}
		when, ok := seen[last]
		if !ok || when[1] == 0 {
			v := seen[0]
			v[0], v[1] = i+1, v[0]
			seen[0] = v
			last = 0
			if debug {
				fmt.Printf("speak zero @%+v\n", v)
			}
		} else {
			val := when[0] - when[1]
			v := seen[val]
			v[0], v[1] = i+1, v[0]
			seen[val] = v
			last = val
			if debug {
				fmt.Printf("speak %d @%+v\n", val, v)
			}
		}
	}
	return last
}

// Task1 ...
func Task1(input []int) int {
	return play(input, 2020, true)
}

// Task2 ...
func Task2(input []int) int {
	return play(input, 30000000, false)
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	nums := strings.Split(scanner.Text(), ",")
	input := make([]int, len(nums))
	for i := range nums {
		input[i], _ = strconv.Atoi(nums[i])
	}
	file.Close()

	result := Task1(input)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(input)
	fmt.Printf("Task 2: %d\n", result)
}
