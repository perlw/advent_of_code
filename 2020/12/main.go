package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Task1 ...
func Task1(input []string) int {
	var x, y int
	dir := 0
	dx := map[int][2]int{
		0:   {1, 0},
		90:  {0, 1},
		180: {-1, 0},
		270: {0, -1},
	}

	for _, instr := range input {
		val, _ := strconv.Atoi(instr[1:])
		switch instr[0] {
		case 'N':
			y -= val
		case 'S':
			y += val
		case 'E':
			x += val
		case 'W':
			x -= val
		case 'L':
			dir -= val
			if dir < 0 {
				dir += 360
			}
		case 'R':
			dir += val
			if dir > 359 {
				dir -= 360
			}
		case 'F':
			d := dx[dir]
			x += val * d[0]
			y += val * d[1]
		}
	}

	return x + y
}

// Task2 ...
func Task2(input []string) int {
	var x, y int
	wx, wy := 10, -1

	for _, instr := range input {
		val, _ := strconv.Atoi(instr[1:])
		switch instr[0] {
		case 'N':
			wy -= val
		case 'S':
			wy += val
		case 'E':
			wx += val
		case 'W':
			wx -= val
		case 'L':
			switch val {
			case 90:
				wx, wy = wy, -wx
			case 180:
				wx, wy = -wx, -wy
			case 270:
				wx, wy = -wy, wx
			}
		case 'R':
			switch val {
			case 90:
				wx, wy = -wy, wx
			case 180:
				wx, wy = -wx, -wy
			case 270:
				wx, wy = wy, -wx
			}
		case 'F':
			x += val * wx
			y += val * wy
		}
	}

	return x + y
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	input := make([]string, 0, 100)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	file.Close()

	result := Task1(input)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(input)
	fmt.Printf("Task 2: %d\n", result)
}
