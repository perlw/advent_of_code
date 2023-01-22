package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:os"
import "core:strconv"
import "core:strings"

Instruction :: struct {
	Op: OpCode,
	X:  int,
}

OpCode :: enum {
	Noop,
	AddX,
}

OpCycles := []int {
	OpCode.Noop = 1,
	OpCode.AddX = 2,
}

// NOTE: Have to free result.
parse_input_file :: proc(filepath: string) -> [dynamic]Instruction {
	data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	buf: bytes.Buffer
	bytes.buffer_init(&buf, data)
	defer bytes.buffer_destroy(&buf)

	instructions := make([dynamic]Instruction, 0, 10)

	line: string
	err: io.Error
	for ; err != .EOF; line, err = bytes.buffer_read_string(&buf, '\n') {
		if line == "" || line == "\n" {
			continue
		}

		line = line[:len(line) - 1]
		instr := strings.split(line, " ")
		defer delete(instr)

		switch instr[0] {
		case "noop":
			append(&instructions, Instruction{Op = .Noop})
		case "addx":
			append(&instructions, Instruction{Op = .AddX, X = strconv.atoi(instr[1])})
		case:
			panic("unknown op")
		}
	}

	return instructions
}

task1 :: proc(instructions: []Instruction, debug := false) -> int {
	result := 0

	x := 1
	cycle: int
	for instruction in instructions {
		for i in 0 ..< OpCycles[instruction.Op] {
			cycle += 1
			if cycle >= 20 && (cycle - 20) % 40 == 0 {
				result += (cycle * x)
			}
		}
		switch instruction.Op {
		case .Noop:
		case .AddX:
			x += instruction.X
		}
		if debug {
			fmt.printf("[%d/%d] -> %v\n", cycle, x, instruction)
		}
	}

	return result
}

task2 :: proc(instructions: []Instruction, debug := false) {
	x := 1
	hscan := 0
	for instruction in instructions {
		for i in 0 ..< OpCycles[instruction.Op] {
			if x >= hscan - 1 && x <= hscan + 1 {
				fmt.printf("#")
			} else {
				fmt.printf(".")
			}

			hscan = (hscan + 1) % 40
			if hscan == 0 {
				fmt.printf("\n")
			}
		}

		switch instruction.Op {
		case .Noop:
		case .AddX:
			x += instruction.X
		}
	}
}

main :: proc() {
	instructions := parse_input_file("input.txt")
	defer delete(instructions)

	result1 := task1(instructions[:])
	fmt.printf("Task 1 result: %d\n", result1)

	task2(instructions[:])
}
