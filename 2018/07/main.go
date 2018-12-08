package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Step struct {
	ID         byte
	Requires   []byte
	RequiredBy []byte
}

type Puzzle struct {
	Steps []Step
}

func findStep(steps []Step, ID byte) int {
	for t := range steps {
		if steps[t].ID == ID {
			return t
		}
	}
	return -1
}

func findFirst(steps []Step) int {
	for t := range steps {
		if len(steps[t].Requires) == 0 {
			return t
		}
	}
	return -1
}

func findLast(steps []Step) int {
	for t := range steps {
		if len(steps[t].RequiredBy) == 0 {
			return t
		}
	}
	return -1
}

// a is a superset of b
func contains(a []byte, b []byte) bool {
	if len(b) == 0 {
		return true
	}
	for _, bc := range b {
		found := false
		for _, ac := range a {
			if ac == bc {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

func (p *Puzzle) Add(ID byte, requires byte) {
	t := findStep(p.Steps, ID)
	if t == -1 {
		p.Steps = append(p.Steps, Step{
			ID:         ID,
			Requires:   make([]byte, 0, 10),
			RequiredBy: make([]byte, 0, 10),
		})
		t = len(p.Steps) - 1
	}
	p.Steps[t].RequiredBy = append(p.Steps[t].RequiredBy, requires)

	t = findStep(p.Steps, requires)
	if t == -1 {
		p.Steps = append(p.Steps, Step{
			ID:       requires,
			Requires: make([]byte, 0, 10),
		})
		t = len(p.Steps) - 1
	}
	p.Steps[t].Requires = append(p.Steps[t].Requires, ID)
}

func (p *Puzzle) Solution1() string {
	avail := make([]byte, 0, 10)
	visited := make([]byte, 0, 10)

	for _, step := range p.Steps {
		if contains(visited, step.Requires) {
			fmt.Println(string(step.ID), "->", string(step.Requires), "/", string(step.RequiredBy))
			avail = append(avail, step.ID)
		}
	}

	for {
		if len(avail) == 0 {
			break
		}

		next := byte('[')
		for _, c := range avail {
			if c < next {
				next = c
			}
		}
		alreadyVisited := false
		for _, c := range visited {
			if c == next {
				alreadyVisited = true
				break
			}
		}
		if alreadyVisited {
			break
		}
		visited = append(visited, next)
		curr := findStep(p.Steps, next)

		for _, c := range p.Steps[curr].RequiredBy {
			step := findStep(p.Steps, c)
			if contains(visited, p.Steps[step].Requires) && !contains(avail, []byte{c}) {
				avail = append(avail, c)
			}
		}

		fmt.Println("\ncurr", string(p.Steps[curr].ID))
		fmt.Println("avail", string(avail))
		fmt.Println("visited", string(visited))

		for t, c := range avail {
			if next == c {
				avail = append(avail[:t], avail[t+1:]...)
				break
			}
		}
	}

	return string(visited)
}

func (p *Puzzle) Solution2() int {
	return 0
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	p := Puzzle{
		Steps: make([]Step, 0, 10),
	}
	for scanner.Scan() {
		words := strings.Split(scanner.Text(), " ")
		p.Add(words[1][0], words[7][0])
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	spew.Dump(p.Steps)
	for _, step := range p.Steps {
		fmt.Println(string(step.ID), "->", string(step.Requires), "/", string(step.RequiredBy))
	}
	fmt.Println(p.Solution1())
	fmt.Println(p.Solution2())
}
