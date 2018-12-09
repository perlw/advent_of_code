package main

import (
	"fmt"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Item struct {
	Prev *Item
	Next *Item
	Val  int
}

type Circular struct {
	Root    *Item
	Current *Item
}

func NewCircular() *Circular {
	i := Item{
		Val: 0,
	}
	i.Prev = &i
	i.Next = &i
	return &Circular{
		Current: &i,
		Root:    &i,
	}
}

func (c *Circular) Add(i int) {
	it := Item{
		Prev: c.Current,
		Next: c.Current.Next,
		Val:  i,
	}
	c.Current.Next.Prev = &it
	c.Current.Next = &it
	c.Current = &it
}

func (c *Circular) Pop() int {
	prev := c.Current.Prev
	next := c.Current.Next
	prev.Next = next
	next.Prev = prev
	val := c.Current.Val
	c.Current = next
	return val
}

func (c *Circular) Rotate(i int) {
	for {
		if i > 0 {
			c.Current = c.Current.Next
			i--
		} else if i < 0 {
			c.Current = c.Current.Prev
			i++
		} else {
			break
		}
	}
}

func (c Circular) String() string {
	output := strings.Builder{}
	it := c.Root
	for {
		if it == c.Current {
			output.WriteString(fmt.Sprintf("\t(%3d)", it.Val))
		} else {
			output.WriteString(fmt.Sprintf("\t%4d", it.Val))
		}
		it = it.Next
		if it == c.Root {
			break
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

	if debug {
		fmt.Printf("[--]%s\n", p.Marbles)
	}
	pnum := 0
	for t := 1; t <= p.LastMarble; t++ {
		if t%23 == 0 {
			p.Marbles.Rotate(-7)
			players[pnum] += t + p.Marbles.Pop()
		} else {
			p.Marbles.Rotate(1)
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
		//Players:    9,
		// LastMarble: 25,
		Players:    439,
		LastMarble: 7130700,
		Marbles:    NewCircular(),
	}

	fmt.Println(p.Solution1(false))
	fmt.Println(p.Solution2())
}
