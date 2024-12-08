package main

import "core:log"
import "core:os"
import "core:slice"
import "core:strings"

Input :: struct {
	grid: [][]u8,
}

Result1 :: distinct int
Result2 :: distinct int


// --- Input --- //
print_input :: proc(input: ^Input) {
	for line in input.grid {
		log.infof("%s", line)
	}
}

parse_input_file :: proc(filepath: string) -> Input {
	input: Input

	raw_data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}
	defer delete(raw_data)

	lines := strings.split_lines(string(raw_data))
	defer delete(lines)

	input.grid = make([][]u8, len(lines) - 1)
	i: int
	for line in lines {
		if len(line) == 0 {
			continue
		}
		input.grid[i] = slice.clone(transmute([]u8)line[:])
		i += 1
	}

	return input
}

free_input :: proc(input: ^Input) {
	for _, i in input.grid {
		delete(input.grid[i])
	}
	delete(input.grid)
}
// --- Input --- //


// --- Helpers --- //
Direction :: enum {
	Up,
	Down,
	Left,
	Right,
	UpLeft,
	UpRight,
	DownLeft,
	DownRight,
}

print_grid :: proc(grid: [][]u8) {
	for line in grid {
		log.debugf("%s", line)
	}
}
// --- Helpers --- //


// --- Task 1 --- //
run_task1 :: proc(input: ^Input, debug: bool) -> Result1 {
	result: Result1

	if debug {
		print_input(input)
	}

	grid := make([][]u8, len(input.grid))
	for _, i in grid {
		grid[i] = make([]u8, len(input.grid[i]))
	}
	defer {
		for _, i in grid {
			delete(grid[i])
		}
		delete(grid)
	}

	chars := []u8{'X', 'M', 'A', 'S'}
	height := len(grid)
	for y in 0 ..< height {
		width := len(grid[y])
		for x in 0 ..< width {
			if input.grid[y][x] != 'X' {
				grid[y][x] = '.'
				continue
			}

			matches: [Direction]int
			for c, i in chars {
				if y - i >= 0 {
					if c == input.grid[y - i][x] {
						matches[.Up] += 1
					}
				}
				if y + i < height {
					if c == input.grid[y + i][x] {
						matches[.Down] += 1
					}
				}
				if x - i >= 0 {
					if c == input.grid[y][x - i] {
						matches[.Left] += 1
					}
				}
				if x + i < width {
					if c == input.grid[y][x + i] {
						matches[.Right] += 1
					}
				}
				if y - i >= 0 && x - i >= 0 {
					if c == input.grid[y - i][x - i] {
						matches[.UpLeft] += 1
					}
				}
				if y - i >= 0 && x + i < width {
					if c == input.grid[y - i][x + i] {
						matches[.UpRight] += 1
					}
				}
				if y + i < height && x - i >= 0 {
					if c == input.grid[y + i][x - i] {
						matches[.DownLeft] += 1
					}
				}
				if y + i < height && x + i < width {
					if c == input.grid[y + i][x + i] {
						matches[.DownRight] += 1
					}
				}
			}

			log.debugf("%v", matches)

			grid[y][x] = 'F'
			for m in matches {
				if m == len(chars) {
					grid[y][x] = 'T'
					result += 1
				}
			}
		}
	}

	print_grid(grid)

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

	grid := make([][]u8, len(input.grid))
	for _, i in grid {
		grid[i] = make([]u8, len(input.grid[i]))
	}
	defer {
		for _, i in grid {
			delete(grid[i])
		}
		delete(grid)
	}

	height := len(grid)
	for y in 1 ..< height - 1 {
		width := len(grid[y])
		for x in 1 ..< width - 1 {
			if input.grid[y][x] != 'A' {
				grid[y][x] = ' '
				continue
			}

			matches: int
			if 'M' == input.grid[y - 1][x - 1] && 'S' == input.grid[y + 1][x + 1] {
				matches += 1
			}
			if 'S' == input.grid[y - 1][x - 1] && 'M' == input.grid[y + 1][x + 1] {
				matches += 1
			}
			if 'M' == input.grid[y + 1][x - 1] && 'S' == input.grid[y - 1][x + 1] {
				matches += 1
			}
			if 'S' == input.grid[y + 1][x - 1] && 'M' == input.grid[y - 1][x + 1] {
				matches += 1
			}

			grid[y][x] = 'T' if matches == 2 else 'F'
			if matches == 2 {
				result += 1
			}
		}
	}

	print_grid(grid)

	return result
}

print_result2 :: proc(result: ^Result2) {
	log.infof("Task 2: %d", result^)
}
// --- Task 2 --- //
