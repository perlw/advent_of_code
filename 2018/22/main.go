package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

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

type Tool int

const (
	ToolNeither Tool = iota
	ToolTorch
	ToolClimbing
)

type Puzzle struct {
	gridW, gridH int
	gridSize     int
	grid         []int
	depth        int
	tX, tY       int
}

func (p *Puzzle) Solve(depth, tX, tY int) {
	p.gridW, p.gridH = tX+10, tY+10
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

func pickTool(kind int, tool Tool) (Tool, bool) {
	switch kind {
	case 0:
		return ToolClimbing, (tool == ToolClimbing)
	case 1:
		return ToolClimbing, (tool == ToolNeither)
	case 2:
		return ToolTorch, (tool == ToolTorch)
	}
	return ToolNeither, (tool == ToolNeither)
}

func (p *Puzzle) FindPath(pretty bool) {
	grid := make([]int, p.gridSize)
	relScores := make([]int, p.gridSize)
	for y := 0; y < p.gridH; y++ {
		i := y * p.gridW
		for x := 0; x < p.gridW; x++ {
			grid[i+x] = 0
		}
	}

	currentTool := ToolNeither
	current := 0
	target := (p.tY * p.gridW) + p.tX
	for {
		up := current - p.gridW
		dn := current + p.gridW
		lt := current - 1
		rt := current + 1

		if up == target || dn == target || lt == target || rt == target {
			_, change := pickTool(p.grid[up]%3, currentTool)
			score := 1
			if change {
				score = 7
			}
			grid[target] = grid[current] + score
			relScores[target] = score
			grid[current] += 9000
			current = target
			break
		}

		if up >= 0 && grid[up] == 0 {
			_, change := pickTool(p.grid[up]%3, currentTool)
			score := 1
			if change {
				score = 7
			}
			relScores[up] = score
			grid[up] = grid[current] + score
		}
		if lt >= 0 && current%p.gridW > 0 && grid[lt] == 0 {
			_, change := pickTool(p.grid[lt]%3, currentTool)
			score := 1
			if change {
				score = 7
			}
			relScores[lt] = score
			grid[lt] = grid[current] + score
		}
		if rt < p.gridSize && current+1%p.gridW != 0 && grid[rt] == 0 {
			_, change := pickTool(p.grid[rt]%3, currentTool)
			score := 1
			if change {
				score = 7
			}
			relScores[rt] = score
			grid[rt] = grid[current] + score
		}
		if dn < p.gridSize && grid[dn] == 0 {
			_, change := pickTool(p.grid[dn]%3, currentTool)
			score := 1
			if change {
				score = 7
			}
			relScores[dn] = score
			grid[dn] = grid[current] + score
		}
		grid[current] += 9000

		lowest := 8999
		for i, score := range grid {
			if score > 0 && score < lowest {
				current = i
				lowest = score
			}
		}
		if lowest == 8999 {
			break
		}
		currentTool, _ = pickTool(p.grid[current]%3, currentTool)

		// Print path
		if pretty {
			fmt.Printf("\033c")
			for y := 0; y < p.gridH; y++ {
				i := y * p.gridW
				for x := 0; x < p.gridW; x++ {
					if x == 0 && y == 0 {
						fmt.Printf("S")
						continue
					}
					if x == p.tX && y == p.tY {
						fmt.Printf("T")
						continue
					}
					if grid[i+x] >= 9000 {
						fmt.Printf("o")
						continue
					}
					if grid[i+x] > 0 {
						fmt.Printf("+")
						continue
					}

					fmt.Printf(".")
				}
				fmt.Printf("\n")
			}
			time.Sleep(10 * time.Millisecond)
		}
	}

	current = target
	for {
		lowest := 99999

		up := current - p.gridW
		dn := current + p.gridW
		lt := current - 1
		rt := current + 1

		if up == 0 || dn == 0 || lt == 0 || rt == 0 {
			break
		}

		if up >= 0 && grid[up] >= 9000 && grid[up] < lowest {
			lowest = grid[up]
			current = up
		}
		if lt >= 0 && current%p.gridW > 0 && grid[lt] >= 9000 && grid[lt] < lowest {
			lowest = grid[lt]
			current = lt
		}
		if rt < p.gridSize && current+1%p.gridW != 0 && grid[rt] >= 9000 && grid[rt] < lowest {
			lowest = grid[rt]
			current = rt
		}
		if dn < p.gridSize && grid[dn] >= 9000 && grid[dn] < lowest {
			lowest = grid[rt]
			current = rt
		}
		grid[current] = 100000

		// Print path
		if pretty {
			fmt.Printf("\033c")
			for y := 0; y < p.gridH; y++ {
				i := y * p.gridW
				for x := 0; x < p.gridW; x++ {
					if x == 0 && y == 0 {
						fmt.Printf("S")
						continue
					}
					if x == p.tX && y == p.tY {
						fmt.Printf("T")
						continue
					}
					if grid[i+x] == 100000 {
						fmt.Printf("*")
						continue
					}
					if grid[i+x] >= 9000 {
						fmt.Printf("o")
						continue
					}
					if grid[i+x] > 0 {
						fmt.Printf("+")
						continue
					}

					fmt.Printf(".")
				}
				fmt.Printf("\n")
			}
			time.Sleep(10 * time.Millisecond)
		}
	}

	// Print path
	score := 0
	p.PrintState()
	for y := 0; y < p.gridH; y++ {
		i := y * p.gridW
		for x := 0; x < p.gridW; x++ {
			if x == 0 && y == 0 {
				fmt.Printf("S")
				continue
			}
			if x == p.tX && y == p.tY {
				fmt.Printf("T")
				continue
			}
			if grid[i+x] == 100000 {
				fmt.Printf("*")
				score += relScores[i+x]
				continue
			}
			if grid[i+x] >= 9000 {
				fmt.Printf("o")
				continue
			}
			if grid[i+x] > 0 {
				fmt.Printf("+")
				continue
			}

			fmt.Printf(".")
		}
		fmt.Printf("\n")
	}
	fmt.Printf("\n")
	for y := 0; y < p.gridH; y++ {
		i := y * p.gridW
		for x := 0; x < p.gridW; x++ {
			fmt.Printf("%d", relScores[i+x])
		}
		fmt.Printf("\n")
	}
	fmt.Printf("Minutes: %d\n", score)
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
	t.FindPath(true)

	// Puzzle
	p := Puzzle{}
	p.Solve(9171, 7, 721)
	p.FindPath(false)
}
