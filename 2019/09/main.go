package main

import (
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

func main() {
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

	// Part 1
	com := NewComputer(ops)
	fmt.Printf("Part 1:\n")
	for !com.Halted {
		out, halt, err := com.Iterate(func() int {
			return 1
		}, false, false)
		if err != nil {
			log.Fatal(fmt.Errorf("failure while executing computer: %w", err))
		}
		if !halt {
			fmt.Printf("%d\n", out)
		}
	}

	// Part 2
	com = NewComputer(ops)
	fmt.Printf("Part 2:\n")
	for !com.Halted {
		out, halt, err := com.Iterate(func() int {
			return 2
		}, false, false)
		if err != nil {
			log.Fatal(fmt.Errorf("failure while executing computer: %w", err))
		}
		if !halt {
			fmt.Printf("%d\n", out)
		}
	}
}
