package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"os"
	"sort"

	col "github.com/fatih/color"
	"github.com/pkg/errors"
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
	Color *col.Color
}

func (t TileDef) String() string {
	return t.Color.Sprintf("%c", t.Char)
}

var TileMap = map[TileID]*TileDef{
	TileIDUnknown: &TileDef{
		ID:    TileIDUnknown,
		Char:  '?',
		Color: col.New(col.FgBlack).Add(col.BgMagenta),
	},
	TileIDGround: &TileDef{
		ID:    TileIDGround,
		Char:  '.',
		Color: col.New(col.FgYellow),
	},
	TileIDWall: &TileDef{
		ID:    TileIDWall,
		Char:  '#',
		Color: col.New(col.FgHiWhite),
	},
}

var CreatureMap = map[TileID]*TileDef{
	TileIDElf: &TileDef{
		ID:    TileIDElf,
		Char:  'E',
		Color: col.New(col.FgRed),
	},
	TileIDGoblin: &TileDef{
		ID:    TileIDGoblin,
		Char:  'G',
		Color: col.New(col.FgGreen),
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
	GetLPos() (int, int)
	GetPower() int
	SetPower(int)
	GetTile() *TileDef
	Move(int, int)
	Hurt(int)
	Clone() Creature
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
	X, Y   int
	LX, LY int
	Hp     int
	Power  int
	Tile   *TileDef
}

func NewElf(x, y int) *Elf {
	return &Elf{
		X:     x,
		Y:     y,
		LX:    -1,
		LY:    -1,
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

func (e *Elf) GetLPos() (int, int) {
	return e.LX, e.LY
}

func (e *Elf) GetPower() int {
	return e.Power
}

func (e *Elf) SetPower(power int) {
	e.Power = power
}

func (e *Elf) GetTile() *TileDef {
	return e.Tile
}

func (e *Elf) Move(dx, dy int) {
	e.LX, e.LY = e.X, e.Y
	e.X += dx
	e.Y += dy
}

func (e *Elf) Hurt(dmg int) {
	e.Hp -= dmg
}

func (e *Elf) Clone() Creature {
	return &Elf{
		X:     e.X,
		Y:     e.Y,
		LX:    e.LX,
		LY:    e.LY,
		Hp:    e.Hp,
		Power: e.Power,
		Tile:  e.Tile,
	}
}

type Goblin struct {
	X, Y   int
	LX, LY int
	Hp     int
	Power  int
	Tile   *TileDef
}

func NewGoblin(x, y int) *Goblin {
	return &Goblin{
		X:     x,
		Y:     y,
		LX:    -1,
		LY:    -1,
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

func (g *Goblin) GetLPos() (int, int) {
	return g.LX, g.LY
}

func (g *Goblin) GetPower() int {
	return g.Power
}

func (g *Goblin) SetPower(power int) {
	g.Power = power
}

func (g *Goblin) GetTile() *TileDef {
	return g.Tile
}

func (g *Goblin) Move(dx, dy int) {
	g.LX, g.LY = g.X, g.Y
	g.X += dx
	g.Y += dy
}

func (g *Goblin) Hurt(dmg int) {
	g.Hp -= dmg
}

func (g *Goblin) Clone() Creature {
	return &Goblin{
		X:     g.X,
		Y:     g.Y,
		LX:    g.LX,
		LY:    g.LY,
		Hp:    g.Hp,
		Power: g.Power,
		Tile:  g.Tile,
	}
}

var GridWidth int
var GridSize int

type Puzzle struct {
	World     []Tile
	Creatures []Creature
}

type Creatures []Creature

func (c Creatures) Filter(id TileID) []Creature {
	creatures := make([]Creature, 0)
	for _, cc := range c {
		if cc.GetTile().ID == id {
			creatures = append(creatures, cc)
		}
	}
	return creatures
}

func (c Creatures) Find(x, y int) Creature {
	for _, cc := range c {
		if cc.GetHp() <= 0 {
			continue
		}
		cx, cy := cc.GetPos()
		if cx == x && cy == y {
			return cc
		}
	}
	return nil
}

func printRuneGrid(grid []rune) {
	for y := 0; y < GridWidth; y++ {
		for x := 0; x < GridWidth; x++ {
			i := (y * GridWidth) + x
			r := grid[i]

			switch r {
			case 0:
				r = 'S'
			case 252:
				r = 't'
			case 253:
				r = ' '
			case 254:
				r = '.'
			case 255:
				r = 'T'
			case 999:
				r = 'o'
			default:
				if r > 255 {
					r = '*'
				} else {
					r = '+'
				}
			}

			fmt.Printf("%c", r)
		}
		fmt.Printf("\n")
	}
	for t := 0; t < 42; t++ {
		fmt.Printf("\n")
	}
}

type GridToGIF struct {
	gridWidth int
	grid      [][]rune
}

func NewGridToGIF(width int) *GridToGIF {
	return &GridToGIF{
		gridWidth: width,
		grid:      make([][]rune, 0, 100),
	}
}

func (g *GridToGIF) AddFrame(grid []rune) {
	gr := make([]rune, len(grid))
	copy(gr, grid)
	g.grid = append(g.grid, gr)
}

func (g *GridToGIF) Generate(filepath string) error {
	palette := []color.Color{
		color.RGBA{0x0, 0x0, 0x0, 0xff},    // Empty
		color.RGBA{0x0, 0xff, 0x0, 0xff},   // Start
		color.RGBA{0xff, 0x0, 0x0, 0xff},   // End
		color.RGBA{0x66, 0x66, 0x66, 0xff}, // Ground
		color.RGBA{0x99, 0x99, 0x99, 0xff}, // Potential
		color.RGBA{0xff, 0xff, 0xff, 0xff}, // Visited
		color.RGBA{0xff, 0x99, 0x0, 0xff},  // Path home
		color.RGBA{0x0, 0x99, 0x0, 0xff},   // Friendly
	}
	size := image.Rect(0, 0, 32, 32)
	images := make([]*image.Paletted, 0, 100)
	delays := make([]int, 0, 100)
	disposals := make([]byte, 0, 100)
	for _, grid := range g.grid {
		image := image.NewPaletted(size, palette)

		for y := 0; y < g.gridWidth; y++ {
			for x := 0; x < g.gridWidth; x++ {
				i := (y * g.gridWidth) + x

				col := palette[0]
				switch grid[i] {
				case 0:
					fallthrough
				case 256:
					col = palette[1]
				case 252:
					col = palette[7]
				case 253:
					col = palette[0]
				case 254:
					col = palette[3]
				case 255:
					col = palette[2]
				case 999:
					col = palette[6]
				default:
					if grid[i] > 255 {
						col = palette[5]
					} else {
						col = palette[4]
					}
				}
				/*for py := 0; py < 8; py++ {
					for px := 0; px < 8; px++ {(
						image.Set((x*8)+px, (y*8)+py, col)
					}
				}*/
				image.Set(x, y, col)
			}
		}

		images = append(images, image)
		delays = append(delays, 10)
		disposals = append(disposals, gif.DisposalPrevious)
	}

	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return errors.Wrap(err, "could not open output file")
	}
	defer f.Close()
	gif.EncodeAll(f, &gif.GIF{
		Image:    images,
		Delay:    delays,
		Disposal: disposals,
	})
	return nil
}

func (p *Puzzle) djikstra(sx, sy int, friendlies []Creature, targets []Creature, gifGen *GridToGIF) (int, int) {
	grid := make([]rune, len(p.World))
	for y := 0; y < GridWidth; y++ {
	row:
		for x := 0; x < GridWidth; x++ {
			i := (y * GridWidth) + x
			if x == sx && y == sy {
				grid[i] = 0
				continue
			}
			for _, c := range friendlies {
				if c.GetHp() <= 0 {
					continue
				}
				cx, cy := c.GetPos()
				if x == cx && y == cy {
					grid[i] = 252
					continue row
				}
			}
			for _, c := range targets {
				if c.GetHp() <= 0 {
					continue
				}
				cx, cy := c.GetPos()
				if x == cx && y == cy {
					grid[i] = 255
					continue row
				}
			}
			if p.World[i].Def.Char == '.' {
				grid[i] = 254
				continue
			}
			grid[i] = 253
		}
	}

	source := (sy * GridWidth) + sx
	current := source
	found := false
	for {
		// Calc neighbor scores
		// This "should" be safe since it can never reach the bounds of the map
		up := current - GridWidth
		dn := current + GridWidth
		lt := current - 1
		rt := current + 1
		for _, c := range targets {
			if c.GetHp() <= 0 {
				continue
			}
			cx, cy := c.GetPos()
			target := (cy * GridWidth) + cx
			if up == target || dn == target || lt == target || rt == target {
				//fmt.Println("found target")
				found = true
				grid[current] += 256
				current = target
				break
			}
		}
		if found {
			break
		}

		if grid[up] == 254 {
			grid[up] = grid[current] + 1
		}
		if grid[dn] == 254 {
			grid[dn] = grid[current] + 1
		}
		if grid[lt] == 254 {
			grid[lt] = grid[current] + 1
		}
		if grid[rt] == 254 {
			grid[rt] = grid[current] + 1
		}
		grid[current] += 256

		// Choose neighbor
		lowest := rune(252)
		for i, score := range grid {
			if score < lowest {
				current = i
				lowest = score
			}
		}
		if lowest == 252 {
			// fmt.Println("END, no match")
			return -1, -1
		}

		//gifGen.AddFrame(grid)
		//printRuneGrid(grid)
		//time.Sleep(10 * time.Millisecond)
	}

	if found {
		for {
			// Choose neighbor
			lowest := rune(998)
			up := current - GridWidth
			dn := current + GridWidth
			lt := current - 1
			rt := current + 1
			if up == source || dn == source || lt == source || rt == source {
				// fmt.Println("found path", current%GridWidth, current/GridWidth)
				return current % GridWidth, current / GridWidth
			}
			if grid[up] > 255 && grid[up] < lowest {
				lowest = grid[up]
				current = up
			}
			if grid[dn] > 255 && grid[dn] < lowest {
				lowest = grid[dn]
				current = dn
			}
			if grid[lt] > 255 && grid[lt] < lowest {
				lowest = grid[lt]
				current = lt
			}
			if grid[rt] > 255 && grid[rt] < lowest {
				lowest = grid[rt]
				current = rt
			}
			grid[current] = 999

			//gifGen.AddFrame(grid)
			//printRuneGrid(grid)
			//waitEnter()
			//time.Sleep(10 * time.Millisecond)
		}
	}
	return -1, -1
}

func (p *Puzzle) getCloseTarget(creature Creature, targets Creatures) (Creature, bool) {
	cx, cy := creature.GetPos()

	up := targets.Find(cx, cy-1)
	dn := targets.Find(cx, cy+1)
	lt := targets.Find(cx-1, cy)
	rt := targets.Find(cx+1, cy)

	// By health first, then reading order
	var target Creature
	lowest := 9999
	if up != nil {
		hp := up.GetHp()
		if hp < lowest {
			lowest = hp
			target = up
		}
	}
	if lt != nil {
		hp := lt.GetHp()
		if hp < lowest {
			lowest = hp
			target = lt
		}
	}
	if rt != nil {
		hp := rt.GetHp()
		if hp < lowest {
			lowest = hp
			target = rt
		}
	}
	if dn != nil {
		hp := dn.GetHp()
		if hp < lowest {
			lowest = hp
			target = dn
		}
	}

	return target, target != nil
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

var aliveColor = col.New(col.FgGreen)
var deadColor = col.New(col.FgRed)

func (p *Puzzle) Solution1(pretty bool) (int, int) {
	turn := 1
	gifGen := NewGridToGIF(GridWidth)
	for {
		fmt.Printf("\n### Turn %d ###\n", turn)
		fmt.Printf("Turn order: ")
		sort.Sort(ByCoord(p.Creatures))
		for _, c := range p.Creatures {
			col := aliveColor
			if c.GetHp() < 1 {
				col = deadColor
			}
			col.Printf("%s (%dhp),", c.GetTile().ID, c.GetHp())
		}
		fmt.Printf("\n")

		for t, c := range p.Creatures {
			if c.GetHp() < 1 {
				continue
			}

			aliveGoblins := 0
			aliveElves := 0
			for _, c := range p.Creatures {
				if c.GetHp() < 1 {
					continue
				}
				switch c.GetTile().ID {
				case TileIDElf:
					aliveElves++
				case TileIDGoblin:
					aliveGoblins++
				}
			}
			if aliveGoblins == 0 || aliveElves == 0 {
				hp := 0
				for _, c := range p.Creatures {
					if c.GetHp() < 1 {
						continue
					}
					hp += c.GetHp()
				}
				fmt.Println("genning gif")
				gifGen.Generate("out.gif")
				fmt.Println("DONE")
				return turn - 1, hp
			}

			fmt.Printf("\n=== #%d, %s ===\n", t, c.GetTile().ID)

			fmt.Printf("\n")
			cx, cy := c.GetPos()
			targetID := TileIDUnknown
			switch c.GetTile().ID {
			case TileIDElf:
				targetID = TileIDGoblin
			case TileIDGoblin:
				targetID = TileIDElf
			}

			target, ok := p.getCloseTarget(c, Creatures(p.Creatures).Filter(targetID))
			if !ok {
				tx, ty := p.djikstra(cx, cy, Creatures(p.Creatures).Filter(c.GetTile().ID), Creatures(p.Creatures).Filter(targetID), gifGen)
				if tx < 0 || ty < 0 {
					fmt.Printf("#%d, %s have no targets and is confused...\n", t, c.GetTile().ID)
					continue
				}
				c.Move(tx-cx, ty-cy)
			}

			target, ok = p.getCloseTarget(c, Creatures(p.Creatures).Filter(targetID))
			if ok {
				fmt.Printf("%s attacks %s: ", c.GetTile().ID, target.GetTile().ID)
				target.Hurt(c.GetPower())
				fmt.Printf("%s takes %ddmg ", target.GetTile().ID, c.GetPower())
				if target.GetHp() < 1 {
					fmt.Printf("and dies...")
				} else {
					fmt.Printf("and have %dHP left", target.GetHp())
				}
				fmt.Printf("\n")
			}

			{
				world := make([]rune, len(p.World))
				elves := Creatures(p.Creatures).Filter(TileIDElf)
				goblins := Creatures(p.Creatures).Filter(TileIDGoblin)
				for y := 0; y < GridWidth; y++ {
				row:
					for x := 0; x < GridWidth; x++ {
						i := (y * GridWidth) + x
						for _, c := range goblins {
							if c.GetHp() <= 0 {
								continue
							}
							cx, cy := c.GetPos()
							if x == cx && y == cy {
								world[i] = 252
								continue row
							}
						}
						for _, c := range elves {
							if c.GetHp() <= 0 {
								continue
							}
							cx, cy := c.GetPos()
							if x == cx && y == cy {
								world[i] = 255
								continue row
							}
						}
						if p.World[i].Def.Char == '.' {
							world[i] = 254
							continue
						}
						world[i] = 253
					}
				}
				gifGen.AddFrame(world)
			}
			// p.PrintState()
			// time.Sleep(10 * time.Millisecond)
			/*waitEnter()*/
		}
		turn++
	}
}

func (p *Puzzle) Solution2(pretty bool) (int, int) {
	creatures := make([]Creature, len(p.Creatures))
	for t := range p.Creatures {
		cx, cy := p.Creatures[t].GetPos()
		creatures[t] = NewCreature(p.Creatures[t].GetTile().ID, cx, cy)
	}
	pow := 4
	for {
		for t := range p.Creatures {
			cx, cy := creatures[t].GetPos()
			p.Creatures[t] = NewCreature(creatures[t].GetTile().ID, cx, cy)
		}
		elves := Creatures(p.Creatures).Filter(TileIDElf)
		for _, e := range elves {
			e.SetPower(pow)
		}
		turns, hp := p.Solution1(true)
		elves = Creatures(p.Creatures).Filter(TileIDElf)
		alive := 0
		for _, e := range elves {
			if e.GetHp() > 0 {
				alive++
			}
		}
		if alive >= len(elves) {
			fmt.Println(pow, turns, hp, turns*hp)
			p.PrintState()
			waitEnter()
			return turns, hp
		}
		pow++
	}
}

func (p *Puzzle) PrintState() {
	for y := 0; y < GridWidth; y++ {
		for x := 0; x < GridWidth; x++ {
			i := (y * GridWidth) + x
			found := false
			for _, c := range p.Creatures {
				if c.GetHp() < 1 {
					continue
				}
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

	p.PrintState()
	turns, hp := p.Solution1(true)
	p.PrintState()

	turns, hp = p.Solution2(true)
	fmt.Println(turns, hp, turns*hp)
}
