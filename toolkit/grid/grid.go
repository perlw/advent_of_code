package grid

import (
	"fmt"
	"strings"
)

// Cell represents a single cell in the grid.
type Cell struct {
	R rune
	// Color
}

// Grid is the struct holding the grid information.
type Grid struct {
	Cells  []Cell
	Width  int
	Height int
	Label  string
}

// TODO: Grow...
// TODO: Shrink...

// Set sets a cell at location.
func (g *Grid) Set(x, y int, c Cell) {
	g.Cells[(y*g.Width)+x] = c
}

// Get gets a cell at location.
func (g *Grid) Get(x, y int) Cell {
	return g.Cells[(y*g.Width)+x]
}

func (g Grid) String() string {
	var result strings.Builder
	result.Grow((g.Width * g.Height) * 2)

	fmt.Fprintf(&result, "╭")
	fmt.Fprintf(&result, "─ %s ─", g.Label)
	for x := len(g.Label) + 4; x < g.Width+2; x++ {
		fmt.Fprintf(&result, "─")
	}
	fmt.Fprintf(&result, "╮\n")

	for y := 0; y < g.Height; y++ {
		fmt.Fprintf(&result, "│ ")

		i := y * g.Width
		for x := 0; x < g.Width; x++ {
			r := g.Cells[i+x].R
			if r < ' ' {
				r = ' '
			}
			fmt.Fprintf(&result, "%c", r)
		}

		fmt.Fprintf(&result, " │\n")
	}

	fmt.Fprintf(&result, "╰")
	for x := 1; x < g.Width+3; x++ {
		fmt.Fprintf(&result, "─")
	}
	fmt.Fprintf(&result, "╯\n")

	return result.String()
}
