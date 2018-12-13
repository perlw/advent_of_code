package main

import (
	"bufio"
	"fmt"
	"os"
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
}

type Puzzle struct {
	Playfield []byte
	Carts     []Cart
}

func (p *Puzzle) FindCollision(x, y int) bool {
	count := 0
	for t := range p.Carts {
		if p.Carts[t].X == x && p.Carts[t].Y == y {
			count++
		}
	}
	return count > 1
}

func (p *Puzzle) FindCart(x, y int) int {
	for t := range p.Carts {
		if p.Carts[t].X == x && p.Carts[t].Y == y {
			return t
		}
	}
	return -1
}

func (p *Puzzle) Tick() (bool, int, int) {
	crash := false
	var cX, cY int
	carts := make([]Cart, 0, len(p.Carts))
	for y := 0; y < gridY; y++ {
		for x := 0; x < gridX; x++ {
			i := (y * gridX) + x

			if ci := p.FindCart(x, y); ci > -1 {
				c := p.Playfield[i]
				cart := p.Carts[ci]

				switch c {
				case '\\':
					switch cart.Dir {
					case '<':
						cart.Dir = '^'
					case '>':
						cart.Dir = 'v'
					case '^':
						cart.Dir = '<'
					case 'v':
						cart.Dir = '>'
					}
				case '/':
					switch cart.Dir {
					case '<':
						cart.Dir = 'v'
					case '>':
						cart.Dir = '^'
					case '^':
						cart.Dir = '>'
					case 'v':
						cart.Dir = '<'
					}
				case '+':
					switch cart.Dir {
					case '<':
						switch cart.LastTurn {
						case 0:
							cart.Dir = 'v'
							cart.LastTurn++
						case 1:
							cart.LastTurn++
						case 2:
							cart.Dir = '^'
							cart.LastTurn = 0
						}
					case '>':
						switch cart.LastTurn {
						case 0:
							cart.Dir = '^'
							cart.LastTurn++
						case 1:
							cart.LastTurn++
						case 2:
							cart.Dir = 'v'
							cart.LastTurn = 0
						}
					case '^':
						switch cart.LastTurn {
						case 0:
							cart.Dir = '<'
							cart.LastTurn++
						case 1:
							cart.LastTurn++
						case 2:
							cart.Dir = '>'
							cart.LastTurn = 0
						}
					case 'v':
						switch cart.LastTurn {
						case 0:
							cart.Dir = '>'
							cart.LastTurn++
						case 1:
							cart.LastTurn++
						case 2:
							cart.Dir = '<'
							cart.LastTurn = 0
						}
					}
				}

				switch cart.Dir {
				case '<':
					cart.X--
				case '>':
					cart.X++
				case '^':
					cart.Y--
				case 'v':
					cart.Y++
				}

				carts = append(carts, cart)
			}

			if p.FindCollision(x, y) {
				crash = true
				cX = x
				cY = y
				break
			}
		}
	}
	p.Carts = carts
	return crash, cX, cY
}

func (p *Puzzle) Solution1(pretty bool) (int, int) {
	var cX, cY int
	for {
		crash := false
		if crash, cX, cY = p.Tick(); crash {
			break
		}
		if pretty {
			for y := 0; y < gridY; y++ {
				for x := 0; x < gridX; x++ {
					i := (y * gridX) + x
					c := p.Playfield[i]

					if cart := p.FindCart(x, y); cart > -1 {
						c = p.Carts[cart].Dir
					}

					fmt.Printf("%c", c)
				}
				fmt.Printf("\n")
			}
		}

		time.Sleep(100 * time.Millisecond)
	}
	for y := 0; y < gridY; y++ {
		for x := 0; x < gridX; x++ {
			i := (y * gridX) + x
			c := p.Playfield[i]

			if x == cX && y == cY {
				c = 'X'
			}

			fmt.Printf("%c", c)
		}
		fmt.Printf("\n")
	}
	return cX, cY
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
