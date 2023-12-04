package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:os"

Input :: struct {
	width: int,
	grid:  []rune,
}

Result1 :: distinct int
Result2 :: distinct int


// --- Input --- //
parse_input_file :: proc(filepath: string) -> Input {
	input: Input

	raw_data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	buf: bytes.Buffer
	bytes.buffer_init(&buf, raw_data)
	defer bytes.buffer_destroy(&buf)

	current: int
	line: string
	err: io.Error
	for ; err != .EOF; line, err = bytes.buffer_read_string(&buf, '\n') {
		if line == "" || line == "\n" {
			continue
		}

		line = line[:len(line) - 1]

		if input.width == 0 {
			input.width = len(line)
			input.grid = make([]rune, input.width * input.width)
		}

		for c in line {
			input.grid[current] = c
			current += 1
		}
	}

	return input
}

free_input :: proc(input: ^Input) {
	delete(input.grid)
}
// --- Input --- //


// --- Task 1 --- //
read_number :: proc(line: []rune) -> (num, steps: int) {
	for i := 0; i < len(line); i += 1 {
		r := line[i]
		if r < '0' || r > '9' {
			break
		}
		num *= 10
		num += int(r - '0')
		steps += 1
	}
	return
}

print_input :: proc(input: ^Input) {
	for r, i in input.grid {
		if i % input.width == 0 {
			fmt.printf("\n")
		}
		fmt.printf("%c", r)
	}
	fmt.printf("\n")
}

run_task1 :: proc(input: ^Input, debug: bool) -> Result1 {
	result: Result1

	if debug {
		print_input(input)
	}

	for i := 0; i < len(input.grid); i += 1 {
		r := input.grid[i]
		if r < '0' || r > '9' {
			continue
		}

		start_i := i
		num, steps := read_number(input.grid[i:])
		i += steps

		x := start_i % input.width
		y := start_i / input.width
		start_x := x - 1 if x > 1 else 0
		start_y := y - 1 if y > 1 else 0
		end_x := x + steps if x + steps < input.width - 1 else input.width - 1
		end_y := y + 1 if y < input.width - 1 else input.width - 1

		if debug {
			fmt.printf(
				"number: %d (%d x %d -> %d x %d -> %d x %d) -> %d\n",
				start_i,
				start_x,
				start_y,
				x,
				y,
				end_x,
				end_y,
				num,
			)
		}

		has_part: bool
		loop: for yy := start_y; yy <= end_y; yy += 1 {
			for xx := start_x; xx <= end_x; xx += 1 {
				check := input.grid[(yy * input.width) + xx]
				if !has_part {
					has_part = (check < '0' || check > '9') && check != '.'
				}
				if !debug && has_part {
					break loop
				}
				if debug {
					fmt.printf("%c", check)
				}
			}
			if debug {
				fmt.printf("\n")
			}
		}
		if debug {
			fmt.printf("has part: %t\n", has_part)
			fmt.printf("\n")
		}

		if has_part {
			result += Result1(num)
		}
	}

	return result
}

print_result1 :: proc(result: ^Result1) {
	fmt.printf("Task 1: %d\n", result^)
}
// --- Task 1 --- //


// --- Task 2 --- //
run_task2 :: proc(input: ^Input, debug: bool) -> Result2 {
	result: Result2

	if debug {
		print_input(input)
	}

	for i := 0; i < len(input.grid); i += 1 {
		r := input.grid[i]
		if r != '*' {
			continue
		}

		x := i % input.width
		y := i / input.width
		start_x := x - 1 if x > 1 else 0
		start_y := y - 1 if y > 1 else 0
		end_x := x + 1 if x < input.width - 1 else input.width - 1
		end_y := y + 1 if y < input.width - 1 else input.width - 1

		if debug {
			fmt.printf("number: %d (%d x %d -> %d x %d -> %d x %d)\n", i, start_x, start_y, x, y, end_x, end_y)
		}

		count: int
		total := 1
		loop: for yy := start_y; yy <= end_y; yy += 1 {
			for xx := start_x; xx <= end_x; xx += 1 {
				grid_i := (yy * input.width) + xx
				check := input.grid[grid_i]
				if check >= '0' && check <= '9' {
					left := xx
					for ; left >= 0; left -= 1 {
						check_i := (yy * input.width) + left
						if input.grid[check_i] < '0' || input.grid[check_i] > '9' {
							break
						}
					}
					left += 1
					num, steps := read_number(input.grid[(yy * input.width) + left:])
					xx = left + steps

					if debug {
						fmt.printf("valid: %d\n", num)
					}
					count += 1
					total *= num
				}
			}
		}
		if debug {
			fmt.printf("ratio: %d, count: %d\n", total, count)
		}

		if count == 2 {
			result += Result2(total)
		}
	}

	return result
}

print_result2 :: proc(result: ^Result2) {
	fmt.printf("Task 2: %d\n", result^)
}
// --- Task 2 --- //
