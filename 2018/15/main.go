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
)

type TileDef struct {
	Char  rune
	Color *color.Color
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

type Tile struct {
	Def *TileDef
}

func (t Tile) String() string {
	return t.Def.Color.Sprintf("%c", t.Def.Char)
}

func findTileDef(char rune) *TileDef {
	for _, t := range TileMap {
		if t.Char == char {
			return t
		}
	}
	return nil
}

type Creature interface {
	GetHp() int
	GetPos() (int, int)
}

type Elf struct {
	X, Y int
	Hp   int
}

type Goblin struct {
	X, Y int
	Hp   int
}

type Puzzle struct {
	GridWidth int
	GridSize  int
	World     []Tile
	Creatures []Creature
}

func (p *Puzzle) Solution1(pretty bool) int {
	p.PrintWorld()

	return 0
}

func (p *Puzzle) PrintWorld() {
	for y := 0; y < p.GridWidth; y++ {
		for x := 0; x < p.GridWidth; x++ {
			i := (y * p.GridWidth) + x
			fmt.Printf("%s", p.World[i])
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
		World: make([]Tile, 0, 32),
	}
	lines := 0
	for scanner.Scan() {
		line := []byte(scanner.Text())
		row := make([]Tile, 0, 32)
		for _, c := range line {
			tileDef := findTileDef(rune(c))
			if tileDef == nil {
				row = append(row, Tile{
					Def: TileMap[TileIDUnknown],
				})
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
