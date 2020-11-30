package astar

import (
	"fmt"
	"testing"

	"github.com/perlw/advent_of_code/toolkit/grid"
)

func TestShouldWorkWhenStartEqGoal(t *testing.T) {
	g := grid.Grid{
		Cells:  make([]grid.Cell, 100),
		Width:  10,
		Height: 10,
		Label:  "test",
	}

	_, pathGrid := FindPath(g, 4, 4, 4, 4, true)
	fmt.Printf("%s\n", pathGrid.String())
}

func TestShouldTakeStep(t *testing.T) {
	input := []rune{
		'#', '#', '#', '#', '#', '#', '#', '#', '#', '#',
		'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
		'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
		'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
		'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
		'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
		'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
		'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
		'#', ' ', ' ', ' ', ' ', ' ', ' ', ' ', ' ', '#',
		'#', '#', '#', '#', '#', '#', '#', '#', '#', '#',
	}
	g := grid.Grid{
		Cells:  make([]grid.Cell, 100),
		Width:  10,
		Height: 10,
		Label:  "test",
	}
	for i, r := range input {
		g.Cells[i].R = r
	}

	_, pathGrid := FindPath(g, 4, 4, 7, 4, true)
	fmt.Printf("%s\n", pathGrid.String())
}
