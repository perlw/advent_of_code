package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/perlw/advent_of_code/toolkit/grid"
)

func testSlope(input *grid.Grid, sx, sy int) int {
	var count int

	var x int
	for y := 0; y < input.Height; y += sy {
		if input.Get(x, y).R == '#' {
			count++
		}
		x = (x + sx) % input.Width
	}

	return count
}

// Task1 ...
func Task1(input *grid.Grid) int {
	return testSlope(input, 3, 1)
}

// Task2 ...
func Task2(input *grid.Grid) int {
	slopes := [][]int{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}

	result := 1
	for _, slope := range slopes {
		result *= testSlope(input, slope[0], slope[1])
	}
	return result
}

func main() {
	var input []rune
	var width, height int
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, []rune(line)...)

		if width == 0 {
			width = len(line)
		}
		height++
	}
	file.Close()

	var g grid.Grid
	g.FromRunes(input)
	g.Width = width
	g.Height = height
	g.Label = "Map"

	result := Task1(g.Clone())
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(g.Clone())
	fmt.Printf("Task 1: %d\n", result)
}
