package main

import (
	"bufio"
	"fmt"
	"os"
)

func waitEnter() {
	fmt.Printf("Press ENTER to continue\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for t := range a {
		if a[t] != b[t] {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func manhattan(a, b Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y) + abs(a.Z-b.Z)
}

type Point struct {
	X, Y, Z int
}

type Puzzle struct {
}

func NewPuzzle(filepath string) *Puzzle {
	p := Puzzle{}

	input, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		var x, y, z, r int
		fmt.Sscanf(scanner.Text(), "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	input.Close()

	return &p
}

func main() {
	// Test1
	t := NewPuzzle("test1.txt")

	// Puzzle1
	/*p := NewPuzzle("input.txt")
	fmt.Println("In range:", len(p.DetectInRangeBot(p.FindStrongestBot())))
	fmt.Printf("Overlapping dist: %+v\n", p.FindMostOverlapDist())*/
}
