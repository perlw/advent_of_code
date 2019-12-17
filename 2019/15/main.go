package main

import (
	"bufio"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

type OpType int

const (
	OpAdd OpType = iota + 1
	OpMul
	OpInput
	OpOutput
	OpJumpTrue
	OpJumpFalse
	OpLessThan
	OpEquals
	OpRelativeBase
	OpEnd = 99
)

type ModeType int

const (
	ModePosition ModeType = iota
	ModeImmediate
	ModeRelative
)

type Computer struct {
	mem    []int
	pc     int
	rb     int
	Halted bool
}

func NewComputer(initial []int) *Computer {
	com := Computer{}
	com.mem = make([]int, 65535)
	copy(com.mem, initial)
	return &com
}

func (com *Computer) getArgPos(pos int, mode ModeType) (int, error) {
	if pos < 0 || pos >= len(com.mem) {
		return -1, fmt.Errorf("invalid index, out of range: %d", pos)
	}
	switch mode {
	case ModePosition:
		if com.mem[pos] < 0 || com.mem[pos] >= len(com.mem) {
			return -1, fmt.Errorf("invalid position index, out of range: %d", com.mem[pos])
		}
		return com.mem[pos], nil
	case ModeRelative:
		rel := com.mem[pos] + com.rb
		if rel < 0 || rel >= len(com.mem) {
			return -1, fmt.Errorf("invalid relative index, out of range: %d", rel)
		}
		return rel, nil
	}
	return pos, nil
}

func (com *Computer) Iterate(input func() int, yieldOnInput bool, debug bool) (int, bool, error) {
	log := ioutil.Discard
	if debug {
		log = os.Stdout
	}
	for com.pc < len(com.mem) {
		op := OpType(com.mem[com.pc])
		if op == OpEnd {
			fmt.Fprintf(log, "Halt!\n")
			com.Halted = true
			return 0, true, nil
		}

		aMode := ModePosition
		bMode := ModePosition
		cMode := ModePosition
		if op > 99 {
			aMode = ModeType((op / 100) % 10)
			bMode = ModeType((op / 1000) % 10)
			cMode = ModeType((op / 10000) % 10)
			op %= 100
		}

		switch op {
		case OpAdd:
			a, err := com.getArgPos(com.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			b, err := com.getArgPos(com.pc+2, bMode)
			if err != nil {
				return 0, false, err
			}
			o, err := com.getArgPos(com.pc+3, cMode)
			if err != nil {
				return 0, false, err
			}
			com.mem[o] = com.mem[a] + com.mem[b]
			fmt.Fprintf(log, "Setting %d @%d\n", com.mem[o], o)
			com.pc += 4
		case OpMul:
			a, err := com.getArgPos(com.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			b, err := com.getArgPos(com.pc+2, bMode)
			if err != nil {
				return 0, false, err
			}
			o, err := com.getArgPos(com.pc+3, cMode)
			if err != nil {
				return 0, false, err
			}
			com.mem[o] = com.mem[a] * com.mem[b]
			fmt.Fprintf(log, "Setting %d @%d\n", com.mem[o], o)
			com.pc += 4
		case OpInput:
			a, err := com.getArgPos(com.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			com.mem[a] = input()
			fmt.Fprintf(log, "Inputting: %d @%d\n", com.mem[a], a)
			com.pc += 2
			if yieldOnInput {
				return 0, false, nil
			}
		case OpOutput:
			a, err := com.getArgPos(com.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			fmt.Fprintf(log, "Outputting: %d @%d\n", com.mem[a], a)
			com.pc += 2
			return com.mem[a], false, nil
		case OpJumpTrue:
			a, err := com.getArgPos(com.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			b, err := com.getArgPos(com.pc+2, bMode)
			if err != nil {
				return 0, false, err
			}
			if com.mem[a] != 0 {
				com.pc = com.mem[b]
				fmt.Fprintf(log, "Jump to %d\n", com.pc)
			} else {
				com.pc += 3
			}
		case OpJumpFalse:
			a, err := com.getArgPos(com.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			b, err := com.getArgPos(com.pc+2, bMode)
			if err != nil {
				return 0, false, err
			}
			if com.mem[a] == 0 {
				com.pc = com.mem[b]
				fmt.Fprintf(log, "Jump to %d\n", com.pc)
			} else {
				com.pc += 3
			}
		case OpLessThan:
			a, err := com.getArgPos(com.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			b, err := com.getArgPos(com.pc+2, bMode)
			if err != nil {
				return 0, false, err
			}
			o, err := com.getArgPos(com.pc+3, cMode)
			if err != nil {
				return 0, false, err
			}
			if com.mem[a] < com.mem[b] {
				com.mem[o] = 1
			} else {
				com.mem[o] = 0
			}
			com.pc += 4
		case OpEquals:
			a, err := com.getArgPos(com.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			b, err := com.getArgPos(com.pc+2, bMode)
			if err != nil {
				return 0, false, err
			}
			o, err := com.getArgPos(com.pc+3, cMode)
			if err != nil {
				return 0, false, err
			}
			if com.mem[a] == com.mem[b] {
				com.mem[o] = 1
			} else {
				com.mem[o] = 0
			}
			com.pc += 4
		case OpRelativeBase:
			a, err := com.getArgPos(com.pc+1, aMode)
			if err != nil {
				return 0, false, err
			}
			com.rb += com.mem[a]
			com.pc += 2
		default:
			return 0, false, fmt.Errorf("invalid op: %d", op)
		}
	}
	return 0, false, fmt.Errorf("fell out of memory.. pc@%d, max mem: %d", com.pc, len(com.mem))
}

func resize(grid []rune, width, height, newWidth, newHeight, offsetx, offsety int) []rune {
	result := make([]rune, newWidth*newHeight)

	for y := 0; y < newHeight; y++ {
		i := y * newWidth
		for x := 0; x < newWidth; x++ {
			if x >= offsetx && y >= offsety && x-offsetx < width && y-offsety < height {
				j := ((y - offsety) * width) + x - offsetx
				result[i+x] = grid[j]
			} else {
				result[i+x] = ' '
			}
		}
	}

	return result
}

func draw(grid []rune, width, height int) {
	fmt.Printf("\033c")
	for y := 0; y < height; y++ {
		i := y * width
		for x := 0; x < width; x++ {
			fmt.Printf("%c", grid[i+x])
		}
		fmt.Printf("\n")
	}
}

func waitEnter() {
	fmt.Printf("Press ENTER to continue\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

type Move int

const (
	MoveNorth Move = iota + 1
	MoveSouth
	MoveWest
	MoveEast
)

type Status int

const (
	StatusWall Status = iota
	StatusMoved
	StatusFoundOxygen
)

func buildMap(ops []int, findOxygen, debug bool) ([]rune, [2]int, [2]int, [2]int) {
	width, height := 10, 10
	grid := make([]rune, width*height)
	for y := 0; y < height; y++ {
		i := y * width
		for x := 0; x < width; x++ {
			grid[i+x] = ' '
		}
	}

	var px, py int
	var sx, sy int
	var oxyx, oxyy int
	var oxygen bool
	com := NewComputer(ops)
	for !com.Halted {
		ox, oy := 0, 0
		// Grow and offset grid
		if px == 0 {
			ox = 1
		}
		if py == 0 {
			oy = 1
		}
		if ox > 0 || oy > 0 {
			grid = resize(grid, width, height, width+ox, height+oy, ox, oy)
			width += ox
			height += oy
			px += ox
			py += oy
			sx += ox
			sy += oy
			if oxygen {
				oxyx += ox
				oxyy += oy
			}
		}
		// Grow
		ox, oy = 0, 0
		if px >= width-1 {
			ox = 1
		}
		if py >= height-1 {
			oy = 1
		}
		if ox > 0 || oy > 0 {
			grid = resize(grid, width, height, width+ox, height+oy, 0, 0)
			width += ox
			height += oy
		}

		nx, ny := px, py
		out, halt, err := com.Iterate(func() int {
			n := ((ny + 1) * width) + nx
			s := ((ny - 1) * width) + nx
			w := (ny * width) + nx - 1
			e := (ny * width) + nx + 1

			i := (ny * width) + nx
			if grid[n] == ' ' {
				if findOxygen {
					grid[i] = 'v'
				} else {
					grid[i] = '.'
				}
				ny++
				return int(MoveNorth)
			}
			if grid[s] == ' ' {
				if findOxygen {
					grid[i] = '^'
				} else {
					grid[i] = '.'
				}
				ny--
				return int(MoveSouth)
			}
			if grid[w] == ' ' {
				if findOxygen {
					grid[i] = '<'
				} else {
					grid[i] = '.'
				}
				nx--
				return int(MoveWest)
			}
			if grid[e] == ' ' {
				if findOxygen {
					grid[i] = '>'
				} else {
					grid[i] = '.'
				}
				nx++
				return int(MoveEast)
			}

			grid[(py*width)+px] = '`'
			if grid[n] == '^' || grid[n] == 'v' || grid[n] == '<' || grid[n] == '>' || grid[n] == '.' {
				ny++
				return int(MoveNorth)
			}
			if grid[s] == '^' || grid[s] == 'v' || grid[s] == '<' || grid[s] == '>' || grid[s] == '.' {
				ny--
				return int(MoveSouth)
			}
			if grid[w] == '^' || grid[w] == 'v' || grid[w] == '<' || grid[w] == '>' || grid[w] == '.' {
				nx--
				return int(MoveWest)
			}
			if grid[e] == '^' || grid[e] == 'v' || grid[e] == '<' || grid[e] == '>' || grid[e] == '.' {
				nx++
				return int(MoveEast)
			}

			if findOxygen {
				fmt.Println("err, no moves >:o")
				os.Exit(-1)
			}
			return 0
		}, false, false)
		if err != nil {
			log.Fatal(fmt.Errorf("failure while executing computer: %w", err))
		}
		if halt {
			break
		}

		if grid[(py*width)+px] == 'D' {
			grid[(py*width)+px] = ' '
		}
		switch Status(out) {
		case StatusWall:
			grid[(ny*width)+nx] = '#'
		case StatusMoved:
			px, py = nx, ny
		case StatusFoundOxygen:
			grid[(ny*width)+nx] = 'O'
			oxyx, oxyy = nx, ny
			px, py = nx, ny
			oxygen = true
		}
		if !oxygen {
			grid[(py*width)+px] = 'D'
		}

		if debug {
			draw(grid, width, height)
		}
		if findOxygen && oxygen {
			break
		}
		if debug {
			waitEnter()
		}
	}

	return grid, [2]int{width, height}, [2]int{sx, sy}, [2]int{oxyx, oxyy}
}

func RunDroid(ops []int, debug, step bool) int {
	grid, size, initial, oxygen := buildMap(ops, true, false)

	width, height := size[0], size[1]
	px, py := oxygen[0], oxygen[1]
	sx, sy := initial[0], initial[1]

	var steps int
	for {
		n := ((py + 1) * width) + px
		s := ((py - 1) * width) + px
		w := (py * width) + px - 1
		e := (py * width) + px + 1
		grid[(py*width)+px] = '='

		if grid[n] == '^' || grid[n] == 'v' || grid[n] == '<' || grid[n] == '>' {
			py++
		}
		if grid[s] == '^' || grid[s] == 'v' || grid[s] == '<' || grid[s] == '>' {
			py--
		}
		if grid[w] == '^' || grid[w] == 'v' || grid[w] == '<' || grid[w] == '>' {
			px--
		}
		if grid[e] == '^' || grid[e] == 'v' || grid[e] == '<' || grid[e] == '>' {
			px++
		}

		if debug {
			draw(grid, width, height)
		}
		steps++

		if px == sx && py == sy {
			break
		}
		if step {
			waitEnter()
		}
	}

	return steps
}

func FillWithOyxgen(ops []int, debug, step bool) int {
	grid, size, _, oxygen := buildMap(ops, false, false)

	for i := range grid {
		if grid[i] == '.' || grid[i] == '`' || grid[i] == 'D' || grid[i] == 'O' {
			grid[i] = '.'
		} else {
			grid[i] = '#'
		}
	}

	width, height := size[0], size[1]
	i := (oxygen[1] * width) + oxygen[0]
	grid[i] = 'O'

	temp := make([]rune, len(grid))
	var minutes int
	for {
		copy(temp, grid)
		full := true
		for i := range grid {
			if grid[i] == 'O' {
				dir := []int{
					i - width,
					i + width,
					i - 1,
					i + 1,
				}
				for _, d := range dir {
					if grid[d] == '.' {
						temp[d] = 'O'
						full = false
					}
				}
			}
		}
		copy(grid, temp)

		if debug {
			draw(grid, width, height)
		}
		if full {
			break
		}
		minutes++
		if step {
			waitEnter()
		}
	}

	return minutes
}

func main() {
	var debug, step bool
	flag.BoolVar(&debug, "debug", false, "draw debug (implicit when step is enabled)")
	flag.BoolVar(&step, "step", false, "enable step execute")
	flag.Parse()

	debug = debug || step

	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("could not open input file: %w", err))
	}
	defer input.Close()

	raw, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(fmt.Errorf("error while reading input: %w", err))
	}
	ops, err := func(val []byte) ([]int, error) {
		vals := strings.Split(strings.TrimSpace(string(val)), ",")
		result := make([]int, len(vals))
		for i := range vals {
			result[i], err = strconv.Atoi(vals[i])
			if err != nil {
				return nil, fmt.Errorf("could not convert to int: %w", err)
			}
		}
		return result, nil
	}(raw)
	if err != nil {
		log.Fatal(fmt.Errorf("could not read ops: %w", err))
	}

	fmt.Printf("Part 1 - Moves: %d\n", RunDroid(ops, debug, step))
	fmt.Printf("Part 2 - Minutes: %d\n", FillWithOyxgen(ops, debug, step))
}
