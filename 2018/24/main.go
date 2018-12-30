package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/davecgh/go-spew/spew"
)

type GroupType int

const (
	GroupImmune GroupType = iota
	GroupInfection
)

func (g GroupType) String() string {
	if g == GroupImmune {
		return "immunesystem"
	}
	return "infection"
}

type Group struct {
	ID         int
	Type       GroupType
	Amount     int
	HP         int
	Weak       []string
	Immune     []string
	Damage     int
	DamageType string
	Initiative int
	EP         int
}

func (g *Group) PotentialDamage(target *Group) int {
	if inSlice(g.DamageType, target.Immune) {
		return 0
	}
	if inSlice(g.DamageType, target.Weak) {
		return g.EP * 2
	}
	return g.EP
}

type ByInit []Group

func (b ByInit) Len() int {
	return len(b)
}

func (b ByInit) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func (b ByInit) Less(i, j int) bool {
	return b[i].Initiative > b[j].Initiative
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

func inSlice(a string, b []string) bool {
	for _, c := range b {
		if a == c {
			return true
		}
	}
	return false
}

func findGroup(groups []Group, id int, groupType GroupType) *Group {
	for t := range groups {
		if groups[t].ID == id && groups[t].Type == groupType {
			return &groups[t]
		}
	}
	return nil
}

func existsInTurnOrder(group *Group, turnorder []Turn) *Group {
	for t := range turnorder {
		if turnorder[t].Target.ID == group.ID && turnorder[t].Target.Type == group.Type {
			return turnorder[t].Target
		}
	}
	return nil
}

func findTarget(attacker *Group, turnorder []Turn, groups []Group) (*Group, int) {
	highest := -1
	index := -1
	init := -1
	hEP := -1
	for t := range groups {
		if groups[t].Amount <= 0 {
			continue
		}
		if (attacker.ID == groups[t].ID && attacker.Type == groups[t].Type) || attacker.Type == groups[t].Type {
			continue
		}
		if existsInTurnOrder(&groups[t], turnorder) != nil {
			continue
		}

		potential := attacker.PotentialDamage(&groups[t])
		if potential > highest {
			highest = potential
			index = t
			init = groups[t].Initiative
			hEP = groups[t].EP
		} else if potential == highest && (groups[t].EP > hEP || groups[t].Initiative > init) {
			highest = potential
			index = t
			init = groups[t].Initiative
			hEP = groups[t].EP
		}
	}
	if index > -1 {
		return &groups[index], highest
	}
	return nil, -1
}

type Point struct {
	X, Y, Z int
}

type Puzzle struct {
	Groups []Group
}

func NewPuzzle(filepath string, boost int) *Puzzle {
	p := Puzzle{
		Groups: make([]Group, 0, 10),
	}

	input, err := os.Open(filepath)
	if err != nil {
		panic(err.Error())
	}
	scanner := bufio.NewScanner(input)

	id := 1
	groupType := GroupImmune
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
				id = 1
				groupType = GroupInfection
				continue
			}

			parts := strings.FieldsFunc(line, func(c rune) bool {
				return c == '(' || c == ')'
			})

			if len(parts) > 1 {
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
			}

			fmt.Sscanf(parts[0], "%d units each with %d hit points", &num, &hp)
			index := 0
			if len(parts) > 1 {
				index = 2
			}
			fmt.Sscanf(parts[index], " with an attack that does %d %s damage at initiative %d", &pow, &damagetype, &init)

			if groupType == GroupImmune {
				pow += boost
			}

			p.Groups = append(p.Groups, Group{
				ID:         id,
				Type:       groupType,
				Amount:     num,
				Damage:     pow,
				DamageType: damagetype,
				HP:         hp,
				Immune:     immune,
				Initiative: init,
				Weak:       weak,
				EP:         num * pow,
			})
			id++
		}
		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}
	}

	// spew.Dump(p)

	input.Close()

	return &p
}

type Turn struct {
	Group  *Group
	Target *Group
}

func (t *Turn) Exec() int {
	damage := t.Group.PotentialDamage(t.Target)
	deaths := damage / t.Target.HP
	t.Target.Amount -= deaths
	return deaths
}

func (p *Puzzle) Sim(pretty bool) GroupType {
	for {
		turnorder := make([]Turn, 0, 10)

		sort.Sort(ByInit(p.Groups))

		if pretty {
			fmt.Println("\nTarget phase:")
		}
		// Decide turn order
		for t := range p.Groups {
			if p.Groups[t].Amount <= 0 {
				continue
			}

			target, potential := findTarget(&p.Groups[t], turnorder, p.Groups)
			if target == nil {
				if pretty {
					fmt.Println("no targets")
				}
				continue
			}

			if pretty {
				fmt.Printf("#%d:%s -> #%d:%s for %d damage\n", p.Groups[t].ID, p.Groups[t].Type, target.ID, target.Type, potential)
			}

			turnorder = append(turnorder, Turn{
				Group:  &p.Groups[t],
				Target: target,
			})
		}

		if len(turnorder) == 0 {
			fmt.Println("Combat over")
			if pretty {
				spew.Dump(p.Groups)
			}
			sum := 0
			groupType := GroupInfection
			for _, g := range p.Groups {
				if g.Amount > 0 {
					sum += g.Amount
					fmt.Printf("#%d:%s contains %d units\n", g.ID, g.Type, g.Amount)
					groupType = g.Type
				}
			}
			fmt.Println("Sum alive:", sum, groupType)
			return groupType
		}

		if pretty {
			fmt.Println("\nAttack phase:")
		}
		// Execute
		for _, t := range turnorder {
			if pretty {
				fmt.Printf("#%d:%s -> #%d:%s", t.Group.ID, t.Group.Type, t.Target.ID, t.Target.Type)
				num := t.Exec()
				fmt.Printf(", killing %d units\n", num)
			} else {
				t.Exec()
			}
		}
	}
}

func main() {
	// Test1
	t := NewPuzzle("test1.txt", 0)
	t.Sim(true)

	// Test2
	boost := 1
	for {
		t := NewPuzzle("test1.txt", boost)
		if t.Sim(false) == GroupImmune {
			break
		}
		boost++
	}

	// Puzzle1
	p := NewPuzzle("input.txt", 0)
	p.Sim(false)

	// Puzzle2
	boost = 1
	for {
		p := NewPuzzle("input.txt", boost)
		if p.Sim(false) == GroupImmune {
			break
		}
		boost++
	}
}
