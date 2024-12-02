package main

import "core:log"
import "core:os"
import "core:slice"
import "core:strconv"
import "core:strings"

Input :: struct {
	left:  [dynamic]int,
	right: [dynamic]int,
}

Result1 :: distinct int
Result2 :: distinct int


// --- Input --- //
print_input :: proc(input: ^Input) {
	log.infof("%v", input)
}

parse_input_file :: proc(filepath: string) -> Input {
	input: Input

	raw_data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}
	defer delete(raw_data)

	input.left = make([dynamic]int, 0, 10)
	input.right = make([dynamic]int, 0, 10)

	lines := strings.split_lines(string(raw_data))
	defer delete(lines)

	for line in lines {
		left_end := strings.index_byte(line, ' ')
		right_start := strings.last_index_byte(line, ' ')
		if left_end == -1 || right_start == -1 {
			continue
		}

		left := strconv.atoi(line[:left_end])
		right := strconv.atoi(line[right_start + 1:])

		append(&input.left, left)
		append(&input.right, right)
	}

	return input
}

free_input :: proc(input: ^Input) {
	delete(input.left)
	delete(input.right)
}
// --- Input --- //


// --- Task 1 --- //
run_task1 :: proc(input: ^Input, debug: bool) -> Result1 {
	result: Result1

	if debug {
		print_input(input)
	}

	left_slice := slice.clone(input.left[:])
	right_slice := slice.clone(input.right[:])
	defer delete(left_slice)
	defer delete(right_slice)

	slice.sort(left_slice)
	slice.sort(right_slice)

	distance: int
	for i in 0 ..< len(left_slice) {
		l, r := left_slice[i], right_slice[i]
		if l < r {
			l, r = r, l
		}
		distance += l - r
	}
	result = Result1(distance)

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

	similarity: int
	for i in 0 ..< len(input.left) {
		left := input.left[i]
		count: int
		for j in 0 ..< len(input.right) {
			count += 1 if input.right[j] == left else 0
		}
		similarity += left * count
	}
	result = Result2(similarity)

	return result
}

print_result2 :: proc(result: ^Result2) {
	log.infof("Task 2: %d", result^)
}
// --- Task 2 --- //
