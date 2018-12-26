package main

import (
	"bufio"
	"container/heap"
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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func manhattan(a, b Point) int {
	return abs(a.X-b.X) + abs(a.Y-b.Y) + abs(a.Z-b.Z)
}

type Point struct {
	X, Y, Z int
}

type NanoBot struct {
	Pos Point
	R   int
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
			Pos: Point{
				X: x,
				Y: y,
				Z: z,
			},
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
		if manhattan(bot.Pos, b.Pos) <= bot.R {
			bots = append(bots, &p.bots[t])
		}
	}
	return bots
}

func (p *Puzzle) FindMostOverlapDist() int {
	pq := make(PriorityQueue, 0, len(p.bots)*2)

	for _, b := range p.bots {
		d := abs(b.Pos.X) + abs(b.Pos.Y) + abs(b.Pos.Z)
		pq = append(pq, &Item{
			value:    1,
			priority: max(0, d-b.R),
		})
		pq = append(pq, &Item{
			value:    -1,
			priority: d + b.R + 1,
		})
	}
	heap.Init(&pq)

	var result, count, maxCount int
	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		dist, e := item.priority, item.value
		count += e
		if count > maxCount {
			result = dist
			maxCount = count
		}
	}

	return result + 1
}

func main() {
	// Test1
	t := NewPuzzle("test1.txt")
	bot := t.FindStrongestBot()
	spew.Dump(bot)
	inRange := t.DetectInRangeBot(bot)
	spew.Dump(inRange)
	fmt.Println("In range:", len(inRange))

	// Test2
	t = NewPuzzle("test2.txt")
	bot = t.FindStrongestBot()
	spew.Dump(bot)
	fmt.Printf("Overlapping dist: %+v\n", t.FindMostOverlapDist())

	// Puzzle1
	p := NewPuzzle("input.txt")
	fmt.Println("In range:", len(p.DetectInRangeBot(p.FindStrongestBot())))
	fmt.Printf("Overlapping dist: %+v\n", p.FindMostOverlapDist())
}
