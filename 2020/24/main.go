package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

// Dir ...
type Dir int

// Dir constants.
const (
	DirNE Dir = iota
	DirE
	DirSE
	DirSW
	DirW
	DirNW
)

func (d Dir) String() string {
	switch d {
	case DirNE:
		return "ne"
	case DirE:
		return "e"
	case DirSE:
		return "se"
	case DirSW:
		return "sw"
	case DirW:
		return "w"
	case DirNW:
		return "nw"
	}
	return "n/a"
}

// MovedPos ...
func (d Dir) MovedPos(initial Pos) Pos {
	move := initial
	switch d {
	case DirNE:
		if move.Y%2 != 0 {
			move.X++
		}
		move.Y--
	case DirE:
		move.X++
	case DirSE:
		if move.Y%2 != 0 {
			move.X++
		}
		move.Y++
	case DirSW:
		if move.Y%2 == 0 {
			move.X--
		}
		move.Y++
	case DirW:
		move.X--
	case DirNW:
		if move.Y%2 == 0 {
			move.X--
		}
		move.Y--
	}
	return move
}

// StrToDir ...
func StrToDir(str string) Dir {
	switch str {
	case "ne":
		return DirNE
	case "e":
		return DirE
	case "se":
		return DirSE
	case "sw":
		return DirSW
	case "w":
		return DirW
	case "nw":
		return DirNW
	}
	return -1
}

// Cell ...
type Cell struct {
	Position Pos
	Flipped  bool
}

// Pos ...
type Pos struct {
	X, Y int
}

// HexGrid ...
type HexGrid struct {
	Root  *Cell
	Cells map[Pos]*Cell
}

// MakeMoves ...
func (hg *HexGrid) MakeMoves(moves string, debug bool) {
	if debug {
		fmt.Printf("\n==>%s\n", moves)
	}
	var currPos Pos
	for i := 0; i < len(moves); i++ {
		var dir Dir
		if moves[i] == 's' || moves[i] == 'n' {
			dir = StrToDir(moves[i : i+2])
			i++
		} else {
			dir = StrToDir(string(moves[i]))
		}

		move := dir.MovedPos(currPos)
		if debug {
			fmt.Printf("%+v move %s to %+v\n", currPos, dir, move)
		}

		cell, ok := hg.Cells[move]
		if !ok {
			cell = &Cell{
				Position: move,
			}
		}
		if i == len(moves)-1 {
			cell.Flipped = !cell.Flipped
			if debug {
				fmt.Printf("flipped %+v, now %v\n", cell.Position, cell.Flipped)
			}
		}
		hg.Cells[move] = cell
		currPos = move
	}
}

// IterateLife ...
func (hg *HexGrid) IterateLife(debug bool) {
	cells := make(map[Pos]*Cell)

	var topLeft, bottomRight Pos
	for _, c := range hg.Cells {
		if c.Position.X < topLeft.X {
			topLeft.X = c.Position.X
		}
		if c.Position.Y < topLeft.Y {
			topLeft.Y = c.Position.Y
		}
		if c.Position.X > bottomRight.X {
			bottomRight.X = c.Position.X
		}
		if c.Position.Y > bottomRight.Y {
			bottomRight.Y = c.Position.Y
		}
	}

	if debug {
		fmt.Printf("%+v -> %+v\n", topLeft, bottomRight)
	}
	// Adding in missing slots in grid to allow growth.
	for y := topLeft.Y - 1; y <= bottomRight.Y+1; y++ {
		for x := topLeft.X - 1; x <= bottomRight.X+1; x++ {
			pos := Pos{x, y}
			if _, ok := hg.Cells[pos]; !ok {
				hg.Cells[pos] = &Cell{
					Position: pos,
				}
			}
		}
	}

	for _, c := range hg.Cells {
		xw, xe, x, y := c.Position.X, c.Position.X, c.Position.X, c.Position.Y
		if y%2 != 0 {
			xe++
		}
		if y%2 == 0 {
			xw--
		}

		neighbors := []Pos{
			{xe, y - 1},
			{x + 1, y},
			{xe, y + 1},
			{xw, y + 1},
			{x - 1, y},
			{xw, y - 1},
		}
		var flipped int
		if debug {
			fmt.Printf("%+v is %v, neighbors...\n", c.Position, c.Flipped)
		}
		for i := range neighbors {
			cell, ok := hg.Cells[neighbors[i]]
			if debug {
				if ok {
					fmt.Printf("\t%+v is %v\n", cell.Position, cell.Flipped)
				} else {
					fmt.Printf("\t%+v [none] \n", neighbors[i])
				}
			}
			if ok && cell.Flipped {
				flipped++
			}
		}
		if debug {
			fmt.Printf("num flipped: %d\n", flipped)
		}

		cell := *c
		if (cell.Flipped && (flipped == 0 || flipped > 2)) ||
			(!cell.Flipped && flipped == 2) {
			if debug {
				fmt.Printf("=> flipping (%v and %d) <==\n", cell.Flipped, flipped)
			}
			cell.Flipped = !cell.Flipped
		}
		cells[cell.Position] = &cell
	}

	hg.Cells = cells
}

// Task1 ...
func Task1(input []string, debug bool) int {
	hg := HexGrid{
		Root:  &Cell{},
		Cells: make(map[Pos]*Cell),
	}
	for _, in := range input {
		hg.MakeMoves(in, debug)
	}

	var result int
	for _, c := range hg.Cells {
		if c.Flipped {
			result++
		}
	}
	return result
}

// Task2 ...
func Task2(input []string, iterate int, debug bool) int {
	hg := HexGrid{
		Root:  &Cell{},
		Cells: make(map[Pos]*Cell),
	}
	for _, in := range input {
		hg.MakeMoves(in, debug)
	}

	for i := 0; i < iterate; i++ {
		if debug {
			fmt.Printf("Day %d:", i+1)
		}
		hg.IterateLife(debug)
		if debug {
			var result int
			for _, c := range hg.Cells {
				if c.Flipped {
					result++
				}
			}
			fmt.Printf(" %2d\n", result)
		}
	}

	var result int
	for _, c := range hg.Cells {
		if c.Flipped {
			result++
		}
	}
	return result
}

func readInput(reader io.Reader) []string {
	input := make([]string, 0, 10)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	return input
}

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()

	var input []string
	file, _ := os.Open("input.txt")
	input = readInput(file)
	file.Close()

	result := Task1(input, debug)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(input, 100, debug)
	fmt.Printf("Task 2: %d\n", result)
}
