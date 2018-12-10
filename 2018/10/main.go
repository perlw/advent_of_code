package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

func waitEnter() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

type Point struct {
	X, Y   int
	VX, VY int
}

type Puzzle struct {
	Points []*Point
}

func (p *Puzzle) Solution1() {
	var minX, minY int
	var maxX, maxY int
	for _, p := range p.Points {
		if p.X < minX {
			minX = p.X
		}
		if p.Y < minY {
			minY = p.Y
		}
		if p.X > maxX {
			maxX = p.X
		}
		if p.Y > maxY {
			maxY = p.Y
		}
	}
	minX++
	minY++
	maxX++
	maxY++

	centerX := (maxX + minX) / 2
	centerY := (maxY + minY) / 2
	fmt.Println(minX, minY, maxX, maxY)
	fmt.Println(centerX, centerY)
	tick := 0
	for {
		for _, p := range p.Points {
			p.X += p.VX
			p.Y += p.VY
		}

		count := 0
		for _, p := range p.Points {
			if p.X > centerX-120 && p.X < centerX+120 && p.Y > centerY-20 && p.Y < centerY+20 {
				count++
			}
		}

		tick++
		if count >= len(p.Points) {
			for y := centerY - 20; y < centerY+20; y++ {
				for x := centerX - 120; x < centerX+120; x++ {
					found := false
					for _, p := range p.Points {
						if p.X == x && p.Y == y {
							found = true
							break
						}
					}
					if found {
						fmt.Printf("#")
					} else {
						fmt.Printf(".")
					}
				}
				fmt.Printf("\n")
			}
			fmt.Println(count, tick)
			waitEnter()
		}
	}
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	p := Puzzle{
		Points: make([]*Point, 0, 10),
	}
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), ">")
		pos := strings.Split(strings.Split(data[0], "<")[1], ", ")
		vel := strings.Split(strings.Split(data[1], "<")[1], ", ")

		x, _ := strconv.Atoi(strings.TrimSpace(pos[0]))
		y, _ := strconv.Atoi(strings.TrimSpace(pos[1]))
		vx, _ := strconv.Atoi(strings.TrimSpace(vel[0]))
		vy, _ := strconv.Atoi(strings.TrimSpace(vel[1]))

		p.Points = append(p.Points, &Point{
			X:  x,
			Y:  y,
			VX: vx,
			VY: vy,
		})
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	spew.Dump(p)

	p.Solution1()
}
