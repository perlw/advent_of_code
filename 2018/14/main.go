package main

import (
	"fmt"
	"time"
)

type Puzzle struct {
	Recipes []int
	Elves   []int
}

func wrap(v, max int) int {
	for v >= max {
		v -= max
	}
	return v
}

func (p *Puzzle) Solution1(input int, pretty bool) []int {
	for {
		value := 0
		for u := range p.Elves {
			value += p.Recipes[p.Elves[u]]
		}
		if value >= 10 {
			p.Recipes = append(p.Recipes, value/10, value%10)
		} else {
			p.Recipes = append(p.Recipes, value)
		}
		for u := range p.Elves {
			v := p.Recipes[p.Elves[u]]
			p.Elves[u] = wrap(p.Elves[u]+1+v, len(p.Recipes))
		}

		if pretty {
			for t := range p.Recipes {
				found := false
				for u := range p.Elves {
					if p.Elves[u] == t {
						fmt.Printf("%c%d%c", 33+u, p.Recipes[t], 33+u)
						found = true
					}
				}
				if !found {
					fmt.Printf(" %d ", p.Recipes[t])
				}
			}
			fmt.Printf("\n")
			time.Sleep(100 * time.Millisecond)
		}

		if len(p.Recipes) >= input+10 {
			break
		}
	}
	return p.Recipes[input : input+10]
}

func (p *Puzzle) Solution2(input []int, pretty bool) int {
	for {
		value := 0
		for u := range p.Elves {
			value += p.Recipes[p.Elves[u]]
		}
		if value >= 10 {
			p.Recipes = append(p.Recipes, value/10, value%10)
		} else {
			p.Recipes = append(p.Recipes, value)
		}
		for u := range p.Elves {
			v := p.Recipes[p.Elves[u]]
			p.Elves[u] = wrap(p.Elves[u]+1+v, len(p.Recipes))
		}

		if pretty {
			for t := range p.Recipes {
				found := false
				for u := range p.Elves {
					if p.Elves[u] == t {
						fmt.Printf("%c%d%c", 33+u, p.Recipes[t], 33+u)
						found = true
					}
				}
				if !found {
					fmt.Printf(" %d ", p.Recipes[t])
				}
			}
			fmt.Printf("\n")
			time.Sleep(100 * time.Millisecond)
		}

		if len(p.Recipes) < len(input)+1 {
			continue
		}

		found := true
		part := p.Recipes[len(p.Recipes)-len(input)-1:]
		for i := range input {
			if part[i] != input[i] {
				found = false
				break
			}
		}
		if found {
			return len(p.Recipes) - len(input) - 1
		}
	}
}

func main() {
	p := Puzzle{
		Recipes: []int{3, 7},
		Elves:   []int{0, 1},
	}
	/* fmt.Println(p.Solution1(5, true))

	p = Puzzle{
		Recipes: []int{3, 7},
		Elves:   []int{0, 1},
	}
	fmt.Println(p.Solution2([]int{5, 1, 5, 8, 9}, true))

	p = Puzzle{
		Recipes: []int{3, 7},
		Elves:   []int{0, 1},
	}
	fmt.Println(p.Solution2([]int{0, 1, 2, 4, 5}, true))

	p = Puzzle{
		Recipes: []int{3, 7},
		Elves:   []int{0, 1},
	}
	fmt.Println(p.Solution2([]int{9, 2, 5, 1, 0}, true))

	p = Puzzle{
		Recipes: []int{3, 7},
		Elves:   []int{0, 1},
	}
	fmt.Println(p.Solution2([]int{5, 9, 4, 1, 4}, false))

	p = Puzzle{
		Recipes: []int{3, 7},
		Elves:   []int{0, 1},
	}*/
	fmt.Println(p.Solution1(633601, false))

	p = Puzzle{
		Recipes: []int{3, 7},
		Elves:   []int{0, 1},
	}
	fmt.Println(p.Solution2([]int{6, 3, 3, 6, 0, 1}, false))
}
