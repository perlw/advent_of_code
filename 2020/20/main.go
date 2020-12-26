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

// Possibilities ...
type Possibilities struct {
	T map[int][]string
	B map[int][]string
	L map[int][]string
	R map[int][]string
}

// Count ...
func (p Possibilities) Count() int {
	return len(p.T) + len(p.B) + len(p.L) + len(p.R)
}

func addPossibility(
	p Possibilities, l, r, t, b bool, transform string, id int,
) Possibilities {
	if t && strings.Contains(transform, "h") {
		if _, ok := p.B[id]; !ok {
			p.B[id] = make([]string, 0, 10)
		}
		p.B[id] = append(p.B[id], transform)
	} else if t && !strings.Contains(transform, "h") {
		if _, ok := p.T[id]; !ok {
			p.T[id] = make([]string, 0, 10)
		}
		p.T[id] = append(p.T[id], transform)
	}
	if b && strings.Contains(transform, "h") {
		if _, ok := p.T[id]; !ok {
			p.T[id] = make([]string, 0, 10)
		}
		p.T[id] = append(p.T[id], transform)
	} else if b && !strings.Contains(transform, "h") {
		if _, ok := p.B[id]; !ok {
			p.B[id] = make([]string, 0, 10)
		}
		p.B[id] = append(p.B[id], transform)
	}
	if l && strings.Contains(transform, "v") {
		if _, ok := p.R[id]; !ok {
			p.R[id] = make([]string, 0, 10)
		}
		p.R[id] = append(p.R[id], transform)
	} else if l && !strings.Contains(transform, "v") {
		if _, ok := p.L[id]; !ok {
			p.L[id] = make([]string, 0, 10)
		}
		p.L[id] = append(p.L[id], transform)
	}
	if r && strings.Contains(transform, "v") {
		if _, ok := p.L[id]; !ok {
			p.L[id] = make([]string, 0, 10)
		}
		p.L[id] = append(p.L[id], transform)
	} else if r && !strings.Contains(transform, "v") {
		if _, ok := p.R[id]; !ok {
			p.R[id] = make([]string, 0, 10)
		}
		p.R[id] = append(p.R[id], transform)
	}

	return p
}

func (ts Tiles) buildPosssibilities(debug bool) map[int]Possibilities {
	result := make(map[int]Possibilities)
	for i := range ts {
		result[i] = Possibilities{
			T: make(map[int][]string),
			B: make(map[int][]string),
			L: make(map[int][]string),
			R: make(map[int][]string),
		}
		var count int
		for j := range ts {
			if j == i {
				continue
			}

			tile := ts[j]
			g := *(&tile).Clone()
			l, r, t, b := compareEdges(ts[i], g)
			if l || r || t || b {
				count++
				result[i] = addPossibility(result[i], l, r, t, b, "none", j)
				if debug {
					fmt.Printf("potential %d -> %d\n", i, j)
				}
			}

			g = *(&tile).Clone()
			hflipGrid(&g)
			l, r, t, b = compareEdges(ts[i], g)
			if l || r || t || b {
				count++
				result[i] = addPossibility(result[i], l, r, t, b, "hflip", j)
				if debug {
					fmt.Printf("potential hflip %d -> %d\n", i, j)
				}
			}

			g = *(&tile).Clone()
			vflipGrid(&g)
			l, r, t, b = compareEdges(ts[i], g)
			if l || r || t || b {
				count++
				result[i] = addPossibility(result[i], l, r, t, b, "vflip", j)
				if debug {
					fmt.Printf("potential vflip %d -> %d\n", i, j)
				}
			}

			g = *(&tile).Clone()
			hflipGrid(&g)
			vflipGrid(&g)
			l, r, t, b = compareEdges(ts[i], g)
			if l || r || t || b {
				count++
				result[i] = addPossibility(result[i], l, r, t, b, "hvflip", j)
				if debug {
					fmt.Printf("potential hvflip %d -> %d\n", i, j)
				}
			}

			g = *(&tile).Clone()
			rotateGrid(&g)
			l, r, t, b = compareEdges(ts[i], g)
			if l || r || t || b {
				count++
				result[i] = addPossibility(
					result[i], l, r, t, b, "rot", j,
				)
				if debug {
					fmt.Printf("potential rot %d -> %d\n", i, j)
				}
			}
			vflipGrid(&g)
			l, r, t, b = compareEdges(ts[i], g)
			if l || r || t || b {
				count++
				result[i] = addPossibility(
					result[i], l, r, t, b, "vfliprot", j,
				)
				if debug {
					fmt.Printf("potential vfliprot %d -> %d\n", i, j)
				}
			}

			hflipGrid(&g)
			l, r, t, b = compareEdges(ts[i], g)
			if l || r || t || b {
				count++
				result[i] = addPossibility(
					result[i], l, r, t, b, "hfliprot", j,
				)
				if debug {
					fmt.Printf("potential hfliprot %d -> %d\n", i, j)
				}
			}
			vflipGrid(&g)
			l, r, t, b = compareEdges(ts[i], g)
			if l || r || t || b {
				count++
				result[i] = addPossibility(
					result[i], l, r, t, b, "hvfliprot", j,
				)
				if debug {
					fmt.Printf("potential hvfliprot %d -> %d\n", i, j)
				}
			}
		}
		if debug {
			fmt.Printf("%d -> %d\n", i, count)
		}
	}

	return result
}

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

