package main

import "core:log"
import "core:os"
import "core:strconv"
import "core:strings"

Input :: struct {
	reports: [dynamic][dynamic]int,
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

	input.reports = make([dynamic][dynamic]int, 0, 10)

	string_data := string(raw_data)
	for line in strings.split_lines_iterator(&string_data) {
		line_slice := line[:]
		values := make([dynamic]int)
		for value in strings.split_by_byte_iterator(&line_slice, ' ') {
			num := strconv.atoi(value)
			append(&values, num)
		}
		append(&input.reports, values)
	}

	return input
}

free_input :: proc(input: ^Input) {
	for i in 0 ..< len(input.reports) {
		delete(input.reports[i])
	}
	delete(input.reports)
}
// --- Input --- //


// --- Helpers --- //
check_levels :: proc(report: []int) -> bool {
	desc := report[0] > report[1]
	for i in 0 ..< len(report) - 1 {
		a := report[i]
		b := report[i + 1]
		dist := abs(a - b)
		if (dist < 1 || dist > 3) || (desc && a < b) || (!desc && a > b) {
			return false
		}
	}
	return true
}

filter_on_index :: proc(report: []int, index: int) -> [dynamic]int {
	result := make([dynamic]int)
	for i in 0 ..< len(report) {
		if i != index {
			append(&result, report[i])
		}
	}
	return result
}
// --- Helpers --- //


// --- Task 1 --- //
run_task1 :: proc(input: ^Input, debug: bool) -> Result1 {
	result: Result1

	if debug {
		print_input(input)
	}

	for report in input.reports {
		if check_levels(report[:]) {
			result += 1
		}
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

	for report in input.reports {
		safe := check_levels(report[:])

		if !safe {
			log.debugf("Not safe, analyzing: %v", report)
			for i in 0 ..< len(report) {
				filtered := filter_on_index(report[:], i)
				defer delete(filtered)
				if check_levels(filtered[:]) {
					log.debugf("Found safe by removing %d at %d", report[i], i)
					safe = true
					break
				}
			}
		}

		if safe {
			result += 1
		}
	}

	return result
}

print_result2 :: proc(result: ^Result2) {
	log.infof("Task 2: %d", result^)
}
// --- Task 2 --- //
