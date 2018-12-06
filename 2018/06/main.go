package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func abs(t int) int {
	if t < 0 {
		return -t
	}
	return t
}

type Point struct {
	ID  byte
	X   int
	Y   int
	Inf bool
}

func (p *Point) Equal(x, y int) bool {
	return (p.X == x && p.Y == y)
}

func (p *Point) Distance(x, y int) int {
	return abs(p.X-x) + abs(p.Y-y)
}

func (p Point) String() string {
	var inf string
	if p.Inf {
		inf = " INF"
	}
	return fmt.Sprintf("%c %dx%d%s", p.ID, p.X, p.Y, inf)
}

type Puzzle struct {
	Points     []Point
	MinX, MinY int
	MaxX, MaxY int
}

func (p *Puzzle) Add(x, y int) {
	p.Points = append(p.Points, Point{
		ID: byte('A' + len(p.Points)),
		X:  x,
		Y:  y,
	})
	if x < p.MinX {
		p.MinX = x
	}
	if y < p.MinY {
		p.MinY = y
	}
	if x > p.MaxX {
		p.MaxX = x
	}
	if y > p.MaxY {
		p.MaxY = y
	}
}

func (p *Puzzle) Solution1() int {
	size := ((p.MaxY + 2) - p.MinY) * ((p.MaxX + 2) - p.MinX)
	ox := -p.MinX
	oy := -p.MinY
	grid := make([]byte, size)
	for y := p.MinY; y < p.MaxY+2; y++ {
		for x := p.MinX; x < p.MaxX+2; x++ {
			i := ((y + oy) * ((p.MaxX + 2) + ox)) + (x + ox)
			c := byte('.')

			if x == p.MinX || x == p.MaxX+1 || y == p.MinY || y == p.MaxY+1 {
				c = '~'
			}

			for _, point := range p.Points {
				if point.Equal(x, y) {
					c = point.ID
					break
				}
			}
			grid[i] = c
		}
	}

	dist := make([]int, len(p.Points))
	for y := p.MinY; y < p.MaxY+2; y++ {
		for x := p.MinX; x < p.MaxX+2; x++ {
			skip := false
			for t, point := range p.Points {
				if point.Equal(x, y) {
					skip = true
					break
				}
				dist[t] = point.Distance(x, y)
			}
			if skip {
				continue
			}

			var id byte
			shortest := math.MaxInt32
			for t, d := range dist {
				if d < shortest {
					shortest = d
					id = p.Points[t].ID + 32
				}
			}
			matches := 0
			for _, d := range dist {
				if d == shortest {
					matches++
				}
			}

			i := ((y + oy) * ((p.MaxX + 2) + ox)) + (x + ox)
			if grid[i] == '~' {
				for t := range p.Points {
					if p.Points[t].ID == id-32 {
						p.Points[t].Inf = true
						break
					}
				}
			} else {
				if matches > 1 {
					id = '.'
				}
				grid[i] = id
			}
		}
	}

	for y := p.MinY; y < p.MaxY+2; y++ {
		for x := p.MinX; x < p.MaxX+2; x++ {
			i := ((y + oy) * ((p.MaxX + 2) + ox)) + (x + ox)
			fmt.Printf("%c", grid[i])
		}
		fmt.Printf("\n")
	}

	largest := 0
	for _, point := range p.Points {
		if point.Inf {
			continue
		}

		count := 0
		for _, c := range grid {
			if c == point.ID || c == point.ID+32 {
				count++
			}
		}
		if count > largest {
			largest = count
		}
	}

	return largest
}

func (p *Puzzle) Solution2() int {
	return 0
}

func (p Puzzle) String() string {
	output := strings.Builder{}
	output.WriteString(fmt.Sprintf("%dx%d -> %dx%d\n", p.MinX, p.MinY, p.MaxX, p.MaxY))
	for y := p.MinY; y < p.MaxY+2; y++ {
		for x := p.MinX; x < p.MaxX+2; x++ {
			c := byte('.')
			for _, point := range p.Points {
				if point.Equal(x, y) {
					c = point.ID
					break
				}
			}
			output.WriteByte(c)
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

	p := Puzzle{
		Points: make([]Point, 0, 1024),
	}
	for scanner.Scan() {
		var x, y int
		fmt.Sscanf(scanner.Text(), "%d,%d", &x, &y)
		p.Add(x, y)
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	fmt.Println(p.Solution1())
	fmt.Println(p.Solution2())
}
