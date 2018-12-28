package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type Group struct {
	Amount     int
	HP         int
	Weak       []string
	Immune     []string
	Damage     int
	DamageType string
	Initiative int
}

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

type Puzzle struct {
	ImmuneSystem []Group
	Infection    []Group
}

func NewPuzzle(filepath string) *Puzzle {
	p := Puzzle{
		ImmuneSystem: make([]Group, 0, 10),
		Infection:    make([]Group, 0, 10),
	}

	input, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	scanner := bufio.NewScanner(input)

	target := p.ImmuneSystem
	for t := 0; t < 2; t++ {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}
		for scanner.Scan() {
			var num, hp, pow, init int
			immune := make([]string, 0, 4)
			weak := make([]string, 0, 4)
			var damagetype string

			line := scanner.Text()
			if line == "" {
				scanner.Scan()
				scanner.Scan()
				p.ImmuneSystem = target
				target = p.Infection
				continue
			}

			parts := strings.FieldsFunc(line, func(c rune) bool {
				return c == '(' || c == ')'
			})

			vparts := strings.Split(parts[1], ";")
			for _, vul := range vparts {
				vul = strings.Trim(vul, " ")
				vulp := strings.FieldsFunc(vul, func(c rune) bool {
					return c == ' ' || c == ','
				})
				for t := 2; t < len(vulp); t++ {
					switch vulp[0] {
					case "weak":
						weak = append(weak, vulp[t])
					case "immune":
						immune = append(immune, vulp[t])
					}
				}
			}

			fmt.Sscanf(parts[0], "%d units each with %d hit points", &num, &hp)
			fmt.Sscanf(parts[2], " with an attack that does %d %s damage at initiative %d", &pow, &damagetype, &init)

			target = append(target, Group{
				Amount:     num,
				Damage:     pow,
				DamageType: damagetype,
				HP:         hp,
				Immune:     immune,
				Initiative: init,
				Weak:       weak,
			})
		}
		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}
	}
	p.Infection = target

	spew.Dump(p)

	input.Close()

	return &p
}

func (p *Puzzle) Sim() {
}

func main() {
	// Test1
	t := NewPuzzle("test1.txt")
	t.Sim()

	// Puzzle1
	/*p := NewPuzzle("input.txt")
	p.Sim()*/
}
