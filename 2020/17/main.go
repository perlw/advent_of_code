package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y, z int
}

// Task1 ...
func Task1(input []rune, stride int) int {
	fmt.Println(input)

	world := make(map[Point]bool)
	for step := 0; step < 1; step++ {
	}

	return -1
}

// Task2 ...
func Task2() int {
	return -1
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var stride int
	input := make([]rune, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()
		if stride == 0 {
			stride = len(line)
		}
		for _, r := range strings.Split(line, ",") {
			input = append(input, rune(r[0]))
		}
	}
	file.Close()

	result := Task1(input, stride)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2()
	fmt.Printf("Task 2: %d\n", result)
}
