package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func search(str string, min, max int) int {
	if max-min <= 2 {
		if str[0] == 'F' || str[0] == 'L' {
			return min
		}
		return max
	}

	step := int(math.Round(float64(max-min) / 2))
	switch str[0] {
	case 'F', 'L':
		max -= step
	case 'B', 'R':
		min += step
	}
	return search(str[1:], min, max)
}

// Seat ...
type Seat string

// GetRow ...
func (s Seat) GetRow() int {
	return search(string(s)[:7], 0, 127)
}

// GetColumn ...
func (s Seat) GetColumn() int {
	return search(string(s)[7:], 0, 7)
}

// GetID ...
func (s Seat) GetID() int {
	return (s.GetRow() * 8) + s.GetColumn()
}

/*func (s Seat) String() string {
	return "n/a"
}*/

// Task1 ...
func Task1(input []Seat) int {
	var result int
	for _, i := range input {
		if id := i.GetID(); id > result {
			result = id
		}
	}
	return result
}

// Task2 ...
func Task2(input []Seat) int {
	seatmap := make([]int, len(input))

	for i, in := range input {
		seatmap[i] = in.GetID()
	}
	sort.Ints(seatmap)
	for i := 0; i < len(seatmap)-2; i++ {
		if seatmap[i+1]-seatmap[i] != 1 {
			return seatmap[i] + 1
		}
	}

	return -1
}

func main() {
	var input []Seat
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		input = append(input, Seat(scanner.Text()))
	}
	file.Close()

	result := Task1(input)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(input)
	fmt.Printf("Task 2: %d\n", result)
}
