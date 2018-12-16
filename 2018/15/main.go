package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/fatih/color"
)

type TileID int

const (
	TileIDUnknown TileID = iota
	TileIDGround
	TileIDWall

	TileIDElf
	TileIDGoblin
)

type TileDef struct {
	Char  rune
	Color *color.Color
}

func (t TileDef) String() string {
	return t.Color.Sprintf("%c", t.Char)
}

var TileMap = map[TileID]*TileDef{
	TileIDUnknown: &TileDef{
		Char:  '?',
		Color: color.New(color.FgBlack).Add(color.BgMagenta),
	},
	TileIDGround: &TileDef{
		Char:  '.',
		Color: color.New(color.FgYellow),
	},
	TileIDWall: &TileDef{
		Char:  '#',
		Color: color.New(color.FgHiWhite),
	},
}

var CreatureMap = map[TileID]*TileDef{
	TileIDElf: &TileDef{
		Char:  'E',
		Color: color.New(color.FgRed),
	},
	TileIDGoblin: &TileDef{
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

type Puzzle struct {
	GridWidth int
	GridSize  int
	World     []Tile
	Creatures []Creature
}

func (p *Puzzle) Solution1(pretty bool) int {
	p.PrintState()

	return 0
}

func (p *Puzzle) PrintState() {
	for y := 0; y < p.GridWidth; y++ {
		for x := 0; x < p.GridWidth; x++ {
			i := (y * p.GridWidth) + x
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
	p.GridWidth = lines
	p.GridSize = p.GridWidth * 2

	fmt.Println(p.Solution1(true))
}
