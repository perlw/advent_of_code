package main

import (
	"bufio"
	"fmt"
	"os"
)

func waitEnter() {
	fmt.Printf("Press ENTER to continue\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

type DoorDir int

const (
	DoorNorth DoorDir = iota
	DoorWest
	DoorEast
	DoorSouth
)

type Puzzle struct {
	gridW, gridH int
	gridSize     int
	grid         []byte
	cX, cY       int
}

func NewPuzzle() *Puzzle {
	return &Puzzle{
		gridW:    3,
		gridH:    3,
		gridSize: 3,
		grid:     []byte{'#', '?', '#', '?', 'X', '?', '#', '?', '#'},
		cX:       1,
		cY:       1,
	}
}

func copyGrid(dst []byte, src []byte, dx, dy, dw, dh, sw, sh int) {
	for y := 0; y < sh; y++ {
		di := ((y + dy) * dw) + dx
		si := y * sw
		for x := 0; x < sw; x++ {
			dst[di+x] = src[si+x]
		}
	}
}

func (p *Puzzle) walk(door DoorDir) {
	var nx, ny int
	switch door {
	case DoorNorth:
		ny = -2
	case DoorWest:
		nx = -2
	case DoorEast:
		nx = 2
	case DoorSouth:
		ny = 2
	}

	oldW, oldH := p.gridW, p.gridH
	oldX, oldY := p.cX, p.cY

	p.cX += nx
	p.cY += ny

	if p.cX < 0 || p.cY < 0 || p.cX >= p.gridW || p.cY >= p.gridH {
		var dx, dy int
		switch door {
		case DoorNorth, DoorSouth:
			p.gridH += 2
			dy = 2
		case DoorWest, DoorEast:
			p.gridW += 2
			dx = 2
		}
		p.gridSize = p.gridW * p.gridH

		newGrid := make([]byte, p.gridSize)
		copyGrid(newGrid, p.grid, dx, dy, p.gridW, p.gridH, oldW, oldH)
		p.grid = newGrid

		p.cX = oldX
		p.cY = oldY
	}

	room := []byte{'#', '?', '#', '?', '.', '?', '#', '?', '#'}
	for y := 0; y < 3; y++ {
		i := ((y + p.cY - 1) * p.gridW) + p.cX - 1
		ri := y * 3
		for x := 0; x < 3; x++ {
			p.grid[i+x] = room[ri+x]
		}
	}

	switch door {
	case DoorNorth, DoorSouth:
		p.grid[((p.cY-(ny/2))*p.gridW)+(p.cX-(nx/2))] = '-'
	case DoorWest, DoorEast:
		p.grid[((p.cY-(ny/2))*p.gridW)+(p.cX-(nx/2))] = '|'
	}

	p.PrintMap()
	waitEnter()
}

func (p *Puzzle) PrintMap() {
	for y := 0; y < p.gridH; y++ {
		i := y * p.gridW
		for x := 0; x < p.gridW; x++ {
			fmt.Printf("%c", p.grid[i+x])
		}
		fmt.Printf("\n")
	}
}

func (p *Puzzle) Map(directions []byte) {
	for _, c := range directions {
		switch c {
		case '^':
		case '$':
		default:
			switch c {
			case 'N':
				p.walk(DoorNorth)
			case 'W':
				p.walk(DoorWest)
			case 'E':
				p.walk(DoorEast)
			case 'S':
				p.walk(DoorSouth)
			}
		}
	}

	for t := range p.grid {
		if p.grid[t] == '?' {
			p.grid[t] = '#'
		}
	}

	p.PrintMap()
}

func (p *Puzzle) MapFile(filepath string) {
	/*
		input, err := os.Open("input.txt")
		if err != nil {
			panic(err.Error())
		}
		defer input.Close()
		scanner := bufio.NewScanner(input)

		for scanner.Scan() {
			var op string
			var a, b, c int
			fmt.Sscanf(scanner.Text(), "%s %d %d %d", &op, &a, &b, &c)
		}
		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}
	*/
}

func main() {
	// Tests
	p := NewPuzzle()
	p.Map([]byte("^WNE$"))
	waitEnter()

	p = NewPuzzle()
	p.MapFile("input.txt")
}
