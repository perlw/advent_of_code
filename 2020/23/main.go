package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Game ...
type Game struct {
	Cups    []int
	Current int
	Move    int
	Lowest  int
	Highest int
}

// Take ...
func (g *Game) Take(num int) []int {
	result := make([]int, 0, num)
	var start, end int

	start = g.Current + 1
	end = start + num
	if start+num < len(g.Cups) {
		result = append(result, g.Cups[start:end]...)
		g.Cups = append(g.Cups[:g.Current+1], g.Cups[end:]...)
	} else {
		end = (start + num) - len(g.Cups)
		result = append(result, g.Cups[start:]...)
		g.Cups = g.Cups[:g.Current+1]
		result = append(result, g.Cups[:end]...)
		g.Cups = g.Cups[end:]
	}
	return result
}

// Turn ...
func (g *Game) Turn(debug bool) {
	current := g.Cups[g.Current]
	target := g.Cups[g.Current] - 1
	g.Move++

	if debug {
		fmt.Printf("\n-- move %d --\n", g.Move)
		fmt.Printf("%s\n", g)
	}
	hand := g.Take(3)
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

	for i := range g.Cups {
		if g.Cups[i] == target {
			tmp := make([]int, len(g.Cups[i+1:]))
			copy(tmp, g.Cups[i+1:])
			g.Cups = append(g.Cups[:i+1], hand...)
			g.Cups = append(g.Cups, tmp...)
			break
		}
	}

	for i := range g.Cups {
		if g.Cups[i] == current {
			g.Current = (i + 1) % len(g.Cups)
			break
		}
	}
}

func (g Game) String() string {
	var result string
	fmt.Printf("cups: ")
	for i, d := range g.Cups {
		if i == g.Current {
			fmt.Printf("(%d)", d)
		} else {
			fmt.Printf(" %d ", d)
		}
	}
	return result
}

// Task1 ...
func Task1(input Game, iterations int, debug bool) string {
	for i := 0; i < iterations; i++ {
		input.Turn(debug)
	}

	fmt.Println("\n-- final--")
	fmt.Printf("%s\n", input)

	var result string
	var current int
	for i := range input.Cups {
		if input.Cups[i] == 1 {
			current = (i + 1) % len(input.Cups)
			break
		}
	}

	for input.Cups[current] != 1 {
		result += strconv.Itoa(input.Cups[current])
		current = (current + 1) % len(input.Cups)
	}

	return result
}

// Task2 ...
func Task2(input Game, debug bool) int {
	return -1
}

func readInput(reader io.Reader) Game {
	input := Game{
		Cups:   make([]int, 0, 10),
		Lowest: 10,
	}

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		for _, b := range scanner.Bytes() {
			d := int(b - '0')
			if d < input.Lowest {
				input.Lowest = d
			}
			if d > input.Highest {
				input.Highest = d
			}
			input.Cups = append(input.Cups, d)
		}
	}

	return input
}

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()

	var input Game
	file, _ := os.Open("input.txt")
	input = readInput(file)
	file.Close()

	result1 := Task1(input, 100, debug)
	fmt.Printf("Task 1: %s\n", result1)

	result2 := Task2(input, debug)
	fmt.Printf("Task 2: %d\n", result2)
}
