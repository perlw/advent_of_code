package grid

import (
	"fmt"
	"testing"
)

func TestShouldPrint(t *testing.T) {
	g := Grid{
		Cells:  make([]Cell, 100),
		Width:  10,
		Height: 10,
		Label:  "test",
	}
	g.Cells[0].R = 'x'
	g.Cells[g.Width-1].R = 'x'
	g.Cells[g.Width*(g.Height-1)].R = 'x'
	g.Cells[(g.Width*g.Height)-1].R = 'x'
	fmt.Printf("%s\n", g)
}
