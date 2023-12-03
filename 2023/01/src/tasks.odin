package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:os"
import "core:strings"

Input :: distinct struct {
	raw_data: []byte,
	lines:    [dynamic]string,
}

Result1 :: distinct int
Result2 :: distinct int


// --- Input --- //
parse_input_file :: proc(filepath: string) -> Input {
	input: Input

	ok: bool
	input.raw_data, ok = os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	buf: bytes.Buffer
	bytes.buffer_init(&buf, input.raw_data)

	input.lines = make([dynamic]string, 0, 10)

	line: string
	err: io.Error
	for ; err != .EOF; line, err = bytes.buffer_read_string(&buf, '\n') {
		if line == "" || line == "\n" {
			continue
		}

		append(&input.lines, line[:len(line) - 1])
	}

	return input
}

free_input :: proc(input: ^Input) {
	delete(input.lines)
	delete(input.raw_data)
}
// --- Input --- //


// --- Task 1 --- //
run_task1 :: proc(input: ^Input, debug: bool) -> Result1 {
	result: Result1

	for line, i in input.lines {
		first, second := -1, -1
		for char in line {
			if char >= '0' && char <= '9' {
				if first == -1 {
					first = int(char - '0')
					second = first
				} else {
					second = int(char - '0')
				}
			}
		}
		r := Result1((first * 10) + second)
		if debug {
			fmt.printf("%d: %d\n", i, r)
		}
		result += r
	}

	return result
}

print_result1 :: proc(result: ^Result1) {
	fmt.printf("Task 1: %d\n", result^)
}
// --- Task 1 --- //


// --- Task 2 --- //
digits := []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

run_task2 :: proc(input: ^Input, debug: bool) -> Result2 {
	result: Result2

	for line, i in input.lines {
		if debug {
			fmt.printf("%d: %s\n", i, line)
		}

		first_index, first_value := -1, -1
		second_index, second_value := -1, -1
		for char, ci in line {
			if char >= '0' && char <= '9' {
				if first_index == -1 {
					first_index = ci
					first_value = int(char - '0')
					second_index = first_index
					second_value = first_value
				} else {
					second_index = ci
					second_value = int(char - '0')
				}
			}
		}
		if debug {
			fmt.printf("%d: %d, %d: %d\n", first_index, first_value, second_index, second_value)
		}
		for digit, digit_index in digits {
			if di := strings.index(line, digit); di > -1 && (first_index < 0 || di < first_index) {
				first_index = di
				first_value = digit_index + 1
				if debug {
					fmt.printf("first match: %s %d -> %d\n", digit, digit_index, i)
				}
			}
			if di := strings.last_index(line, digit); di > -1 && (second_index < 0 || di > second_index) {
				second_index = di
				second_value = digit_index + 1
				if debug {
					fmt.printf("second match: %s %d -> %d\n", digit, digit_index, i)
				}
			}
		}

		r := Result2((first_value * 10) + second_value)
		if debug {
			fmt.printf("%d: %d\n", i, r)
		}
		result += r
	}

	return result
}

print_result2 :: proc(result: ^Result2) {
	fmt.printf("Task 2: %d\n", result^)
}
// --- Task 2 --- //
