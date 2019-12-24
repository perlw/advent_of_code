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

func waitEnter() {
	fmt.Printf("Press ENTER to continue\n")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func drawGrid(grid []rune, width, height int) {
	tiles := map[rune]rune{
		'#': ' ',
		'x': '█',
		'^': '┴',
		'v': '┬',
		'<': '┤',
		'>': '├',
		'.': '░',
		' ': '▓',
	}

	for y := 0; y < height; y++ {
		i := y * width
		for x := 0; x < width; x++ {
			fmt.Printf("%c", tiles[grid[i+x]])
		}
		fmt.Println()
	}
}

func FindAlignment(ops []int, debug bool) (int, []rune, int, int) {
	var width, height int
	grid := make([]rune, 0)

	var x int
	com := NewComputer(ops)
	for !com.Halted {
		out, halt, err := com.Iterate(func() int {
			return 0
		}, false, false)
		if err != nil {
			log.Fatal(fmt.Errorf("failure while executing computer: %w", err))
		}
		if halt {
			break
		}

		fmt.Printf("%c", rune(out))
		if rune(out) == '\n' {
			if width == 0 {
				width = x
			}
			height++
			x = 0
		} else {
			grid = append(grid, rune(out))
			x++
		}
	}
	height--

	temp := make([]rune, len(grid))
	copy(temp, grid)

	grid = make([]rune, (width+2)*(height+2))
	for i := range grid {
		grid[i] = ' '
	}
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			grid[((y+1)*(width+2))+x+1] = temp[(y*width)+x]
		}
	}
	width += 2
	height += 2
	if debug {
		drawGrid(grid, width, height)
	}

	intersections := make([][2]int, 0)
	for y := 0; y < height; y++ {
		i := y * width
		for x := 0; x < width; x++ {
			if grid[i+x] != '#' {
				continue
			}

			dirs := []int{
				i + x + width,
				i + x - width,
				i + x - 1,
				i + x + 1,
			}
			crossing := true
			for _, d := range dirs {
				if d < 0 || d >= len(grid) {
					crossing = false
					break
				}
				if grid[d] != '#' {
					crossing = false
					break
				}
			}
			if crossing {
				intersections = append(intersections, [2]int{x, y})
			}
		}
	}

	if debug {
		fmt.Printf("%+v\n", intersections)
	}

	var result int
	for i := range intersections {
		result += intersections[i][0] * intersections[i][1]
	}
	return result, grid, width, height
}

type Step struct {
	r rune
	i int
}

func scanPath(grid []rune, width, height int, debug bool) []Step {
	result := make([]Step, 0)

	transitions := map[rune]map[rune]Step{
		'^': {'<': {r: 'L', i: -1}, '>': {r: 'R', i: 1}},
		'v': {'>': {r: 'L', i: 1}, '<': {r: 'R', i: -1}},
		'<': {'v': {r: 'L', i: width}, '^': {r: 'R', i: -width}},
		'>': {'^': {r: 'L', i: -width}, 'v': {r: 'R', i: width}},
	}

	px, py, pd := (func() (int, int, rune) {
		for y := 0; y < height; y++ {
			i := y * width
			for x := 0; x < width; x++ {
				r := grid[i+x]
				if r == '^' || r == 'v' || r == '<' || r == '>' {
					return x, y, r
				}
			}
		}
		return 0, 0, 0
	})()

	size := width * height
	for {
		fmt.Printf("@%dx%d,%c\n", px, py, pd)

		i := (py * width) + px
		var step Step
		var found bool
		for d, s := range transitions[pd] {
			it := s.i + i
			if it < 0 || it >= size {
				continue
			}

			if grid[it] == '#' {
				pd = d
				step = s
				found = true
				break
			}
		}
		if !found {
			break
		}

		var num int
		for {
			it := (step.i * num) + i
			if it < 0 || it >= size {
				break
			}

			if grid[it] == '.' || grid[it] == ' ' {
				break
			} else {
				px = it % width
				py = it / width
				num++
				grid[it] = 'x'
			}
		}
		num--
		grid[(py*width)+px] = pd

		result = append(result, Step{
			r: step.r,
			i: num,
		})
		fmt.Printf("Move: %c%d\n", step.r, num)
		if debug {
			drawGrid(grid, width, height)
			waitEnter()
		}
	}
	drawGrid(grid, width, height)

	return result
}

func NotifyRobots(ops []int, grid []rune, width, height int, debug bool) int {
	steps := scanPath(grid, width, height, debug)
	for i := range steps {
		fmt.Printf("%c%d ", steps[i].r, steps[i].i)
	}
	fmt.Println()
	fmt.Println(len(steps))
	/*
		com := NewComputer(ops)
		com.mem[0]=2
		for !com.Halted {
			out, halt, err := com.Iterate(func() int {
				return 0
			}, false, false)
			if err != nil {
				log.Fatal(fmt.Errorf("failure while executing computer: %w", err))
			}
			if halt {
				break
			}

			if debug {
				waitEnter()
			}
		}
	*/

	return 0
}

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "draw debug")
	flag.Parse()

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

	alignment, grid, width, height := FindAlignment(ops, debug)
	dust := NotifyRobots(ops, grid, width, height, debug)
	fmt.Printf("Part 1 - Alighment: %d\n", alignment)
	fmt.Printf("Part 2 - Dust: %d\n", dust)
}
