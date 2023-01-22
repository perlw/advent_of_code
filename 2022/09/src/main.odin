package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:math"
import "core:os"
import "core:slice"

Instruction :: [2]byte

parse_input_file :: proc(filepath: string, buffer: []byte) -> []Instruction {
	data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	buf: bytes.Buffer
	bytes.buffer_init(&buf, data)
	defer bytes.buffer_destroy(&buf)

	index: int
	line: []u8
	err: io.Error
	for ; err != .EOF; line, err = bytes.buffer_read_bytes(&buf, '\n') {
		if len(line) == 0 {
			continue
		}

		buffer[index] = line[0]
		val := line[2] - 48
		if line[3] >= 48 {
			val = (val * 10) + line[3] - 48
		}
		buffer[index + 1] = val
		index += 2
	}

	return transmute([]Instruction)buffer[:index / 2]
}

Grid :: struct {
	data: []byte,
	w:    int,
	h:    int,
}

pretty_print :: proc {
	pretty_print_basic,
	pretty_print_with_tail,
}

pretty_print_basic :: proc(grid: ^Grid) {
	for char, i in grid.data {
		if i > 0 && i % grid.w == 0 {
			fmt.printf("\n")
		}
		fmt.printf("%c", char)
	}
	fmt.printf("\n")
}

pretty_print_with_tail :: proc(grid: ^Grid, head: int, tail: []int) {
	for char, i in grid.data {
		if i > 0 && i % grid.w == 0 {
			if math.floor_div(i, grid.w) == 1 {
				fmt.printf("H:%v", []int{head % grid.w, head / grid.w})
			} else if math.floor_div(i, grid.w) - 2 < len(tail) {
				ti := math.floor_div(i, grid.w) - 2
				fmt.printf("%d:%v", ti + 1, []int{tail[ti] % grid.w, tail[ti] / grid.w})
			}
			fmt.printf("\n")
		}
		fmt.printf("%c", char)
	}
	fmt.printf("\n\n")
}

is_touching :: proc(head, tail: int, playfield: ^Grid) -> bool {
	return(
		(head >= tail - playfield.w - 1 && head <= tail - playfield.w + 1) ||
		(head >= tail - 1 && head <= tail + 1) ||
		(head >= tail + playfield.w - 1 && head <= tail + playfield.w + 1) \
	)
}

get_new_tail :: proc(head, tail: int, playfield: ^Grid) -> int {
	cardinals := []int{-playfield.w, playfield.w, -1, 1}
	diagonals := []int{-playfield.w, +playfield.w, -1, 1}
	for c in cardinals {
		if head == tail + (c * 2) {
			return tail + c
		}
	}
	for d in diagonals {
		if head == tail + (d * 2) - 1 {
			return tail + d - 1
		}
		if head == tail + (d * 2) + 1 {
			return tail + d + 1
		}
		if head == tail + (d * 2) - playfield.w {
			return tail + d - playfield.w
		}
		if head == tail + (d * 2) + playfield.w {
			return tail + d + playfield.w
		}

		if head == tail + (d * 2) - 2 {
			return tail + d - 1
		}
		if head == tail + (d * 2) + 2 {
			return tail + d + 1
		}
	}

	fmt.printf("head %d %v\n", head, [2]int{head % playfield.w, head / playfield.w})
	fmt.printf("tail %d %v\n", tail, [2]int{tail % playfield.w, tail / playfield.w})
	panic("unreachable!")
}

task1 :: proc(instructions: []Instruction, playfield: ^Grid, debug := false) -> int {
	result := 1

	slice.fill(playfield.data, '.')

	head := (playfield.w * (playfield.h / 2)) + (playfield.w / 2)
	tail := head
	for instruction in instructions {
		step: int
		step = -playfield.w if instruction[0] == 'U' else step
		step = playfield.w if instruction[0] == 'D' else step
		step = -1 if instruction[0] == 'L' else step
		step = 1 if instruction[0] == 'R' else step
		for _ in 0 ..< instruction[1] {
			head += step
			if !is_touching(head, tail, playfield) {
				tail = get_new_tail(head, tail, playfield)
				if playfield.data[tail] != 'T' {
					result += 1
				}
			}

			playfield.data[tail] = 'T'
		}
	}
	if debug {
		pretty_print(playfield)
	}

	return result
}

task2 :: proc(instructions: []Instruction, playfield: ^Grid, debug := false) -> int {
	result := 1

	slice.fill(playfield.data, '.')

	head := (playfield.w * (playfield.h / 2)) + (playfield.w / 2)
	tail := [9]int{head, head, head, head, head, head, head, head, head}
	for instruction in instructions {
		step: int
		step = -playfield.w if instruction[0] == 'U' else step
		step = playfield.w if instruction[0] == 'D' else step
		step = -1 if instruction[0] == 'L' else step
		step = 1 if instruction[0] == 'R' else step

		for _ in 0 ..< instruction[1] {
			head += step

			prev_index := head
			for _, i in tail {
				tail_byte := byte('1' + i)

				if !is_touching(prev_index, tail[i], playfield) {
					tail[i] = get_new_tail(prev_index, tail[i], playfield)
					if i + 1 == 9 && playfield.data[tail[i]] != '9' {
						result += 1
					}
				}

				if i == len(tail) - 1 {
					playfield.data[tail[i]] = tail_byte
				}

				prev_index = tail[i]
			}
		}
	}
	if debug {
		pretty_print(playfield)
	}

	return result
}

main :: proc() {
	buffer: [16383]byte
	instructions := parse_input_file("input.txt", buffer[:])

	width :: 1024
	playfield_data: [width * width]byte
	playfield := Grid {
		data = playfield_data[:],
		w    = width,
		h    = width,
	}

	result1 := task1(instructions, &playfield)
	fmt.printf("Task 1 result: %d\n", result1)

	result2 := task2(instructions, &playfield)
	fmt.printf("Task 2 result: %d\n", result2)
}
