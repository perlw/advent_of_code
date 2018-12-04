package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	gridX    = 1000
	gridY    = 1000
	gridSize = gridX * gridY
)

type Puzzle struct {
	Collides map[int]bool
	Grid     [gridSize]int
}

func (p *Puzzle) Add(num, x, y, w, h int) {
	if x+w > gridX || y+h > gridY {
		return
	}

	if p.Collides == nil {
		p.Collides = make(map[int]bool)
	}
	if _, ok := p.Collides[num]; !ok {
		p.Collides[num] = false
	}

	for yy := y; yy < y+h; yy++ {
		for xx := x; xx < x+w; xx++ {
			i := (yy * gridX) + xx
			if p.Grid[i] == 0 {
				p.Grid[i] = num
			} else {
				p.Collides[num] = true
				p.Collides[p.Grid[i]] = true
				p.Grid[i] = -1
			}
		}
	}
}

func (p *Puzzle) NumCollisions() int {
	collisions := 0
	for y := 0; y < gridY; y++ {
		for x := 0; x < gridX; x++ {
			i := (y * gridX) + x
			if p.Grid[i] > 1 {
				collisions++
			}
		}
	}
	return collisions
}

func (p *Puzzle) Spared() int {
	for id, collides := range p.Collides {
		if !collides {
			return id
		}
	}
	return -1
}

func (p Puzzle) String() string {
	output := strings.Builder{}
	output.Grow(gridSize)
	for y := 0; y < gridY; y++ {
		for x := 0; x < gridX; x++ {
			i := (y * gridX) + x
			switch p.Grid[i] {
			case 0:
				output.WriteByte('.')
			case -1:
				output.WriteByte('x')
			default:
				output.WriteByte('+')
			}
		}
		output.WriteByte('\n')
	}
	return output.String()
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	p := Puzzle{}
	for scanner.Scan() {
		var num, x, y, w, h int
		fmt.Sscanf(scanner.Text(), "#%d @ %d,%d: %dx%d", &num, &x, &y, &w, &h)
		p.Add(num, x, y, w, h)
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	fmt.Println("Collisions:", p.NumCollisions())
	fmt.Println("Spared:", p.Spared())
}
