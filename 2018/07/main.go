package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

type Worker struct {
	Current byte
	Left    int
}

func (w *Worker) Assign(c byte) {
	w.Current = c
	w.Left = int(c-'A'+1) + 60
}

func (w *Worker) Tick() {
	w.Left--
}

func (w *Worker) Done() bool {
	return (w.Left == 0)
}

func (p *Puzzle) Solution2() int {
	workers := make([]Worker, 5)

	avail := make([]byte, 0, 10)
	visited := make([]byte, 0, 10)
	done := make([]byte, 0, 10)

	for _, step := range p.Steps {
		if contains(visited, step.Requires) {
			fmt.Println(string(step.ID), "->", string(step.Requires), "/", string(step.RequiredBy))
			avail = append(avail, step.ID)
		}
	}

	count := 0
	fmt.Printf("Second")
	for t := range workers {
		workers[t].Current = '.'
		fmt.Printf("\tWorker%d", t)
	}
	fmt.Printf("\tAvail")
	fmt.Println("\tDone")
	for {
		for i := range workers {
			if workers[i].Done() {
				if workers[i].Current != 0 && workers[i].Current != '.' {
					if !contains(done, []byte{workers[i].Current}) {
						done = append(done, workers[i].Current)
					}
					for _, step := range p.Steps {
						if contains(done, step.Requires) && !contains(visited, []byte{step.ID}) &&
							!contains(avail, []byte{step.ID}) && !contains(done, []byte{step.ID}) {
							avail = append(avail, step.ID)
						}
					}
					workers[i].Current = '.'
				}

				ni := -1
				next := byte('[')
				for t, c := range avail {
					if c < next {
						next = c
						ni = t
					}
				}
				if next == '[' {
					continue
				}
				avail = append(avail[:ni], avail[ni+1:]...)

				if !contains(visited, []byte{next}) {
					visited = append(visited, next)
				}

				workers[i].Assign(next)
			}

			workers[i].Tick()
		}

		fmt.Printf("%4d", count)
		for _, worker := range workers {
			fmt.Printf("\t%4c", worker.Current)
		}
		fmt.Printf("\t%s", avail)
		fmt.Printf("\t%s\n", done)
		if len(done) == len(p.Steps) {
			break
		}
		count++
	}

	return count
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

	fmt.Println(p.Solution1())
	fmt.Println("-->", p.Solution2(), "<--")
}
