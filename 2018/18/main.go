package main

import (
	"bufio"
	"fmt"
	"os"

	col "github.com/fatih/color"
)

func waitEnter() {
	fmt.Printf("Press ENTER to continue\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func copyExpandGrid(currW, currH, newW, newH int, grid []rune) []rune {
	out := make([]rune, newW*newH)
	for y := 0; y < currH; y++ {
		for x := 0; x < currW; x++ {
			out[(y*newW)+x] = grid[(y*currW)+x]
		}
	}
	return out
}

var (
	openColor       = col.New(col.FgWhite)
	forestColor     = col.New(col.FgHiGreen)
	lumberyardColor = col.New(col.FgHiBlue)
)

type Puzzle struct {
	gridW, gridH int
	gridSize     int
	grid         []rune
	numLines     int
}

func NewPuzzle() *Puzzle {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	p := Puzzle{}
	for scanner.Scan() {
		p.Add(scanner.Bytes())
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	return &p
}

func (p *Puzzle) Add(b []byte) {
	oldW := p.gridW
	oldH := p.gridH
	oldSize := p.gridSize
	if len(b) > p.gridW {
		p.gridW = len(b)
		p.gridH = len(b)
	}
	p.gridSize = p.gridW * p.gridH
	if oldSize != p.gridSize {
		p.grid = copyExpandGrid(oldW, oldH, p.gridW, p.gridH, p.grid)
	}

	i := p.numLines * p.gridW
	for x, c := range b {
		p.grid[i+x] = rune(c)
	}
	p.numLines++
}

func (p *Puzzle) PrintState() {
	for y := 0; y < p.gridH; y++ {
		for x := 0; x < p.gridW; x++ {
			r := p.grid[(y*p.gridW)+x]
			switch r {
			case '.':
				openColor.Printf("%c", r)
			case '|':
				forestColor.Printf("%c", '♠')
			case '#':
				lumberyardColor.Printf("%c", '⌂')
			default:
				fmt.Printf("%c", r)
			}
		}
		fmt.Printf("\n")
	}
}

func (p *Puzzle) collect(sx, sy int) (int, int) {
	var trees, lumberyards int
	for y := sy - 1; y < sy+2; y++ {
		if y < 0 || y >= p.gridH {
			continue
		}

		i := y * p.gridW
		for x := sx - 1; x < sx+2; x++ {
			if x < 0 || x >= p.gridW {
				continue
			}
			if x == sx && y == sy {
				continue
			}

			switch p.grid[i+x] {
			case '|':
				trees++
			case '#':
				lumberyards++
			}
		}
	}
	return trees, lumberyards
}

func (p *Puzzle) StepSim() {
	tmp := make([]rune, p.gridSize)
	for y := 0; y < p.gridH; y++ {
		i := y * p.gridW
		for x := 0; x < p.gridW; x++ {
			trees, lumberyards := p.collect(x, y)

			switch p.grid[i+x] {
			case '.':
				if trees >= 3 {
					tmp[i+x] = '|'
					continue
				}
			case '|':
				if lumberyards >= 3 {
					tmp[i+x] = '#'
					continue
				}
			case '#':
				if trees == 0 || lumberyards == 0 {
					tmp[i+x] = '.'
					continue
				}
			}
			tmp[i+x] = p.grid[i+x]
		}
	}
	p.grid = tmp
}

func (p *Puzzle) Score() int {
	var trees, lumberyards int
	for _, r := range p.grid {
		if r == '#' {
			lumberyards++
		}
		if r == '|' {
			trees++
		}
	}
	return trees * lumberyards
}

func main() {
	// Solution 1
	fmt.Println("Solution 1")
	fmt.Println("==========")
	p := NewPuzzle()
	p.PrintState()
	waitEnter()
	for tick := 0; tick < 10; tick++ {
		p.StepSim()
		p.PrintState()
		waitEnter()
	}
	fmt.Println("Score:", p.Score())

	// Solution 2
	fmt.Println("\nSolution 2")
	fmt.Println("==========")
	p = NewPuzzle()
	p.PrintState()
	fmt.Printf("\n")
	seenGrids := make([][]rune, 0, 100000)
	var from, to int
	for tick := 0; tick < 1000000000; tick++ {
		p.StepSim()

		for prev, seen := range seenGrids {
			equal := true
			for t := range p.grid {
				if seen[t] != p.grid[t] {
					equal = false
				}
			}

			if equal {
				from = prev
				to = tick
				fmt.Printf("Found cycle between tick %d and %d, score: %d, offset: %d\n", prev, tick, p.Score())
				p.PrintState()
				waitEnter()
				break
			}
		}
		tmp := make([]rune, p.gridSize)
		copy(tmp, p.grid)
		seenGrids = append(seenGrids, tmp)

		if from > 0 && to > 0 {
			break
		}
	}

	dist := to - from
	offset := (1000000000 - from) % dist
	p = NewPuzzle()
	for t := 0; t < from+offset; t++ {
		p.StepSim()
	}
	p.PrintState()
	fmt.Println("Score after 1000000000 iterations:", p.Score())
}
