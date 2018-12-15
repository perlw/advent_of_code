package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"time"
)

const (
	gridX = 150
	gridY = 150
)

var gridSize = gridX * gridY

type Cart struct {
	X, Y     int
	Dir      byte
	LastTurn int
	Dead     bool
}

func (c *Cart) Turn() {
	switch c.LastTurn {
	case 0:
		switch c.Dir {
		case '>':
			c.Dir = '^'
		case '<':
			c.Dir = 'v'
		case '^':
			c.Dir = '<'
		case 'v':
			c.Dir = '>'
		}
		c.LastTurn++
	case 1:
		c.LastTurn++
	case 2:
		switch c.Dir {
		case '>':
			c.Dir = 'v'
		case '<':
			c.Dir = '^'
		case '^':
			c.Dir = '>'
		case 'v':
			c.Dir = '<'
		}
		c.LastTurn = 0
	}
}

func (c *Cart) Move(b byte) {
	switch b {
	case '\\':
		switch c.Dir {
		case '>':
			c.Dir = 'v'
		case '<':
			c.Dir = '^'
		case '^':
			c.Dir = '<'
		case 'v':
			c.Dir = '>'
		}
	case '/':
		switch c.Dir {
		case '>':
			c.Dir = '^'
		case '<':
			c.Dir = 'v'
		case '^':
			c.Dir = '>'
		case 'v':
			c.Dir = '<'
		}
	case '+':
		c.Turn()
	}

	switch c.Dir {
	case '>':
		c.X++
	case '<':
		c.X--
	case '^':
		c.Y--
	case 'v':
		c.Y++
	}
}

type ByCoord []Cart

func (b ByCoord) Len() int {
	return len(b)
}
func (b ByCoord) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}
func (b ByCoord) Less(i, j int) bool {
	ia := (b[i].Y * gridX) + b[i].X
	ib := (b[j].Y * gridX) + b[j].X
	return ia < ib
}

type Collision struct {
	Ca, Cb int
	X, Y   int
}

type Puzzle struct {
	Playfield []byte
	Carts     []Cart
}

func (p *Puzzle) Tick() []Collision {
	sort.Sort(ByCoord(p.Carts))
	collisions := make([]Collision, 0, 2)
	for t, c := range p.Carts {
		if c.Dead {
			continue
		}

		b := p.Playfield[(c.Y*gridX)+c.X]
		p.Carts[t].Move(b)

		for u, cb := range p.Carts {
			if cb.Dead || t == u {
				continue
			}

			if p.Carts[t].X == cb.X && p.Carts[t].Y == cb.Y {
				p.Carts[t].Dead = true
				p.Carts[u].Dead = true
				collisions = append(collisions, Collision{
					Ca: t,
					Cb: u,
					X:  p.Carts[t].X,
					Y:  p.Carts[t].Y,
				})
				break
			}
		}
	}
	return collisions
}

func (p *Puzzle) Solution1(pretty bool) (int, int) {
	grid := make([]byte, gridSize)

	collisions := make([]Collision, 0, 10)
	for {
		count := 0
		for _, c := range p.Carts {
			if c.Dead {
				count++
			}
		}
		if count >= len(p.Carts)-1 {
			break
		}

		c := p.Tick()
		if len(c) > 0 {
			collisions = append(collisions, c...)
		}

		if pretty {
			copy(grid, p.Playfield)
			for _, c := range p.Carts {
				if c.Dead {
					continue
				}

				i := (c.Y * gridX) + c.X
				grid[i] = c.Dir
			}

			for y := 0; y < gridY; y++ {
				fmt.Println(string(grid[y*gridX : (y*gridX)+gridX]))
			}

			time.Sleep(100 * time.Millisecond)
		}
	}
	copy(grid, p.Playfield)
	for _, c := range collisions {
		i := (c.Y * gridX) + c.X
		grid[i] = 'X'
	}
	for _, c := range p.Carts {
		if c.Dead {
			continue
		}

		i := (c.Y * gridX) + c.X
		grid[i] = c.Dir
	}

	for y := 0; y < gridY; y++ {
		fmt.Println(string(grid[y*gridX : (y*gridX)+gridX]))
	}

	var aX, aY int
	for _, c := range p.Carts {
		if c.Dead {
			fmt.Println("dead", c.X, c.Y)
			continue
		}
		fmt.Println("alive", c.X, c.Y)
		aX = c.X
		aY = c.Y
	}
	return aX, aY
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	p := Puzzle{
		Playfield: make([]byte, gridSize),
		Carts:     make([]Cart, 0, 2),
	}
	for t := range p.Playfield {
		p.Playfield[t] = ' '
	}
	lines := 0
	for scanner.Scan() {
		line := []byte(scanner.Text())
		for x, c := range line {
			i := (lines * gridX) + x

			if c == '>' || c == 'v' || c == '<' || c == '^' {
				p.Carts = append(p.Carts, Cart{
					X:        x,
					Y:        lines,
					Dir:      c,
					LastTurn: 0,
				})
				// Seems there's never a train started on an intersection
				if c == '>' || c == '<' {
					c = '-'
				} else {
					c = '|'
				}
			}

			p.Playfield[i] = c
		}
		lines++
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	/*
		palette := []color.Color{
			color.RGBA{0, 0, 0, 0xff},
			color.RGBA{0xff, 0xff, 0xff, 0xff},
		}
		size := image.Rect(0, 0, 128, 128)
		images := make([]*image.Paletted, 10)
		delays := make([]int, 10)
		for t := 0; t < 10; t++ {
			images[t] = image.NewPaletted(size, palette)

			for y := 0; y < 128; y++ {
				for x := 0; x < 128; x++ {
					c := byte((x + t) ^ (y + t))
					images[t].Set(x, y, color.RGBA{
						c, c, c, 0xff,
					})
				}
			}
		}

		f, err := os.OpenFile("out.gif", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		gif.EncodeAll(f, &gif.GIF{
			Image: images,
			Delay: delays,
		})
	*/

	fmt.Println(p.Solution1(false))
}
