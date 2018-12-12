package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var lineSize = 2000
var offset = 10

type Rule struct {
	Match []byte
	Out   byte
}

type Puzzle struct {
	Plants []byte
	Rules  []Rule
}

func sliceEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for t := 0; t < len(a); t++ {
		if a[t] != b[t] {
			return false
		}
	}
	return true
}

func (p *Puzzle) findMatch(pattern []byte) byte {
	for _, r := range p.Rules {
		if sliceEqual(r.Match, pattern) {
			return r.Out
		}
	}
	return '.'
}

func (p *Puzzle) Solution1() int {
	fmt.Printf("%2d %s\n", 0, string(p.Plants))
	for i := 0; i < 20; i++ {
		out := make([]byte, lineSize)
		copy(out, p.Plants)
		for t := 0; t < lineSize-offset; t++ {
			pattern := p.Plants[t : t+5]
			out[t+2] = p.findMatch(pattern)
		}
		p.Plants = out

		fmt.Printf("%2d %s\n", i+1, string(p.Plants))
	}
	points := 0
	for t, c := range p.Plants {
		if c == '#' {
			points += (t - offset)
		}
	}
	return points
}

func (p *Puzzle) Solution2() int {
	points := 0
	lastPoints := 0
	for i := 0; i < 200; i++ {
		fmt.Println(i, points, points-lastPoints)
		out := make([]byte, lineSize)
		for t := range p.Plants {
			out[t] = '.'
		}
		for t := 0; t < lineSize-offset; t++ {
			pattern := p.Plants[t : t+5]
			out[t+2] = p.findMatch(pattern)
		}
		p.Plants = out
		fmt.Printf("%2d %s\n", i+1, string(p.Plants))

		lastPoints = points
		points = 0
		for t, c := range p.Plants {
			if c == '#' {
				points += (t - offset)

				if t > lineSize-20 {
					fmt.Println("oh fuck")
					return 0
				}
			}
		}
	}
	return points
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	p := Puzzle{
		Plants: make([]byte, lineSize),
		Rules:  make([]Rule, 0, 10),
	}
	scanner.Scan()
	initial := []byte(strings.Split(scanner.Text(), ": ")[1])
	for t := range p.Plants {
		p.Plants[t] = '.'
	}
	for t := range initial {
		p.Plants[t+offset] = initial[t]
	}
	scanner.Scan()
	for scanner.Scan() {
		ruleParts := strings.Split(scanner.Text(), " => ")
		p.Rules = append(p.Rules, Rule{
			Match: []byte(ruleParts[0]),
			Out:   []byte(ruleParts[1])[0],
		})
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

	// fmt.Println(p.Solution1())
	fmt.Println(p.Solution2())
}
