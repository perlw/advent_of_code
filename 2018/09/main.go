package main

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Circular struct {
	Data  []int
	Index int
}

func NewCircular() *Circular {
	return &Circular{
		Data:  make([]int, 0, 100),
		Index: 0,
	}
}

func (c *Circular) Add(i int) {
	if len(c.Data) == c.Index {
		c.Data = append(c.Data, i)
	} else {
		tmp := append([]int{i}, c.Data[c.Index:]...)
		c.Data = append(c.Data[:c.Index], tmp...)
	}
	c.Index++
}

func (c *Circular) Del() int {
	t := c.Data[c.Index]
	c.Data = append(c.Data[:c.Index], c.Data[c.Index+1:]...)
	c.Move(1)
	return t
}

func (c *Circular) Move(i int) {
	c.Index += i
	if c.Index >= len(c.Data) {
		c.Index -= len(c.Data)
	} else if c.Index < 0 {
		c.Index = len(c.Data) + c.Index
	}
}

func (c Circular) String() string {
	output := strings.Builder{}
	for t, m := range c.Data {
		if t == c.Index-1 {
			output.WriteString(fmt.Sprintf("\t(%3d)", m))
		} else {
			output.WriteString(fmt.Sprintf("\t%4d", m))
		}
	}
	return output.String()
}

type Puzzle struct {
	Players    int
	LastMarble int
	Marbles    *Circular
}

func (p *Puzzle) Solution1(debug bool) int {
	players := make([]int, p.Players)

	p.Marbles.Add(0)
	if debug {
		fmt.Printf("[--]%s\n", p.Marbles)
	}
	pnum := 0
	for t := 1; t <= p.LastMarble; t++ {
		if t%23 == 0 {
			p.Marbles.Move(-8)
			i := p.Marbles.Del()
			players[pnum] += t + i
		} else {
			p.Marbles.Move(1)
			p.Marbles.Add(t)
		}
		if debug {
			fmt.Printf("[%2d]%s\n", pnum+1, p.Marbles)
		}
		pnum++
		if pnum >= p.Players {
			pnum = 0
		}
	}

	if debug {
		spew.Dump(players)
	}

	highest := 0
	for _, p := range players {
		if p > highest {
			highest = p
		}
	}
	return highest
}

func (p *Puzzle) Solution2() int {
	return 0
}

func main() {
	p := Puzzle{
		Players:    439,
		LastMarble: 71307,
		Marbles:    NewCircular(),
	}

	fmt.Println(p.Solution1(false))
	fmt.Println(p.Solution2())
}
