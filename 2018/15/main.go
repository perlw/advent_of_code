package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"

	"github.com/fatih/color"
)

func waitEnter() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

type TileID int

const (
	TileIDUnknown TileID = iota
	TileIDGround
	TileIDWall

	TileIDElf
	TileIDGoblin
)

func (t TileID) String() string {
	switch t {
	case TileIDUnknown:
		return "unknown"
	case TileIDGround:
		return "ground"
	case TileIDWall:
		return "wall"
	case TileIDElf:
		return "elf"
	case TileIDGoblin:
		return "goblin"
	default:
		return "n/a"
	}
}

type TileDef struct {
	ID    TileID
	Char  rune
	Color *color.Color
}

func (t TileDef) String() string {
	return t.Color.Sprintf("%c", t.Char)
}

var TileMap = map[TileID]*TileDef{
	TileIDUnknown: &TileDef{
		ID:    TileIDUnknown,
		Char:  '?',
		Color: color.New(color.FgBlack).Add(color.BgMagenta),
	},
	TileIDGround: &TileDef{
		ID:    TileIDGround,
		Char:  '.',
		Color: color.New(color.FgYellow),
	},
	TileIDWall: &TileDef{
		ID:    TileIDWall,
		Char:  '#',
		Color: color.New(color.FgHiWhite),
	},
}

var CreatureMap = map[TileID]*TileDef{
	TileIDElf: &TileDef{
		ID:    TileIDElf,
		Char:  'E',
		Color: color.New(color.FgRed),
	},
	TileIDGoblin: &TileDef{
		ID:    TileIDGoblin,
		Char:  'G',
		Color: color.New(color.FgGreen),
	},
}

type Tile struct {
	Def *TileDef
}

func findTileDef(tileMap map[TileID]*TileDef, char rune) (*TileDef, TileID) {
	for id, t := range tileMap {
		if t.Char == char {
			return t, id
		}
	}
	return nil, TileIDUnknown
}

type Creature interface {
	GetHp() int
	GetPos() (int, int)
	GetPower() int
	GetTile() *TileDef
	Move(int, int)
}

func NewCreature(tileID TileID, x, y int) Creature {
	var creature Creature
	switch tileID {
	case TileIDElf:
		creature = NewElf(x, y)
	case TileIDGoblin:
		creature = NewGoblin(x, y)
	}
	return creature
}

type Elf struct {
	X, Y  int
	Hp    int
	Power int
	Tile  *TileDef
}

func NewElf(x, y int) *Elf {
	return &Elf{
		X:     x,
		Y:     y,
		Hp:    200,
		Power: 3,
		Tile:  CreatureMap[TileIDElf],
	}
}

func (e *Elf) GetHp() int {
	return e.Hp
}

func (e *Elf) GetPos() (int, int) {
	return e.X, e.Y
}

func (e *Elf) GetPower() int {
	return e.Power
}

func (e *Elf) GetTile() *TileDef {
	return e.Tile
}

func (e *Elf) Move(dx, dy int) {
	e.X += dx
	e.Y += dy
}

type Goblin struct {
	X, Y  int
	Hp    int
	Power int
	Tile  *TileDef
}

func NewGoblin(x, y int) *Goblin {
	return &Goblin{
		X:     x,
		Y:     y,
		Hp:    200,
		Power: 3,
		Tile:  CreatureMap[TileIDGoblin],
	}
}

func (g *Goblin) GetHp() int {
	return g.Hp
}

func (g *Goblin) GetPos() (int, int) {
	return g.X, g.Y
}

func (g *Goblin) GetPower() int {
	return g.Power
}

func (g *Goblin) GetTile() *TileDef {
	return g.Tile
}

func (g *Goblin) Move(dx, dy int) {
	g.X += dx
	g.Y += dy
}

var GridWidth int
var GridSize int

type Puzzle struct {
	World     []Tile
	Creatures []Creature
}

type FloodDir int

const (
	FloodDirN FloodDir = iota
	FloodDirS
	FloodDirW
	FloodDirE
)

