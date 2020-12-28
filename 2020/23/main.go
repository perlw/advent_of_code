package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

// CircElem ...
type CircElem struct {
	Next *CircElem
	Val  int
}

// CircList ...
type CircList struct {
	Root     *CircElem
	Last     *CircElem
	Count    int
	ValElems map[int]*CircElem
}

// Push ...
func (c *CircList) Push(a int) *CircElem {
	cl := &CircElem{}
	if c.Root == nil {
		cl.Next = cl
		cl.Val = a
		c.Root = cl
		c.Last = cl
		c.ValElems = make(map[int]*CircElem)
	} else {
		cl.Next = c.Root
		cl.Val = a
		c.Last.Next = cl
		c.Last = cl
	}
	c.Count++
	c.ValElems[a] = cl
	return cl
}

// Find ...
func (c *CircList) Find(a int) *CircElem {
	var i int
	for e := c.Root; e != nil; e = e.Next {
		if e.Val == a {
			return e
		}

		i++
		if i >= c.Count {
			break
		}
	}
	return nil
}

// Game ...
type Game struct {
	Cups    *CircList
	Current *CircElem
	Move    int
	Lowest  int
	Highest int
}

// Turn ...
func (g *Game) Turn(debug bool) {
	target := g.Current.Val - 1
	g.Move++

	if debug {
		fmt.Printf("\n-- move %d --\n", g.Move)
		fmt.Printf("%s\n", g)
	}

	e := g.Current
	e1 := e.Next
	e2 := e1.Next
	e3 := e2.Next
	e.Next = e3.Next
	hand := []int{e1.Val, e2.Val, e3.Val}
	if debug {
		fmt.Printf("pick up: %+v\n", hand)
	}

	if target < g.Lowest {
		target = g.Highest
	}
	for {
		var found bool
		for _, d := range hand {
			if target == d {
				found = true
				target--
				if target < g.Lowest {
					target = g.Highest
				}
			}
		}
		if !found {
			break
		}
	}
	if debug {
		fmt.Printf("destination: %d\n", target)
	}

	targetE := g.Cups.ValElems[target]
	e3.Next = targetE.Next
	targetE.Next = e1

	g.Current = g.Current.Next
}

func (g Game) String() string {
	var result string
	result += fmt.Sprintf("cups: ")
	var i int
	for e := g.Cups.Root; e != nil; e = e.Next {
		if e.Val == g.Current.Val {
			result += fmt.Sprintf("(%d)", e.Val)
		} else {
			result += fmt.Sprintf(" %d ", e.Val)
		}
		i++
		if i >= g.Cups.Count {
			break
		}
	}
	return result
}

// Task1 ...
func Task1(input Game, iterations int, debug bool) string {
	for i := 0; i < iterations; i++ {
		input.Turn(debug)
	}

	if debug {
		fmt.Println("\n-- final--")
		fmt.Printf("%s\n", input)
	}

	var result string
	for e := input.Cups.ValElems[1].Next; e.Val != 1; e = e.Next {
		result += strconv.Itoa(e.Val)
	}
	return result
}

// Task2 ...
func Task2(input Game, debug bool) int {
	for i := input.Highest + 1; i <= 1000000; i++ {
		input.Cups.Push(i)
	}
	input.Highest = 1000000

	for i := 0; i < 10000000; i++ {
		input.Turn(debug)
	}

	e := input.Cups.ValElems[1]
	return e.Next.Val * e.Next.Next.Val
}

func readInput(reader io.Reader) Game {
	input := Game{
		Cups:   &CircList{},
		Lowest: 1,
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		for _, b := range scanner.Bytes() {
			d := int(b - '0')
			if d > input.Highest {
				input.Highest = d
			}
			input.Cups.Push(d)
		}
	}

	input.Current = input.Cups.Root

	return input
}

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()

	var input1, input2 Game
	file, _ := os.Open("input.txt")
	input1 = readInput(file)
	file.Close()
	file, _ = os.Open("input.txt")
	input2 = readInput(file)
	file.Close()

	result1 := Task1(input1, 100, debug)
	fmt.Printf("Task 1: %s\n", result1)

	result2 := Task2(input2, debug)
	fmt.Printf("Task 2: %d\n", result2)
}
