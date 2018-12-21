package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	col "github.com/fatih/color"
)

func waitEnter() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

var (
	sandColor         = col.New(col.FgRed)
	clayColor         = col.New(col.FgWhite)
	fallingWaterColor = col.New(col.FgHiBlue)
	stillWaterColor   = col.New(col.FgBlue)
)

type Puzzle struct {
	gridW, gridH int
	gridSize     int
	grid         []rune
}

func (p *Puzzle) AddVein(x1, y1, x2, y2 int) {
	oldW, oldH := p.gridW, p.gridH

	if x2+2 > p.gridW {
		p.gridW = x2 + 2
	}
	if y2+2 > p.gridH {
		p.gridH = y2 + 2
	}
	p.gridSize = p.gridW * p.gridH

	if p.grid == nil {
		p.grid = make([]rune, p.gridSize)
	}

	updated := make([]rune, p.gridSize)
	for y := 0; y < p.gridH; y++ {
		for x := 0; x < p.gridW; x++ {
			i := (y * p.gridW) + x
			if i == 500 {
				updated[i] = '+'
				continue
			}
			if x >= x1 && x <= x2 && y >= y1 && y <= y2 {
				updated[i] = '#'
				continue
			}
			if x < oldW && y < oldH {
				j := (y * oldW) + x
				updated[i] = p.grid[j]
				continue
			}
			updated[i] = '.'
		}
	}

	p.grid = updated
}

func (p *Puzzle) PrintState() {
	for y := 0; y < p.gridH; y++ {
		for x := 494; x < p.gridW; x++ {
			r := p.grid[(y*p.gridW)+x]
			switch r {
			case '.':
				sandColor.Printf("%c", r)
			case '#':
				clayColor.Printf("%c", r)
			case '+', '|':
				fallingWaterColor.Printf("%c", r)
			case '~':
				stillWaterColor.Printf("%c", r)
			default:
				fmt.Printf("%c", r)
			}
		}
		fmt.Printf("\n")
	}
}

func (p *Puzzle) Sim() {
	for {
		p.PrintState()
		fmt.Printf("\n")

		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	p := Puzzle{}
	for scanner.Scan() {
		var xs, xe, ys, ye int
		fmt.Sscanf(scanner.Text(), "x=%d, y=%d..%d", &xs, &ys, &ye)
		if xs == 0 && ys == 0 && ye == 0 {
			fmt.Sscanf(scanner.Text(), "y=%d, x=%d..%d", &ys, &xs, &xe)
			ye = ys
		} else {
			xe = xs
		}
		p.AddVein(xs, ys, xe, ye)
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	p.Sim()
}
