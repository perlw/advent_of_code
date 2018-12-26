package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func waitEnter() {
	fmt.Printf("Press ENTER to continue\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for t := range a {
		if a[t] != b[t] {
			return false
		}
	}
	return true
}

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

func manhattan(a, b *NanoBot) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y) + abs(a.Z-b.Z)
}

type NanoBot struct {
	X, Y, Z int
	R       int
}

type Puzzle struct {
	bots []NanoBot
}

func NewPuzzle(filepath string) *Puzzle {
	p := Puzzle{
		bots: make([]NanoBot, 0, 400),
	}

	input, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	scanner := bufio.NewScanner(input)

	for scanner.Scan() {
		var x, y, z, r int
		fmt.Sscanf(scanner.Text(), "pos=<%d,%d,%d>, r=%d", &x, &y, &z, &r)

		p.bots = append(p.bots, NanoBot{
			X: x,
			Y: y,
			Z: z,
			R: r,
		})
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	input.Close()

	return &p
}

func (p *Puzzle) FindStrongestBot() *NanoBot {
	var bot *NanoBot
	largest := 0
	for t, b := range p.bots {
		if b.R > largest {
			largest = b.R
			bot = &p.bots[t]
		}
	}
	return bot
}

func (p *Puzzle) DetectInRangeBot(bot *NanoBot) []*NanoBot {
	bots := make([]*NanoBot, 0, 400)
	for t, b := range p.bots {
		if manhattan(bot, &b) <= bot.R {
			bots = append(bots, &p.bots[t])
		}
	}
	return bots
}

func main() {
	// Test
	t := NewPuzzle("test.txt")
	bot := t.FindStrongestBot()
	spew.Dump(bot)
	inRange := t.DetectInRangeBot(bot)
	spew.Dump(inRange)
	fmt.Println("In range:", len(inRange))

	// Puzzle
	p := NewPuzzle("input.txt")
	fmt.Println("In range:", len(p.DetectInRangeBot(p.FindStrongestBot())))
}
