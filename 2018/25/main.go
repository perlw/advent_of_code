package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

/*
type ByInit []Group

func (b ByInit) Len() int {
	return len(b)
}

func (b ByInit) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByInit) Less(i, j int) bool {
	return b[i].Initiative > b[j].Initiative
}
*/

func waitEnter() {
	fmt.Printf("Press ENTER to continue\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for t := range a {
		if a[t] != b[t] {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func manhattan(a, b Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y) + abs(a.Z-b.Z) + abs(a.W-b.W)
}

func inSlice(a string, b []string) bool {
	for _, c := range b {
		if a == c {
			return true
		}
	}
	return false
}

type Point struct {
	X, Y, Z, W int
	C          int
}

type Puzzle struct {
	Points []Point
}

func NewPuzzle(filepath string) *Puzzle {
	p := Puzzle{
		Points: make([]Point, 0, 100),
	}

	input, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	scanner := bufio.NewScanner(input)

	count := 0
	for scanner.Scan() {
		var x, y, z, w int
		fmt.Sscanf(scanner.Text(), "%d,%d,%d,%d", &x, &y, &z, &w)
		p.Points = append(p.Points, Point{x, y, z, w, count})
		count++
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	input.Close()

	for i := 0; i < len(p.Points); i++ {
		pi := p.Points[i]
		for j := i + 1; j < len(p.Points); j++ {
			pj := p.Points[j]
			if manhattan(pi, pj) <= 3 && pi.C != pj.C {
				for k := 0; k < len(p.Points); k++ {
					if p.Points[k].C == pj.C {
						p.Points[k].C = pi.C
					}
				}
			}
		}
	}

	return &p
}

func (p *Puzzle) Part1() int {
	constellations := make(map[int]struct{})
	for _, p := range p.Points {
		constellations[p.C] = struct{}{}
	}
	return len(constellations)
}

func main() {
	p := NewPuzzle("input.txt")
	spew.Dump(p)
	fmt.Println(p.Part1())
}
