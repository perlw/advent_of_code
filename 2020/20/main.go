package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/perlw/advent_of_code/toolkit/grid"
)

// Tiles ...
type Tiles map[int]grid.Grid

func rotateGrid(g *grid.Grid) {
	t := g.Clone()
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width; x++ {
			a := (y * g.Width) + x
			b := (x * g.Width) + (g.Width - y - 1)
			t.Cells[b] = g.Cells[a]
		}
	}
	*g = *t
}

func vflipGrid(g *grid.Grid) {
	for y := 0; y < g.Height; y++ {
		for x := 0; x < g.Width/2; x++ {
			a := (y * g.Width) + x
			b := (y * g.Width) + (g.Width - x - 1)
			g.Cells[a], g.Cells[b] = g.Cells[b], g.Cells[a]
		}
	}
}

func hflipGrid(g *grid.Grid) {
	for y := 0; y < g.Height/2; y++ {
		for x := 0; x < g.Width; x++ {
			a := (y * g.Width) + x
			b := ((g.Height - y - 1) * g.Width) + x
			g.Cells[a], g.Cells[b] = g.Cells[b], g.Cells[a]
		}
	}
}

// TODO: Typedeffed constants instead.
func compareEdges(a, b grid.Grid) (bool, bool, bool, bool) {
	lEdge, rEdge, tEdge, bEdge := true, true, true, true

	for x := 0; x < a.Width; x++ {
		if a.Cells[x] != b.Cells[x] {
			tEdge = false
			break
		}
	}
	for x := 0; x < a.Width; x++ {
		if a.Cells[((a.Height-1)*a.Height)+x] != b.Cells[((b.Height-1)*b.Height)+x] {
			bEdge = false
			break
		}
	}
	for y := 0; y < a.Height; y++ {
		if a.Cells[y*a.Width] != b.Cells[y*b.Width] {
			lEdge = false
			break
		}
	}
	for y := 0; y < a.Height; y++ {
		if a.Cells[(y*a.Width)+a.Width-1] != b.Cells[(y*b.Width)+b.Width-1] {
			rEdge = false
			break
		}
	}

	return lEdge, rEdge, tEdge, bEdge
}

// Task1 ...
func Task1(input Tiles) int {
	for i := range input {
		for j := range input {
			if j == i {
				continue
			}

			g := input[i]
			l, r, t, b := compareEdges(input[i], input[j])
			if l || r || t || b {
				fmt.Printf("potential %d -> %d\n", i, j)
			}

			hflipGrid(&g)
			l, r, t, b = compareEdges(input[i], input[j])
			if l || r || t || b {
				fmt.Printf("potential hflip %d -> %d\n", i, j)
			}

			vflipGrid(&g)
			l, r, t, b = compareEdges(input[i], input[j])
			if l || r || t || b {
				fmt.Printf("potential hvflip %d -> %d\n", i, j)
			}

			hflipGrid(&g)
			l, r, t, b = compareEdges(input[i], input[j])
			if l || r || t || b {
				fmt.Printf("potential vflip %d -> %d\n", i, j)
			}

			g = input[i]
			for u := 0; u < 3; u++ {
				rotateGrid(&g)
				l, r, t, b = compareEdges(input[i], input[j])
				if l || r || t || b {
					fmt.Printf("potential %drot %d -> %d\n", u+1, i, j)
				}
			}
		}
	}
	return 0
}

// Task2 ...
func Task2() int {
	return 0
}

func readInput(reader io.Reader) Tiles {
	input := make(Tiles)

	scanner := bufio.NewScanner(reader)
	var currentID int
	cells := make([]rune, 0, 10)
	var stride int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			input[currentID] = grid.FromRunes(cells, stride, stride)
			cells = make([]rune, 0, 10)
			continue
		}

		if strings.HasPrefix(line, "Tile") {
			fmt.Sscanf(line, "Tile %d", &currentID)
		} else {
			if len(cells) == 0 {
				stride = len(line)
			}
			cells = append(cells, []rune(line)...)
		}
	}

	return input
}

func main() {
	var input Tiles

	file, _ := os.Open("input.txt")
	input = readInput(file)
	file.Close()

	result := Task1(input)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2()
	fmt.Printf("Task 2: %d\n", result)
}
