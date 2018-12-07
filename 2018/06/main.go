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
	MaxX, MaxY int
}

func (p *Puzzle) Add(x, y int) {
	p.Points = append(p.Points, Point{
		ID: byte('"' + len(p.Points)),
		X:  x,
		Y:  y,
	})
	if x > p.MaxX {
		p.MaxX = x
	}
	if y > p.MaxY {
		p.MaxY = y
	}
}

func (p *Puzzle) Solution1() int {
	size := (p.MaxY + 2) * (p.MaxX + 2)
	grid := make([]byte, size)
	for y := 0; y < p.MaxY+2; y++ {
		for x := 0; x < p.MaxX+2; x++ {
			i := (y * (p.MaxX + 2)) + x
			c := byte('.')

			if x == 0 || x == p.MaxX+1 || y == 0 || y == p.MaxY+1 {
				c = 219
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
	for y := 0; y < p.MaxY+2; y++ {
		for x := 0; x < p.MaxX+2; x++ {
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
					id = p.Points[t].ID
				}
			}
			matches := 0
			for _, d := range dist {
				if d == shortest {
					matches++
				}
			}

			i := (y * (p.MaxX + 2)) + x
			if grid[i] == 219 {
				for t := range p.Points {
					if p.Points[t].ID == id {
						p.Points[t].Inf = true
						break
					}
				}
			} else {
				if matches > 1 {
					id = '!'
				}
				grid[i] = id
			}
		}
	}

	/*
		for y := 0; y < p.MaxY+2; y++ {
			for x := 0; x < p.MaxX+2; x++ {
				i := (y * (p.MaxX + 2)) + x
				fmt.Printf("%c", grid[i])
			}
			fmt.Printf("\n")
		}
	*/

	largest := 0
	for _, point := range p.Points {
		if point.Inf {
			continue
		}

		count := 0
		for _, c := range grid {
			if c == point.ID {
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
	size := (p.MaxY + 2) * (p.MaxX + 2)
	grid := make([]byte, size)
	for y := 0; y < p.MaxY+2; y++ {
		for x := 0; x < p.MaxX+2; x++ {
			i := (y * (p.MaxX + 2)) + x
			c := byte('.')

			if x == 0 || x == p.MaxX+1 || y == 0 || y == p.MaxY+1 {
				c = 219
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

	drawn := 0
	for y := 0; y < p.MaxY+2; y++ {
		for x := 0; x < p.MaxX+2; x++ {
			total := 0
			for _, point := range p.Points {
				total += point.Distance(x, y)
			}

			i := (y * (p.MaxX + 2)) + x
			if grid[i] == 219 {
				continue
			} else {
				if total < 10000 {
					grid[i] = '#'
					drawn++
				}
			}
		}
	}

	/*for y := 0; y < p.MaxY+2; y++ {
		for x := 0; x < p.MaxX+2; x++ {
			i := (y * (p.MaxX + 2)) + x
			fmt.Printf("%c", grid[i])
		}
		fmt.Printf("\n")
	}*/

	return drawn
}

func (p Puzzle) String() string {
	output := strings.Builder{}
	output.WriteString(fmt.Sprintf("%dx%d\n", p.MaxX, p.MaxY))
	for y := 0; y < p.MaxY+2; y++ {
		for x := 0; x < p.MaxX+2; x++ {
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
