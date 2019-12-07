package main

import (
	"bytes"
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

func RunCircuit(ops []int, inputs []int) (int, []int, error) {
	prog := make([]int, len(ops))
	input := 0
	outputs := []int{}
	for _, i := range inputs {
		copy(prog, ops)

		ii := i
		var b bytes.Buffer
		_, err := ExecuteProgram(prog, func() int {
			j := ii
			ii = input
			return j
		}, io.Writer(&b), false)
		if err != nil {
			return -1, nil, fmt.Errorf("error while executing: %w", err)
		}
		input, _ = strconv.Atoi(b.String())
		outputs = append(outputs, input)
	}

	return input, outputs, nil
}

func verifyInput(input []int) bool {
	seen := make(map[int]bool)
	for _, i := range input {
		if _, ok := seen[i]; ok {
			return false
		}
		seen[i] = true
	}
	return true
}

func FindPhaseSettings(ops []int) (int, error) {
	result := 0
	for i := 0; i <= 4; i++ {
		for j := 0; j <= 4; j++ {
			for k := 0; k <= 4; k++ {
				for l := 0; l <= 4; l++ {
					for m := 0; m <= 4; m++ {
						input := []int{i, j, k, l, m}
						if !verifyInput(input) {
							continue
						}

						amp, _, err := RunCircuit(ops, input)
						if err != nil {
							return 0, fmt.Errorf("could not run phase check: %w", err)
						}
						if amp > result {
							result = amp
						}
					}
				}
			}
		}
	}

	return result, nil
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

	fmt.Println(ops)

	// Part 1
	result, err := FindPhaseSettings(ops)
	if err != nil {
		log.Fatal(fmt.Errorf("could not run part 1: %w", err))
	}
	fmt.Printf("Part 1: %d\n", result)
}