func abs(v int) int {
	if v < 0 {
		v = -v
	}
	return v
}

func (p *Puzzle) flood(reach []rune, sx, sy int, dir FloodDir, creature TileID) {
	switch dir {
	case FloodDirN:
		sy--
	case FloodDirS:
		sy++
	case FloodDirW:
		sx--
	case FloodDirE:
		sx++
	}

	if sx < 0 || sy < 0 || sx > GridWidth-1 || sy > GridWidth-1 {
		return
	}

	i := (sy * GridWidth) + sx
	if reach[i] == '+' || p.World[i].Def.Char == '#' {
		return
	}
	close := false
	for _, c := range p.Creatures {
		cx, cy := c.GetPos()
		if sx == cx && sy == cy {
			if c.GetTile().ID != creature {
				reach[i] = 'X'
			} else {
				reach[i] = '-'
			}
			return
		}
		if c.GetTile().ID != creature {
			d := abs(cx-sx) + abs(cy-sy)
			if d < 2 {
				close = true
			}
		}
	}

	if close {
		reach[i] = '!'
	} else {
		reach[i] = '+'
	}

	switch dir {
	case FloodDirN:
		p.flood(reach, sx, sy, FloodDirN, creature)
		p.flood(reach, sx, sy, FloodDirW, creature)
		p.flood(reach, sx, sy, FloodDirE, creature)
	case FloodDirS:
		p.flood(reach, sx, sy, FloodDirS, creature)
		p.flood(reach, sx, sy, FloodDirW, creature)
		p.flood(reach, sx, sy, FloodDirE, creature)
	case FloodDirW:
		p.flood(reach, sx, sy, FloodDirN, creature)
		p.flood(reach, sx, sy, FloodDirS, creature)
		p.flood(reach, sx, sy, FloodDirW, creature)
	case FloodDirE:
		p.flood(reach, sx, sy, FloodDirN, creature)
		p.flood(reach, sx, sy, FloodDirS, creature)
		p.flood(reach, sx, sy, FloodDirE, creature)
	}
}

func (p *Puzzle) floodFillReach(sx, sy int, creature TileID) []rune {
	reach := make([]rune, GridSize)
	for y := 0; y < GridWidth; y++ {
		for x := 0; x < GridWidth; x++ {
			i := (y * GridWidth) + x
			reach[i] = ' '
		}
	}

	p.flood(reach, sx, sy, FloodDirN, creature)
	p.flood(reach, sx, sy, FloodDirS, creature)
	p.flood(reach, sx, sy, FloodDirW, creature)
	p.flood(reach, sx, sy, FloodDirE, creature)

	return reach
}

type ByCoord []Creature

func (c ByCoord) Len() int {
	return len(c)
}

func (c ByCoord) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c ByCoord) Less(i, j int) bool {
	cax, cay := c[i].GetPos()
	cbx, cby := c[j].GetPos()
	ia := (cay * GridWidth) + cax
	ib := (cby * GridWidth) + cbx
	return ia < ib
}

type Point []int

