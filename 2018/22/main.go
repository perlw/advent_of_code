package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
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

type Puzzle struct {
	gridW, gridH int
	gridSize     int
	grid         []int
	depth        int
	tX, tY       int
}

func (p *Puzzle) Solve(depth, tX, tY int) {
	p.gridW, p.gridH = tX+1, tY+1
	p.gridSize = p.gridW * p.gridH
	p.grid = make([]int, p.gridSize)
	p.depth = depth
	p.tX, p.tY = tX, tY

	for y := 0; y < p.gridH; y++ {
		i := y * p.gridW
		for x := 0; x < p.gridW; x++ {
			if x == tX && y == tY {
				p.grid[i+x] = 0
				continue
			}

			if y == 0 {
				p.grid[i+x] = x * 16807
			} else if x == 0 {
				p.grid[i+x] = y * 48271
			} else {
				p.grid[i+x] = p.grid[i+x-1] * p.grid[i+x-p.gridW]
			}
			p.grid[i+x] = (p.grid[i+x] + p.depth) % 20183
		}
	}
	p.grid[0] = 0
}

func (p *Puzzle) PrintState() {
	fmt.Printf("\033c")

	score := 0

	rocky := color.New(color.FgHiWhite)
	wet := color.New(color.FgHiBlue)
	narrow := color.New(color.FgWhite)
	for y := 0; y < p.gridH; y++ {
		i := y * p.gridW
		for x := 0; x < p.gridW; x++ {
			if i+x == 0 {
				fmt.Printf("%c", 'S')
				continue
			}
			if x == p.tX && y == p.tY {
				fmt.Printf("%c", 'T')
				continue
			}

			erosion := p.grid[i+x] % 3
			switch erosion {
			case 0:
				rocky.Printf("%c", '.')
			case 1:
				score++
				wet.Printf("%c", '=')
			case 2:
				score += 2
				narrow.Printf("%c", '|')
			default:
				fmt.Printf("%c", '?')
			}
		}
		fmt.Printf("\n")
	}

	fmt.Println("Risk level:", score)
}

func main() {
	// Test
	t := Puzzle{}
	t.Solve(510, 10, 10)
	t.PrintState()

	// Puzzle
	p := Puzzle{}
	p.Solve(9171, 7, 721)
	p.PrintState()
}
