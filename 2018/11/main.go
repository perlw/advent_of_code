package main

import (
	"fmt"
	"strings"
)

const (
	gridX = 300
	gridY = 300
)

var gridSize = gridX * gridY

type Puzzle struct {
	Serial int
	Cells  []int
}

func FuelAtCell(serial, x, y int) int {
	id := x + 10
	fuel := id * y
	fuel += serial
	fuel *= id
	fuel /= 100
	fuel %= 10
	fuel &= 0xf
	fuel -= 5
	return fuel
}

func NewPuzzle(serial int) Puzzle {
	p := Puzzle{
		Serial: serial,
		Cells:  make([]int, gridSize),
	}
	for y := 1; y < gridY+1; y++ {
		for x := 1; x < gridX+1; x++ {
			i := ((y - 1) * gridX) + (x - 1)
			p.Cells[i] = FuelAtCell(p.Serial, x, y)
		}
	}
	return p
}

func (p Puzzle) String() string {
	output := strings.Builder{}
	for y := 0; y < gridY; y++ {
		for x := 0; x < gridX; x++ {
			i := (y * gridX) + x
			output.WriteString(fmt.Sprintf("%3d", p.Cells[i]))
		}
		output.WriteByte('\n')
	}
	return output.String()
}

func (p *Puzzle) Solution1() (int, int) {
	var rx, ry int
	largest := 0
	for y := 0; y < gridY-3; y++ {
		for x := 0; x < gridX-3; x++ {
			total := 0
			for yy := y; yy < y+3; yy++ {
				for xx := x; xx < x+3; xx++ {
					i := (yy * gridY) + xx
					total += p.Cells[i]
				}
			}
			if total > largest {
				largest = total
				rx = x + 1
				ry = y + 1
			}
		}
	}

	return rx, ry
}

func (p *Puzzle) Solution2() (int, int, int) {
	var rx, ry int
	largestTotal := 0
	largestGrid := 0

	type chanStruct struct {
		total   int
		x, y, s int
	}
	done := make(chan chanStruct)
	for grid := 0; grid < 300; grid++ {
		go (func(grid int) {
			var rx, ry int
			largestTotal := 0

			fmt.Printf("Grid %dx%d\n", grid, grid)
			for y := 0; y < gridY-grid; y++ {
				for x := 0; x < gridX-grid; x++ {
					total := 0
					for yy := y; yy < y+grid; yy++ {
						for xx := x; xx < x+grid; xx++ {
							i := (yy * gridY) + xx
							total += p.Cells[i]
						}
					}
					if total > largestTotal {
						largestTotal = total
						rx = x + 1
						ry = y + 1
					}
				}
			}
			done <- chanStruct{
				total: largestTotal,
				x:     rx,
				y:     ry,
				s:     grid,
			}
		})(grid)
	}

	count := 0
	for count < 300 {
		result := <-done
		if result.total > largestTotal {
			largestTotal = result.total
			largestGrid = result.s
			rx = result.x
			ry = result.y
		}
		fmt.Printf("%d %dx%d -> %d,%d(%d)\n", count, result.s, result.s, result.x, result.y, result.total)
		count++
	}

	return rx, ry, largestGrid
}

func main() {
	p := NewPuzzle(7315)
	// p := NewPuzzle(42)
	fmt.Println(p.Solution1())
	fmt.Println(p.Solution2())
}
