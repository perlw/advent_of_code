package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:os"
import "core:slice"
import "core:strconv"

State :: struct {
	columns: [dynamic][dynamic]u8,
}

Instruction :: struct {
	count: int,
	from:  int,
	to:    int,
}

parse_input_file :: proc(filepath: string) -> (State, [dynamic]Instruction) {
	data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	buf: bytes.Buffer
	bytes.buffer_init(&buf, data)
	defer bytes.buffer_destroy(&buf)

	state := State {
		columns = make([dynamic][dynamic]u8, 0, 100),
	}
	instructions := make([dynamic]Instruction, 0, 100)

	line: []u8
	err: io.Error
	for ; err != .EOF; line, err = bytes.buffer_read_bytes(&buf, '\n') {
		if len(line) == 0 {
			continue
		}

		if line[0] == '\n' {
			break
		}
		if bytes.index_byte(line, '[') == -1 {
			continue
		}

		line = line[:len(line) - 1]
		for i := 0; i < len(line); i += 4 {
			position := i / 4
			if len(state.columns) < position + 1 {
				append(&state.columns, make([dynamic]u8, 0, 10))
			}
			if line[i + 1] == ' ' {
				continue
			}
			append(&state.columns[position], line[i + 1])
		}
	}
	for column in state.columns {
		for i := 0; i < len(column) / 2; i += 1 {
			column[i], column[len(column) - i - 1] = column[len(column) - i - 1], column[i]
		}
	}

	for ; err != .EOF; line, err = bytes.buffer_read_bytes(&buf, '\n') {
		if len(line) == 0 || line[0] == '\n' {
			continue
		}

		numbers: [3]int
		num_read: int
		line = line[:len(line) - 1]
		i: int
		for {
			start := bytes.index_byte(line[i:], ' ') + i
			end := bytes.index_byte(line[start + 1:], ' ') + start + 1
			if start < 0 || end < 0 {
				break
			}
			if start == end {
				end = len(line)
			}
			numbers[num_read] = strconv.atoi(string(line[start + 1:end]))
			num_read += 1

			i = end + 1
			if i >= len(line) {
				break
			}
		}
		append(&instructions, Instruction{count = numbers[0], from = numbers[1], to = numbers[2]})
	}

	return state, instructions
}

free_input_data :: proc(state: State, instructions: [dynamic]Instruction) {
	free_state(state)
	delete(instructions)
}

clone_state :: proc(state: State) -> State {
	state_clone := State {
		columns = make([dynamic][dynamic]u8, 0, 100),
	}
	for column in state.columns {
		append(&state_clone.columns, slice.clone_to_dynamic(column[:]))
	}
	return state_clone
}

free_state :: proc(state: State) {
	for column in state.columns {
		delete(column)
	}
	delete(state.columns)
}

task1 :: proc(state: State, instructions: []Instruction) -> [dynamic]u8 {
	state_clone := clone_state(state)
	defer free_state(state_clone)

	for instruction in instructions {
		for _ in 0 ..< instruction.count {
			from_column := state_clone.columns[instruction.from - 1]

			moved_char := from_column[len(from_column) - 1]
			state_clone.columns[instruction.from - 1] = slice.clone_to_dynamic(
				from_column[:len(from_column) - 1],
			)
			delete(from_column)

			append(&state_clone.columns[instruction.to - 1], moved_char)
		}
	}

	result := make([dynamic]u8, 0, 10)
	for column in state_clone.columns {
		append(&result, column[len(column) - 1])
	}
	return result
}

task2 :: proc(state: State, instructions: []Instruction) -> [dynamic]u8 {
	state_clone := clone_state(state)
	defer free_state(state_clone)

	for instruction in instructions {
		from_column := state_clone.columns[instruction.from - 1]
		for i in 0 ..< instruction.count {
			ri := instruction.count - 1 - i
			moved_char := from_column[len(from_column) - 1 - ri]
			append(&state_clone.columns[instruction.to - 1], moved_char)
		}
		state_clone.columns[instruction.from - 1] = slice.clone_to_dynamic(
			from_column[:len(from_column) - instruction.count],
		)
		delete(from_column)
	}

	result := make([dynamic]u8, 0, 10)
	for column in state_clone.columns {
		append(&result, column[len(column) - 1])
	}
	return result
}

main :: proc() {
	state, instructions := parse_input_file("input.txt")
	defer free_input_data(state, instructions)

	result1 := task1(state, instructions[:])
	fmt.printf("Task 1 result: %s\n", result1)

	result2 := task2(state, instructions[:])
	fmt.printf("Task 2 result: %s\n", result2)
}
