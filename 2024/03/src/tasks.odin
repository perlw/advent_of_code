package main

import "core:log"
import "core:os"
import "core:strconv"
import "core:strings"

Input :: struct {
	muls:            [dynamic][2]int,
	controlled_muls: [dynamic][2]int,
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

	input.muls = make([dynamic][2]int, 0, 10)
	input.controlled_muls = make([dynamic][2]int, 0, 10)

	should_do := true
	string_data := string(raw_data)
	for line in strings.split_lines_iterator(&string_data) {
		for i in 0 ..< len(line) {
			if i + 4 < len(line) && line[i:i + 4] == "mul(" {
				current := line[i + 4:]
				mid := strings.index(current, ",")
				end := strings.index(current, ")")
				if mid == -1 || end == -1 || mid >= end {
					continue
				}

				a, ok_a := strconv.parse_int(current[:mid])
				b, ok_b := strconv.parse_int(current[mid + 1:end])
				if !ok_a || !ok_b {
					continue
				}

				append(&input.muls, [2]int{a, b})
				if should_do {
					log.debugf("should %s", current[:end])
					append(&input.controlled_muls, [2]int{a, b})
				} else {
					log.debugf("shouldn't %s", current[:end])
				}
			} else if i + 4 < len(line) && line[i:i + 4] == "do()" {
				should_do = true
			} else if i + 7 < len(line) && line[i:i + 7] == "don't()" {
				should_do = false
			}
		}
	}

	return input
}

free_input :: proc(input: ^Input) {
	delete(input.controlled_muls)
	delete(input.muls)
}
// --- Input --- //


// --- Helpers --- //
// --- Helpers --- //


// --- Task 1 --- //
run_task1 :: proc(input: ^Input, debug: bool) -> Result1 {
	result: Result1

	if debug {
		print_input(input)
	}

	for m in input.muls {
		result += Result1(m[0] * m[1])
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

	for m in input.controlled_muls {
		result += Result2(m[0] * m[1])
	}

	return result
}

print_result2 :: proc(result: ^Result2) {
	log.infof("Task 2: %d", result^)
}
// --- Task 2 --- //
