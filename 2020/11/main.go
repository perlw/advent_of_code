package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// SeatMap ...
type SeatMap []rune

func (s SeatMap) ray(stride int, x, y int, dir [2]int) ([2]int, bool) {
	h := len(s) / stride
	x += dir[0]
	y += dir[1]
	for y >= 0 && y < h && x >= 0 && x < stride {
		r := s[(y*stride)+x]
		if r == '#' {
			return [2]int{x, y}, true
		} else if r == 'L' {
			break
		}

		x += dir[0]
		y += dir[1]
	}
	return [2]int{-1, -1}, false
}

// StepAdjacent ...
func (s SeatMap) StepAdjacent(stride int) bool {
	oldSeatMap := make(SeatMap, len(s))
	copy(oldSeatMap, s)

	var dirty bool
	for y := 0; y < len(s)/stride; y++ {
		i := y * stride
		for x := 0; x < stride; x++ {
			j := i + x
			if oldSeatMap[j] == '.' {
				continue
			}

			var numTakenSeats int
			for v := y - 1; v < y+2; v++ {
				for u := x - 1; u < x+2; u++ {
					if u == x && v == y {
						continue
					}

					if u >= 0 && u < stride && v >= 0 && v < len(s)/stride {
						if oldSeatMap[(v*stride)+u] == '#' {
							numTakenSeats++
						}
					}
				}
			}

			if oldSeatMap[j] == 'L' && numTakenSeats == 0 {
				s[j] = '#'
				dirty = true
			} else if oldSeatMap[j] == '#' && numTakenSeats >= 4 {
				s[j] = 'L'
				dirty = true
			}
		}
	}

	return dirty
}

// StepRay ...
func (s SeatMap) StepRay(stride int) bool {
	rays := [][2]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
	}

	oldSeatMap := make(SeatMap, len(s))
	copy(oldSeatMap, s)

	var dirty bool
	for y := 0; y < len(s)/stride; y++ {
		i := y * stride
		for x := 0; x < stride; x++ {
			j := i + x
			if oldSeatMap[j] == '.' {
				continue
			}

			var numTakenSeats int
			for _, ray := range rays {
				if _, ok := oldSeatMap.ray(stride, x, y, ray); ok {
					numTakenSeats++
				}
			}

			if oldSeatMap[j] == 'L' && numTakenSeats == 0 {
				s[j] = '#'
				dirty = true
			} else if oldSeatMap[j] == '#' && numTakenSeats >= 5 {
				s[j] = 'L'
				dirty = true
			}
		}
	}

	return dirty
}

func (s SeatMap) String(stride int) string {
	var result strings.Builder
	result.Grow(len(s) * 2)

	fmt.Fprintf(&result, "╭")
	for x := 1; x < stride+3; x++ {
		fmt.Fprintf(&result, "─")
	}
	fmt.Fprintf(&result, "╮\n")

	for y := 0; y < len(s)/stride; y++ {
		fmt.Fprintf(&result, "│ ")

		i := y * stride
		for x := 0; x < stride; x++ {
			fmt.Fprintf(&result, "%c", s[i+x])
		}

		fmt.Fprintf(&result, " │\n")
	}

	fmt.Fprintf(&result, "╰")
	for x := 1; x < stride+3; x++ {
		fmt.Fprintf(&result, "─")
	}
	fmt.Fprintf(&result, "╯")

	return result.String()
}

// Task1 ...
func Task1(seatMap SeatMap, stride int, print bool) int {
	if print {
		fmt.Println(seatMap.String(stride))
	}
	for seatMap.StepAdjacent(stride) {
		if print {
			fmt.Println(seatMap.String(stride))
		}
	}

	var result int
	for _, r := range seatMap {
		if r == '#' {
			result++
		}
	}
	return result
}

// Task2 ...
func Task2(seatMap SeatMap, stride int, print bool) int {
	if print {
		fmt.Println(seatMap.String(stride))
	}
	for seatMap.StepRay(stride) {
		if print {
			fmt.Println(seatMap.String(stride))
		}
	}

	var result int
	for _, r := range seatMap {
		if r == '#' {
			result++
		}
	}
	return result
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

		input = append(input, []rune(line)...)
	}
	file.Close()

	seatMap := make(SeatMap, len(input))
	copy(seatMap, input)
	result := Task1(seatMap, stride, false)
	fmt.Printf("Task 1: %d\n", result)

	copy(seatMap, input)
	result = Task2(seatMap, stride, false)
	fmt.Printf("Task 2: %d\n", result)
}
