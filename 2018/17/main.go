package main

import (
	"bufio"
	"fmt"
	"os"

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
	ssY, seY     int
}

func (p *Puzzle) AddVein(x1, y1, x2, y2 int) {
	oldW, oldH := p.gridW, p.gridH

	if x2+2 > p.gridW {
		p.gridW = x2 + 2
	}
	if y2+2 > p.gridH {
		p.gridH = y2 + 2
	}
	if y1 < p.ssY {
		p.ssY = y1
	}
	if y2 > p.seY {
		p.seY = y2
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
		for x := 0; x < p.gridW; x++ {
			r := p.grid[(y*p.gridW)+x]
			switch r {
			case '.':
				sandColor.Printf("%c", r)
			case '#':
				clayColor.Printf("%c", r)
			case '+', '|', '/', '\\':
				fallingWaterColor.Printf("%c", r)
			case '~':
				stillWaterColor.Printf("%c", r)
			default:
				if y == p.ssY {
					fmt.Printf("%c", '-')
					continue
				}
				if y == p.seY {
					fmt.Printf("%c", '-')
					continue
				}

				fmt.Printf("%c", r)
			}
		}
		fmt.Printf("\n")
	}
}

func (p *Puzzle) flow(x, y int) bool {
	for ; y < p.gridH; y++ {
		i := (y * p.gridW) + x

		r := p.grid[i]
		switch r {
		case '#', '~':
			y--
			i := (y * p.gridW)
			var rend int
			var rhole int
			for ; x < p.gridW; x++ {
				if p.grid[i+x] == '#' {
					rend = x - 1
					break
				}
				if p.grid[i+x] == '\\' || p.grid[i+x+p.gridW] == '.' {
					rhole = x
					break
				}
				if p.grid[i+x] == '.' {
					p.grid[i+x] = '|'
					return true
				}
			}
			x--
			var lend int
			var lhole int
			for ; x > 0; x-- {
				if p.grid[i+x] == '#' {
					lend = x + 1
					break
				}
				if p.grid[i+x] == '/' || p.grid[i+x+p.gridW] == '.' {
					lhole = x
					break
				}
				if p.grid[i+x] == '.' {
					p.grid[i+x] = '|'
					return true
				}
			}

			if rhole > 0 || lhole > 0 {
				if rhole > 0 {
					p.grid[i+rhole] = '\\'
					if p.flow(rhole, y+1) {
						return true
					}
				}
				if lhole > 0 {
					p.grid[i+lhole] = '/'
					if p.flow(lhole, y+1) {
						return true
					}
				}
				return false
			}

			for x = lend; x <= rend; x++ {
				p.grid[i+x] = '~'
			}

			return true
		case '.':
			p.grid[i] = '|'
			return true
		}
	}
	return false
}

func abs(a int) int {
	if a < 0 {
		return a
	}
	return a
}

func (p *Puzzle) Sim() {
	/*p.PrintState()
	fmt.Printf("\n")*/
	for p.flow(500, 0) {
		/*p.PrintState()
		fmt.Printf("\n")
		time.Sleep(100 * time.Millisecond)*/
	}

	count := 0
	for y := p.ssY; y < p.seY+1; y++ {
		for x := 0; x < p.gridW; x++ {
			r := p.grid[(y*p.gridW)+x]
			if r == '|' || r == '\\' || r == '/' || r == '~' {
				count++
			}
		}
	}
	p.PrintState()
	fmt.Println("DONE")
	fmt.Println("Total water:", count)
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	p := Puzzle{
		ssY: 99999,
	}
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
