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

var doors = map[byte]Point{
	'N': Point{0, -1},
	'W': Point{-1, 0},
	'E': Point{1, 0},
	'S': Point{0, 1},
}

type Puzzle struct {
}

func (p *Puzzle) PrintMap() {
	/*	for y := 0; y < p.gridH; y++ {
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
	*/
}

type Point struct {
	X, Y int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func (p *Puzzle) Map(directions []byte) {
	positions := make([]Point, 0, 100)
	distances := make(map[Point]int)

	pos := Point{5000, 5000}
	prev := pos
	for _, c := range directions[1 : len(directions)-1] {
		switch c {
		case '(':
			positions = append(positions, pos)
		case ')':
			pos, positions = positions[len(positions)-1], positions[:len(positions)-1]
		case '|':
			pos = positions[len(positions)-1]
		default:
			target := doors[c]
			pos.X += target.X
			pos.Y += target.Y
			if _, ok := distances[pos]; ok {
				distances[pos] = min(distances[pos], distances[prev]+1)
			} else {
				distances[pos] = distances[prev] + 1
			}
		}
		prev = pos
	}

	dist := 0
	count := 0
	for _, d := range distances {
		if d > dist {
			dist = d
		}
		if d >= 1000 {
			count++
		}
	}
	fmt.Println("Solution 1:", dist, "steps")
	fmt.Println("Solution 2:", count, "rooms")
}

func (p *Puzzle) MapFile(filepath string) {
	var buf []byte
	input, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		buf = scanner.Bytes()
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	p.Map(buf)
}

func main() {
	// Tests
	p := Puzzle{}
	p.Map([]byte("^WNE$"))
	p.Map([]byte("^ENWWW(NEEE|SSE(EE|N))$"))
	p.Map([]byte("^ENNWSWW(NEWS|)SSSEEN(WNSE|)EE(SWEN|)NNN$"))
	p.MapFile("input.txt")
	waitEnter()
}
