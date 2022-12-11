package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:os"
import "core:slice"

Rucksack :: struct {
	items: []u8,
}

parse_input_file :: proc(filepath: string) -> [dynamic]Rucksack {
	data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	buf: bytes.Buffer
	bytes.buffer_init(&buf, data)
	defer bytes.buffer_destroy(&buf)

	rucksacks := make([dynamic]Rucksack, 0, 100)

	line: []u8
	err: io.Error
	for ; err != .EOF; line, err = bytes.buffer_read_bytes(&buf, '\n') {
		if len(line) == 0 {
			continue
		}

		append(&rucksacks, Rucksack{items = slice.clone(line[:len(line) - 1])})
	}

	return rucksacks
}

char_to_score :: proc(char: u8) -> int {
	score: int
	if char <= 'Z' {
		score = int(char - 'A') + 27
	} else {
		score = int(char - 'a') + 1
	}
	return score
}

task1 :: proc(rucksacks: []Rucksack) -> int {
	score: int

	for rucksack in rucksacks {
		first := make(map[u8]int)
		defer delete(first)
		last := make(map[u8]int)
		defer delete(last)

		half := len(rucksack.items) / 2
		for i := 0; i < half; i += 1 {
			first[rucksack.items[i]] = 1
			last[rucksack.items[half + i]] = 1
		}

		for item in first {
			if item in last {
				score += char_to_score(item)
			}
		}
	}

	return score
}

task2 :: proc(rucksacks: []Rucksack) -> int {
	score: int

	for i := 0; i < len(rucksacks); i += 3 {
		items := make(map[u8]int)
		defer delete(items)
		for rucksack in rucksacks[i:i + 3] {
			unique_items := make(map[u8]int)
			defer delete(unique_items)

			for item in rucksack.items {
				unique_items[item] += 1
			}

			for item in unique_items {
				items[item] += 1
			}
		}

		for item, amount in items {
			if amount >= 3 {
				score += char_to_score(item)
			}
		}
	}

	return score
}

main :: proc() {
	rucksacks := parse_input_file("input.txt")
	defer delete(rucksacks)

	result1 := task1(rucksacks[:])
	fmt.printf("Task 1 result: %v\n", result1)

	result2 := task2(rucksacks[:])
	fmt.printf("Task 2 result: %v\n", result2)
}