func compareEdges(a, b grid.Grid) (bool, bool, bool, bool) {
	lEdge, rEdge, tEdge, bEdge := true, true, true, true

	opEdge := (b.Height - 1) * b.Width
	for x := 0; x < a.Width; x++ {
		if a.Cells[x] != b.Cells[opEdge+x] {
			tEdge = false
			break
		}
	}
	for x := 0; x < a.Width; x++ {
		if a.Cells[((a.Height-1)*a.Height)+x] != b.Cells[x] {
			bEdge = false
			break
		}
	}
	opEdge = b.Width - 1
	for y := 0; y < a.Height; y++ {
		if a.Cells[y*a.Width] != b.Cells[(y*b.Width)+opEdge] {
			lEdge = false
			break
		}
	}
	for y := 0; y < a.Height; y++ {
		if a.Cells[(y*a.Width)+a.Width-1] != b.Cells[y*b.Width] {
			rEdge = false
			break
		}
	}

	return lEdge, rEdge, tEdge, bEdge
}

func inSlice(s []int, a int) bool {
	for i := range s {
		if s[i] == a {
			return true
		}
	}
	return false
}

// Task1 ...
func Task1(input Tiles, debug bool) int {
	possibilities := input.buildPosssibilities(debug)
	if debug {
		fmt.Printf("%+v\n", possibilities)
	}

	result := 1
	for i := range possibilities {
		if possibilities[i].Count() == 2 {
			result *= i
		}
	}
	return result
}

// Task2 ...
func Task2(input Tiles, debug bool) int {
	possibilities := input.buildPosssibilities(debug)
	if debug {
		fmt.Printf("%+v\n", possibilities)
	}

	corners := struct {
		tl, tr, bl, br int
	}{}
	fmt.Println()
	for i, p := range possibilities {
		if p.Count() == 2 {
			if len(p.T) > 0 && len(p.R) > 0 {
				fmt.Printf("tl: %d %+v\n", i, p)
				corners.tl = i
			}
			if len(p.T) > 0 && len(p.L) > 0 {
				fmt.Printf("tr: %d %+v\n", i, p)
				corners.tr = i
			}
			if len(p.B) > 0 && len(p.R) > 0 {
				fmt.Printf("bl: %d %+v\n", i, p)
				corners.bl = i
			}
			if len(p.B) > 0 && len(p.L) > 0 {
				fmt.Printf("br: %d %+v\n", i, p)
				corners.br = i
			}
		}
	}

	fmt.Printf("%+v\n", corners)
	// find corner to corner paths
	/*for _, i := range corners {
		fmt.Printf("searching from %d\n", i)
		c := i
	found:
		for {
		next:
			for j := range possibilities[c] {
				if c == j {
					continue
				}

				fmt.Printf("checking %d\n", j)
				for k := range possibilities[j] {
					if c == k {
						continue
					}
					if len(possibilities[k]) == 2 {
						fmt.Printf("corner %d\n", k)
						// TODO: store
						break found
					} else if len(possibilities[k]) == 3 {
						fmt.Printf("found next %d\n", k)
						c = k
						break next
					}
				}
			}
		}
	}*/

	return -1
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
	input[currentID] = grid.FromRunes(cells, stride, stride)

	return input
}

func main() {
	var input Tiles

	file, _ := os.Open("input.txt")
	input = readInput(file)
	file.Close()

	result := Task1(input, false)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(input, false)
	fmt.Printf("Task 2: %d\n", result)
}
