package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func digitsToString(input []int) string {
	var str string
	for i := range input {
		str += fmt.Sprintf("%d", input[i])
	}
	return str
}

func digitsToInt(input []int) int {
	var result int
	tens := 1000000
	for i := range input {
		result += input[i] * tens
		tens /= 10
	}
	return result
}

func getPattern(phase int) []int {
	base := []int{0, 1, 0, -1}

	pattern := make([]int, 0)
	for i := range base {
		for j := 0; j < phase; j++ {
			pattern = append(pattern, base[i])
		}
	}
	return pattern
}

func phase(input []int, debug bool) []int {
	result := make([]int, len(input))
	for i := range input {
		pattern := getPattern(i + 1)
		p := 1
		for j := range input {
			var prefix string
			if j > 0 {
				prefix = " + "
			}
			result[i] += input[j] * pattern[p]
			if debug {
				fmt.Printf("%s%d*%2d", prefix, input[j], pattern[p])
			}
			p = (p + 1) % len(pattern)
		}
		if result[i] < 0 {
			result[i] = -result[i]
		}
		if result[i] > 9 {
			result[i] %= 10
		}
		if debug {
			fmt.Printf(" = %d\n", result[i])
		}
	}
	return result
}

func FFT(input []int, phases int, debug bool) []int {
	result := make([]int, len(input))
	copy(result, input)

	if debug {
		fmt.Printf("Input signal: %s\n\n", digitsToString(input))
	}
	var current int
	for current < phases {
		current++
		result = phase(result, debug)
		if debug {
			fmt.Printf("\nAfter %d phase: %s\n\n", current, digitsToString(result))
		}
	}

	return result
}

func FFT2(input []int, phases, offset int, debug bool) []int {
	length := (len(input) * 10000) - offset
	result := make([]int, length)

	step := offset
	for i := 0; i < length; i++ {
		result[i] = input[step%len(input)]
		step++
	}

	if debug {
		fmt.Printf("Input signal: %s\n\n", digitsToString(input))
	}
	var current int
	for current < phases {
		current++

		for i := length - 2; i >= 0; i-- {
			result[i] = (result[i] + result[i+1]) % 10
		}

		if debug {
			fmt.Printf("\nAfter %d phase: %s\n\n", current, digitsToString(result))
		}
	}

	return result
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("could not open input file: %w", err))
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)

	digits := make([]int, 0)
	for scanner.Scan() {
		r := scanner.Bytes()
		for i := range r {
			digits = append(digits, int(r[i]-48))
		}
	}

	p1 := FFT(digits, 100, false)
	p2 := FFT2(digits, 100, digitsToInt(digits[:7]), false)
	fmt.Printf("Part 1: %s\n", digitsToString(p1[:8]))
	fmt.Printf("Part 2: %s\n", digitsToString(p2[:8]))
}
