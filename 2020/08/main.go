package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Instr ...
type Instr string

// ...
const (
	InstrAcc Instr = "acc"
	InstrJmp       = "jmp"
	InstrNop       = "nop"
)

// Op ...
type Op struct {
	Instr Instr
	Val   int
}

func (o Op) String() string {
	return fmt.Sprintf("%s  %+4d", o.Instr, o.Val)
}

// ElfCPU ...
type ElfCPU struct {
	PC  int
	Acc int

	Ops []Op
}

// Reset ...
func (e *ElfCPU) Reset() {
	e.PC = 0
	e.Acc = 0
}

// Step ...
func (e *ElfCPU) Step() bool {
	if e.PC >= len(e.Ops) {
		return false
	}

	op := e.Ops[e.PC]
	switch op.Instr {
	case InstrAcc:
		e.Acc += op.Val
		e.PC++
	case InstrJmp:
		e.PC += op.Val
	case InstrNop:
		e.PC++
	default:
		e.PC++
	}

	return true
}

// runCPU ...
func runCPU(elf *ElfCPU) (int, bool) {
	seenPCs := make(map[int]struct{})

	for {
		if _, ok := seenPCs[elf.PC]; ok {
			return elf.Acc, false
		}
		seenPCs[elf.PC] = struct{}{}

		if !elf.Step() {
			break
		}
	}

	return elf.Acc, true
}

// Task1 ...
func Task1(elf *ElfCPU) int {
	result, _ := runCPU(elf)
	return result
}

// Task2 ...
func Task2(elf *ElfCPU) int {
	type step struct {
		PC int
		Op Op
	}

	seenPCs := make(map[int]struct{})
	seenOps := make([]step, 0, 100)

	for {
		if _, ok := seenPCs[elf.PC]; ok {
			break
		}

		seenPCs[elf.PC] = struct{}{}
		seenOps = append(seenOps, step{
			PC: elf.PC,
			Op: elf.Ops[elf.PC],
		})
		if !elf.Step() {
			break
		}
	}

	ops := make([]Op, len(elf.Ops))
	copy(ops, elf.Ops)

	// Search backwards, change a jmp or nop, test, and start over.
	for i := len(seenOps) - 1; i >= 0; i-- {
		copy(elf.Ops, ops)
		elf.Reset()

		var try bool
		if seenOps[i].Op.Instr == InstrJmp {
			elf.Ops[seenOps[i].PC].Instr = InstrNop
			try = true
		} else if seenOps[i].Op.Instr == InstrNop {
			elf.Ops[seenOps[i].PC].Instr = InstrJmp
			try = true
		}

		if try {
			fmt.Printf("modify: %s -> %s", seenOps[i].Op, elf.Ops[seenOps[i].PC])
			if acc, ok := runCPU(elf); ok {
				fmt.Printf(" FOUND!\n")
				return acc
			}
			fmt.Printf(" no dice\n")
		}
	}

	return -1
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	elf := ElfCPU{
		Ops: make([]Op, 0, 100),
	}
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " ")
		val, _ := strconv.Atoi(parts[1])
		elf.Ops = append(elf.Ops, Op{
			Instr: Instr(parts[0]),
			Val:   val,
		})
	}
	file.Close()

	elf.Reset()
	result := Task1(&elf)
	fmt.Printf("Task 1: %d\n", result)

	elf.Reset()
	result = Task2(&elf)
	fmt.Printf("Task 2: %d\n", result)
}