func (p *Puzzle) Solution1(pretty bool) int {
	turn := 1
	for {
		fmt.Printf("\n### Turn %d ###\n", turn)
		for _, c := range p.Creatures {
			fmt.Printf("%s,", c.GetTile().ID)
		}
		fmt.Printf("\n")

		// fmt.Printf("Turn order: ")
		sort.Sort(ByCoord(p.Creatures))
		for _, c := range p.Creatures {
			fmt.Printf("%s,", c.GetTile().ID)
		}
		fmt.Printf("\n")

		for t, c := range p.Creatures {
			fmt.Printf("\n=== #%d, %s ===\n", t, c.GetTile().ID)

			p.PrintState()

			fmt.Printf("\n")
			cx, cy := c.GetPos()
			reach := p.floodFillReach(cx, cy, c.GetTile().ID)
			for y := 0; y < GridWidth; y++ {
				for x := 0; x < GridWidth; x++ {
					i := (y * GridWidth) + x
					fmt.Printf("%c", reach[i])
				}
				fmt.Printf("\n")
			}

			nextToTarget := -1
			targets := make([]Creature, 0, len(reach))
			positions := make([]Point, 0, len(reach))
			for y := 0; y < GridWidth; y++ {
				for x := 0; x < GridWidth; x++ {
					i := (y * GridWidth) + x
					if reach[i] == 'X' {
						for t, cc := range p.Creatures {
							cxx, cyy := cc.GetPos()
							if x == cxx && y == cyy {
								if nextToTarget < 0 {
									d := abs(cx-cxx) + abs(cy-cyy)
									if d < 2 {
										nextToTarget = t
									}
								}

								targets = append(targets, cc)
							}
						}
					} else if reach[i] == '!' {
						positions = append(positions, Point{x, y})
					}
				}
			}

			if len(targets) == 0 {
				fmt.Println("ALL DEAD")
				return 0
			}

			fmt.Println(len(targets), targets)
			fmt.Println(len(positions), positions)

			fmt.Println("next to:", nextToTarget)
			if nextToTarget < 0 {
				// n, w, e, s
				lowest := 100
				dir := -1
				dirs := []int{100, 100, 100, 100}
				for _, p := range positions {
					var d int
					d = abs(cx-p[0]) + abs(cy-p[1]-1)
					if d < dirs[0] {
						dirs[0] = d
					}
					if d < lowest {
						lowest = d
						dir = 0
					}
					d = abs(cx-p[0]-1) + abs(cy-p[1])
					if d < dirs[1] {
						dirs[1] = d
					}
					if d < lowest {
						lowest = d
						dir = 1
					}
					d = abs(cx-p[0]+1) + abs(cy-p[1])
					if d < dirs[2] {
						dirs[2] = d
					}
					if d < lowest {
						lowest = d
						dir = 2
					}
					d = abs(cx-p[0]) + abs(cy-p[1]+1)
					if d < dirs[3] {
						dirs[3] = d
					}
					if d < lowest {
						lowest = d
						dir = 3
					}
				}

				switch dir {
				case 0:
					c.Move(0, -1)
				case 1:
					c.Move(-1, 0)
				case 2:
					c.Move(1, 0)
				case 3:
					c.Move(0, 1)
				}
			}

			if nextToTarget > -1 {
			}

			waitEnter()
		}
		turn++
	}

	return 0
}

func (p *Puzzle) PrintState() {
	for y := 0; y < GridWidth; y++ {
		for x := 0; x < GridWidth; x++ {
			i := (y * GridWidth) + x
			found := false
			for _, c := range p.Creatures {
				cx, cy := c.GetPos()
				if cx == x && cy == y {
					found = true
					fmt.Printf("%s", c.GetTile())
					break
				}
			}
			if !found {
				fmt.Printf("%s", p.World[i].Def)
			}
		}
		fmt.Printf("\n")
	}
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	p := Puzzle{
		World:     make([]Tile, 0, 32),
		Creatures: make([]Creature, 0, 32),
	}
	lines := 0
	for scanner.Scan() {
		line := []byte(scanner.Text())
		row := make([]Tile, 0, 32)
		for cNum, c := range line {
			tileDef, _ := findTileDef(TileMap, rune(c))
			if tileDef == nil {
				_, tileID := findTileDef(CreatureMap, rune(c))
				if tileID == TileIDUnknown {
					row = append(row, Tile{
						Def: TileMap[TileIDUnknown],
					})
					continue
				}

				row = append(row, Tile{
					Def: TileMap[TileIDGround],
				})

				p.Creatures = append(p.Creatures, NewCreature(tileID, cNum, lines))
				continue
			}

			row = append(row, Tile{
				Def: tileDef,
			})
		}
		p.World = append(p.World, row...)
		lines++
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	GridWidth = lines
	GridSize = GridWidth * GridWidth

	fmt.Println(p.Solution1(true))
}
