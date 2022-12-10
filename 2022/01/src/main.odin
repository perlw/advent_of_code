package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:os"
import "core:slice"
import "core:sort"
import "core:strconv"

Elf :: struct {
	total_calories: int,
}

parse_input_file :: proc(filepath: string) -> [dynamic]Elf {
	data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	buf: bytes.Buffer
	bytes.buffer_init(&buf, data)
	defer bytes.buffer_destroy(&buf)

	elves := make([dynamic]Elf, 0, 100)

	line: string
	err: io.Error
	current_calories: int
	for ; err != .EOF; line, err = bytes.buffer_read_string(&buf, '\n') {
		if line == "\n" {
			append(&elves, Elf{total_calories = current_calories})
			current_calories = 0
			continue
		}

		current_calories += strconv.atoi(line)
	}

	return elves
}

task1 :: proc(elves: []Elf) -> int {
	highest_calories: int
	for elf in elves {
		if elf.total_calories > highest_calories {
			highest_calories = elf.total_calories
		}
	}
	return highest_calories
}

sum_calories :: proc(elves: []Elf) -> int {
	sum: int
	for elf in elves {
		sum += elf.total_calories
	}
	return sum
}

task2 :: proc(elves: []Elf) -> int {
	sorted_elves := slice.clone(elves)
	defer delete(sorted_elves)
	sort.quick_sort_proc(sorted_elves, proc(a, b: Elf) -> int {
		return b.total_calories - a.total_calories
	})
	return sum_calories(sorted_elves[:3])
}

main :: proc() {
	elves := parse_input_file("input.txt")
	defer delete(elves)

	result1 := task1(elves[:])
	fmt.printf("Task 1 result: %v\n", result1)

	result2 := task2(elves[:])
	fmt.printf("Task 2 result: %v\n", result2)
}
