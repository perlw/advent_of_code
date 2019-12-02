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
	OpAdd OpType = 1
	OpMul        = 2
	OpEnd        = 99
)

func ExecuteProgram(ops []int) ([]int, error) {
	for i := 0; i < len(ops); i += 4 {
		op := OpType(ops[i])
		if op == OpEnd {
			break
		}

		if i+3 >= len(ops) {
			return nil, fmt.Errorf("broken input, could not index: %d", i+3)
		}
		a := ops[i+1]
		b := ops[i+2]
		out := ops[i+3]
		if a >= len(ops) {
			return nil, fmt.Errorf("attempt to index invalid out position: %d", a)
		}
		if b >= len(ops) {
			return nil, fmt.Errorf("attempt to index invalid out position: %d", b)
		}
		if out >= len(ops) {
			return nil, fmt.Errorf("attempt to index invalid out position: %d", out)
		}

		switch op {
		case OpAdd:
			ops[out] = ops[a] + ops[b]
		case OpMul:
			ops[out] = ops[a] * ops[b]
		default:
			return nil, fmt.Errorf("invalid op: %d", op)
		}
	}
	return ops, nil
}

func Run(ops []int, noun, verb int) (int, error) {
	cOps := make([]int, len(ops))
	copy(cOps, ops)
	cOps[1] = noun
	cOps[2] = verb

	result, err := ExecuteProgram(cOps)
	if err != nil {
		return 0, fmt.Errorf("could not execute program: %w", err)
	}

	return result[0], nil
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
	result, err := Run(ops, 12, 2)
	if err != nil {
		log.Fatal(fmt.Errorf("could not run part 1: %w", err))
	}
	fmt.Printf("Part 1, %d\n", result)

	// Part 2
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			result, err := Run(ops, noun, verb)
			if err != nil {
				log.Fatal(fmt.Errorf("could not run part 2: %w", err))
			}
			if result == 19690720 {
				fmt.Printf("Part 2, %d\n", (noun*100)+verb)
				return
			}
		}
	}
}
