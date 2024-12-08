package main

import "core:log"
import "core:os"
import "core:slice"
import "core:strconv"
import "core:strings"

Rule :: struct {
	before: [dynamic]int,
	after:  [dynamic]int,
}

Input :: struct {
	rules:   map[int]Rule,
	updates: [dynamic][]int,
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

	input.rules = map[int]Rule{}
	input.updates = make([dynamic][]int, 0, 10)

	rules_done: bool
	string_data := string(raw_data)
	for line in strings.split_lines_iterator(&string_data) {
		if len(line) == 0 {
			rules_done = true
			continue
		}

		if !rules_done {
			split := strings.index_rune(line, '|')
			a := strconv.atoi(line[:split])
			b := strconv.atoi(line[split + 1:])

			if a not_in input.rules {
				input.rules[a] = {
					before = make([dynamic]int, 0, 10),
					after  = make([dynamic]int, 0, 10),
				}
			}
			if b not_in input.rules {
				input.rules[b] = {
					before = make([dynamic]int, 0, 10),
					after  = make([dynamic]int, 0, 10),
				}
			}
			rule_a := &input.rules[a]
			rule_b := &input.rules[b]
			append(&rule_a.before, b)
			append(&rule_b.after, a)
		} else {
			raw_values := strings.split(line, ",")
			defer delete(raw_values)

			values := make([]int, len(raw_values))
			for v, i in raw_values {
				values[i] = strconv.atoi(v)
			}
			append(&input.updates, values)
		}
	}

	return input
}

free_input :: proc(input: ^Input) {
	for _, i in input.updates {
		delete(input.updates[i])
	}
	delete(input.updates)

	for k in input.rules {
		delete(input.rules[k].before)
		delete(input.rules[k].after)
	}
	delete(input.rules)
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

	for values in input.updates {
		valid := true
		for v in values {
			after: bool
			for u in values {
				if u == v {
					after = true
					continue
				}

				if (!after && !slice.contains(input.rules[v].after[:], u)) ||
				   (after && !slice.contains(input.rules[v].before[:], u)) {
					valid = false
					break
				}
			}
		}
		if valid {
			log.debugf("%v VALID", values)
			result += Result1(values[len(values) / 2])
		} else {
			log.debugf("%v INVALID", values)
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

	for values in input.updates {
		valid := true
		for v in values {
			after: bool
			for u in values {
				if u == v {
					after = true
					continue
				}

				if (!after && !slice.contains(input.rules[v].after[:], u)) ||
				   (after && !slice.contains(input.rules[v].before[:], u)) {
					valid = false
					break
				}
			}
		}
		if !valid {
			log.debugf("%v INVALID", values)
			less :: proc(a, b: int) -> bool {
				rules := cast(^map[int]Rule)context.user_ptr
				return slice.contains(rules[a].before[:], b)
			}
			context.user_ptr = &input.rules
			slice.sort_by(values, less)
			result += Result2(values[len(values) / 2])
		}
	}

	return result
}

print_result2 :: proc(result: ^Result2) {
	log.infof("Task 2: %d", result^)
}
// --- Task 2 --- //
