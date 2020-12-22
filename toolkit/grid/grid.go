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

// FromRunes creates a new grid from a rune slice.
func FromRunes(runes []rune, width, height int) Grid {
	g := Grid{
		Width:  width,
		Height: height,
	}
	g.Cells = make([]Cell, len(runes))
	for i, r := range runes {
		g.Cells[i].R = r
	}
	return g
}

// TODO: Grow.
// TODO: Shrink.
// TODO: Cellmap for prettyprints.

// FromRunes reads in cells from a rune-slice.
// NOTE: Does not update width/height.
func (g *Grid) FromRunes(runes []rune) {
	g.Cells = make([]Cell, len(runes))
	for i, r := range runes {
		g.Cells[i].R = r
	}
}

// Clone returns a deep copy of the grid.
func (g *Grid) Clone() *Grid {
	clone := Grid{
		Cells:  make([]Cell, len(g.Cells)),
		Width:  g.Width,
		Height: g.Height,
		Label:  g.Label,
	}
	copy(clone.Cells, g.Cells)
	return &clone
}

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
	fmt.Fprintf(&result, "╯")

	return result.String()
}
