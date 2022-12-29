package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:os"
import "core:strconv"
import "core:strings"

Pair :: struct {
	elf1: [2]int,
	elf2: [2]int,
}

parse_input_file :: proc(filepath: string) -> [dynamic]Pair {
	data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	buf: bytes.Buffer
	bytes.buffer_init(&buf, data)
	defer bytes.buffer_destroy(&buf)

	pairs := make([dynamic]Pair, 0, 100)

	line: string
	err: io.Error
	for ; err != .EOF; line, err = bytes.buffer_read_string(&buf, '\n') {
		if line == "" {
			continue
		}

		line = line[:len(line) - 1]
		elves := strings.split(line, ",")
		elf1 := strings.split(elves[0], "-")
		elf2 := strings.split(elves[1], "-")

		pair := Pair {
			elf1 = {strconv.atoi(elf1[0]), strconv.atoi(elf1[1])},
			elf2 = {strconv.atoi(elf2[0]), strconv.atoi(elf2[1])},
		}

		append(&pairs, pair)
	}

	return pairs
}

// is b contained inside a?
pair_inside :: proc(a, b: [2]int) -> bool {
	return b[0] >= a[0] && b[1] <= a[1]
}

pair_overlap :: proc(a, b: [2]int) -> bool {
	return (a[0] <= b[0] && a[1] >= b[0]) || (a[0] <= b[1] && a[1] >= b[1])
}

task1 :: proc(pairs: []Pair) -> int {
	count: int

	for pair in pairs {
		if pair_inside(pair.elf1, pair.elf2) || pair_inside(pair.elf2, pair.elf1) {
			count += 1
		}
	}

	return count
}

task2 :: proc(pairs: []Pair) -> int {
	count: int

	for pair in pairs {
		if pair_inside(pair.elf1, pair.elf2) ||
		   pair_inside(pair.elf2, pair.elf1) ||
		   pair_overlap(pair.elf1, pair.elf2) {
			count += 1
		}
	}

	return count
}

main :: proc() {
	rucksacks := parse_input_file("input.txt")
	defer delete(rucksacks)

	result1 := task1(rucksacks[:])
	fmt.printf("Task 1 result: %v\n", result1)

	result2 := task2(rucksacks[:])
	fmt.printf("Task 2 result: %v\n", result2)
}
