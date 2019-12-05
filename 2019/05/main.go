package main

import (
	"fmt"
	"io"
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
	OpEnd = 99
)

type ModeType int

const (
	ModePosition  ModeType = 0
	ModeImmediate          = 1
)

func getArgPos(ops []int, pos int, mode ModeType) (int, error) {
	if pos >= len(ops) {
		return -1, fmt.Errorf("invalid index, out of range: %d", pos)
	}
	if mode == ModePosition {
		if pos >= len(ops) {
			return -1, fmt.Errorf("invalid immediate index, out of range: %d", pos)
		}
		return ops[pos], nil
	}
	return pos, nil
}

func ExecuteProgram(ops []int, input func() int, output io.Writer, debug bool) ([]int, error) {
	log := ioutil.Discard
	if debug {
		log = os.Stdout
	}
	var pc int
	for pc < len(ops) {
		op := OpType(ops[pc])
		if op == OpEnd {
			break
		}

		aMode := ModePosition
		bMode := ModePosition
		if op > 99 {
			aMode = ModeType((op / 100) % 10)
			bMode = ModeType((op / 1000) % 10)
			op %= 100
		}

		switch op {
		case OpAdd:
			a, err := getArgPos(ops, pc+1, aMode)
			if err != nil {
				return nil, err
			}
			b, err := getArgPos(ops, pc+2, bMode)
			if err != nil {
				return nil, err
			}
			o, err := getArgPos(ops, pc+3, ModePosition)
			if err != nil {
				return nil, err
			}
			ops[o] = ops[a] + ops[b]
			fmt.Fprintf(log, "Setting %d @%d\n", ops[o], o)
			pc += 4
		case OpMul:
			a, err := getArgPos(ops, pc+1, aMode)
			if err != nil {
				return nil, err
			}
			b, err := getArgPos(ops, pc+2, bMode)
			if err != nil {
				return nil, err
			}
			o, err := getArgPos(ops, pc+3, ModePosition)
			if err != nil {
				return nil, err
			}
			ops[o] = ops[a] * ops[b]
			fmt.Fprintf(log, "Setting %d @%d\n", ops[o], o)
			pc += 4
		case OpInput:
			a, err := getArgPos(ops, pc+1, ModePosition)
			if err != nil {
				return nil, err
			}
			ops[a] = input()
			fmt.Fprintf(log, "Inputting: %d @%d\n", ops[a], a)
			pc += 2
		case OpOutput:
			a, err := getArgPos(ops, pc+1, aMode)
			if err != nil {
				return nil, err
			}
			if output != nil {
				fmt.Fprintf(output, "%d", ops[a])
			}
			pc += 2
		case OpJumpTrue:
			a, err := getArgPos(ops, pc+1, aMode)
			if err != nil {
				return nil, err
			}
			b, err := getArgPos(ops, pc+2, bMode)
			if err != nil {
				return nil, err
			}
			if ops[a] != 0 {
				pc = ops[b]
				fmt.Fprintf(log, "Jump to %d\n", pc)
			} else {
				pc += 3
			}
		case OpJumpFalse:
			a, err := getArgPos(ops, pc+1, aMode)
			if err != nil {
				return nil, err
			}
			b, err := getArgPos(ops, pc+2, bMode)
			if err != nil {
				return nil, err
			}
			if ops[a] == 0 {
				pc = ops[b]
				fmt.Fprintf(log, "Jump to %d\n", pc)
			} else {
				pc += 3
			}
		case OpLessThan:
			a, err := getArgPos(ops, pc+1, aMode)
			if err != nil {
				return nil, err
			}
			b, err := getArgPos(ops, pc+2, bMode)
			if err != nil {
				return nil, err
			}
			o, err := getArgPos(ops, pc+3, ModePosition)
			if err != nil {
				return nil, err
			}
			if ops[a] < ops[b] {
				ops[o] = 1
			} else {
				ops[o] = 0
			}
			pc += 4
		case OpEquals:
			a, err := getArgPos(ops, pc+1, aMode)
			if err != nil {
				return nil, err
			}
			b, err := getArgPos(ops, pc+2, bMode)
			if err != nil {
				return nil, err
			}
			o, err := getArgPos(ops, pc+3, ModePosition)
			if err != nil {
				return nil, err
			}
			if ops[a] == ops[b] {
				ops[o] = 1
			} else {
				ops[o] = 0
			}
			pc += 4
		default:
			return nil, fmt.Errorf("invalid op: %d", op)
		}
	}
	return ops, nil
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

	prog := make([]int, len(ops))
	copy(prog, ops)
	// Part 1
	fmt.Println("Part 1...")
	_, err = ExecuteProgram(prog, func() int {
		return 1
	}, os.Stdout, false)
	if err != nil {
		log.Fatal(fmt.Errorf("could not run part 1: %w", err))
	}
	fmt.Println()

	copy(prog, ops)
	fmt.Println("\nPart 2...")
	_, err = ExecuteProgram(prog, func() int {
		return 5
	}, os.Stdout, false)
	if err != nil {
		log.Fatal(fmt.Errorf("could not run part 2: %w", err))
	}
	fmt.Println()
}
