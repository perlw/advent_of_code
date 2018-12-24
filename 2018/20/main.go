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

func (p *Puzzle) walk(door DoorDir) (int, int) {
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

	var dx, dy int
	if p.cX < 0 || p.cY < 0 || p.cX >= p.gridW || p.cY >= p.gridH {
		switch door {
		case DoorNorth:
			p.cY = oldY
			dy = 2
			fallthrough
		case DoorSouth:
			p.gridH += 2
		case DoorWest:
			p.cX = oldX
			dx = 2
			fallthrough
		case DoorEast:
			p.gridW += 2
		}
		p.gridSize = p.gridW * p.gridH

		newGrid := make([]byte, p.gridSize)
		copyGrid(newGrid, p.grid, dx, dy, p.gridW, p.gridH, oldW, oldH)
		p.grid = newGrid
	}

	room := []byte{'#', '?', '#', '?', '.', '?', '#', '?', '#'}
	for y := 0; y < 3; y++ {
		i := ((y + p.cY - 1) * p.gridW) + p.cX - 1
		ri := y * 3
		for x := 0; x < 3; x++ {
			if p.grid[i+x] == 0 {
				p.grid[i+x] = room[ri+x]
			}
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

	return dx, dy
}

func (p *Puzzle) PrintMap() {
	for y := 0; y < p.gridH; y++ {
		i := y * p.gridW
		for x := 0; x < p.gridW; x++ {
			c := p.grid[i+x]
			if c == 0 {
				c = ' '
			}
			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
}

func (p *Puzzle) Map(directions []byte) (int, int) {
	var oX, oY int
	fmt.Println("run", string(directions))
	for t := 0; t < len(directions); t++ {
		c := directions[t]
		switch c {
		case '^':
		case '$':
			for t := range p.grid {
				if p.grid[t] == '?' {
					p.grid[t] = '#'
				}
			}
			break
		case '(':
			var branch []byte
			fmt.Printf("branch at %d,%d index %d\n", p.cX, p.cY, t)
			count := 1
			split := 0
			for i := t + 1; i < len(directions); i++ {
				if directions[i] == '(' {
					count++
					continue
				}
				if split == 0 && directions[i] == '|' {
					split = i - (t + 1)
					continue
				}
				if directions[i] == ')' {
					count--
					if count == 0 {
						branch = directions[t+1 : i]
						t = i + 1
						break
					}
				}
			}
			oldX, oldY := p.cX, p.cY
			x, y := p.Map(branch[:split])
			oX += x
			oY += y
			p.cX, p.cY = oldX+x, oldY+y
			x, y = p.Map(branch[split+1:])
			oX += x
			oY += y
			p.cX, p.cY = oldX+x, oldY+y
			fmt.Printf("continuing with %s at %d,%d index %d\n", string(directions[t:]), p.cX, p.cY, t)
		case ')', '|':
			break
		default:
			switch c {
			case 'N':
				x, y := p.walk(DoorNorth)
				oX += x
				oY += y
			case 'W':
				x, y := p.walk(DoorWest)
				oX += x
				oY += y
			case 'E':
				x, y := p.walk(DoorEast)
				oX += x
				oY += y
			case 'S':
				x, y := p.walk(DoorSouth)
				oX += x
				oY += y
			}
		}
	}
	return oX, oY
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
	//p.Map([]byte("^WNE$"))
	//p.Map([]byte("^WNEE(SSSW|EEN(ES|EN))NN(WW|)NN$"))
	p.Map([]byte("^ESSWWN(E|NNENN(EESS(WNSE|)SSS|WWWSSSSE(SW|NNNE)))$"))
	p.PrintMap()
	waitEnter()

	p = NewPuzzle()
	p.MapFile("input.txt")
}
