package main

import "core:fmt"
import "core:os"

parse_input_file :: proc(filepath: string) -> []u8 {
	data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	return data
}

is_uniq :: proc(message: []u8) -> bool {
	for i := 0; i < len(message); i += 1 {
		for j := i + 1; j < len(message); j += 1 {
			if message[i] == message[j] {
				return false
			}
		}
	}
	return true
}

task1 :: proc(datastream: []u8) -> int {
	marker: int
	message: [4]u8
	for char, index in datastream {
		for i := 0; i < 3; i += 1 {
			message[i] = message[i + 1]
		}
		message[3] = char

		if index > 2 {
			if is_uniq(message[:]) {
				marker = index + 1
				break
			}
		}
	}
	return marker
}

task2 :: proc(datastream: []u8) -> int {
	marker: int
	message: [14]u8
	for char, index in datastream {
		for i := 0; i < 13; i += 1 {
			message[i] = message[i + 1]
		}
		message[13] = char

		if index > 12 {
			if is_uniq(message[:]) {
				marker = index + 1
				break
			}
		}
	}
	return marker
}

main :: proc() {
	datastream := parse_input_file("input.txt")
	defer delete(datastream)

	result1 := task1(datastream)
	fmt.printf("Task 1 result: %d\n", result1)

	result2 := task2(datastream)
	fmt.printf("Task 2 result: %d\n", result2)
}
