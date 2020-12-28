package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

func findLoopSize(subject int, target int, debug bool) int {
	val := 1
	loopSize := 1
	for {
		val *= subject
		val %= 20201227
		if debug {
			fmt.Printf("%d\n", val)
		}
		if val == target {
			break
		}
		loopSize++
	}
	return loopSize
}

func transformKey(subject, loopSize int) int {
	val := 1
	for i := 0; i < loopSize; i++ {
		val *= subject
		val %= 20201227
	}
	return val
}

// Task1 ...
func Task1(input []int, debug bool) int {
	subjectNumber := 7
	loopA := findLoopSize(subjectNumber, input[0], debug)
	loopB := findLoopSize(subjectNumber, input[1], debug)

	fmt.Printf("loopSize for %d: %d\n", input[0], loopA)
	fmt.Printf("loopSize for %d: %d\n", input[1], loopB)

	a := transformKey(input[0], loopB)
	b := transformKey(input[1], loopA)

	fmt.Printf("encryption for %d@%d: %d\n", input[0], loopB, a)
	fmt.Printf("encryption for %d@%d: %d\n", input[1], loopA, b)

	if a == b {
		return a
	}
	return -1
}

// Task2 ...
func Task2(input []int, debug bool) int {
	return -1
}

func readInput(reader io.Reader) []int {
	input := make([]int, 0, 10)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		input = append(input, val)
	}
	return input
}

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()

	file, _ := os.Open("input.txt")
	input := readInput(file)
	file.Close()

	result := Task1(input, debug)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(input, debug)
	fmt.Printf("Task 2: %d\n", result)
}
