package main

import "core:fmt"
import "core:log"
import "core:os"

Input :: struct {}

Result1 :: distinct int
Result2 :: distinct int


// --- Input --- //
print_input :: proc(input: ^Input) {
	fmt.printf("%v\n", input)
}

parse_input_file :: proc(filepath: string) -> Input {
	input: Input

	raw_data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}
	defer delete(raw_data)

	return input
}

free_input :: proc(input: ^Input) {
}
// --- Input --- //


// --- Task 1 --- //
run_task1 :: proc(input: ^Input, debug: bool) -> Result1 {
	result: Result1

	if debug {
		print_input(input)
	}

	return result
}

print_result1 :: proc(result: ^Result1) {
	log.infof("Task 1: %d", result^)
}
// --- Task 1 --- //


// --- Task 2 --- //
run_task2 :: proc(input: ^Input, debug: bool) -> Result2 {
	result: Result2

	if debug {
		print_input(input)
	}

	return result
}

print_result2 :: proc(result: ^Result2) {
	log.infof("Task 2: %d", result^)
}
// --- Task 2 --- //
