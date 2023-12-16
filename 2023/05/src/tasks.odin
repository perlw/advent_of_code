package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:os"
import "core:strconv"
import "core:strings"

Input :: struct {
	seeds: [dynamic]int,
	maps:  [dynamic][dynamic][3]int,
}

Result1 :: distinct int
Result2 :: distinct int


// --- Input --- //
print_input :: proc(input: ^Input) {
	fmt.printf("%#v\n", input)
}

parse_input_file :: proc(filepath: string) -> Input {
	input: Input

	raw_data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	buf: bytes.Buffer
	bytes.buffer_init(&buf, raw_data)
	defer bytes.buffer_destroy(&buf)

	input.seeds = make([dynamic]int, 0, 10)
	input.maps = make([dynamic][dynamic][3]int, 0, 10)

	line: string
	line, _ = bytes.buffer_read_string(&buf, '\n')

	seeds_string := strings.split(line, ": ")
	defer delete(seeds_string)
	seeds := strings.fields(seeds_string[1])
	defer delete(seeds)
	for seed in seeds {
		append(&input.seeds, strconv.atoi(seed))
	}

	bytes.buffer_read_string(&buf, '\n')
	bytes.buffer_read_string(&buf, '\n')

	current_map := make([dynamic][3]int, 0, 10)
	err: io.Error
	for ;; line, err = bytes.buffer_read_string(&buf, '\n') {
		if line == "" || line == "\n" {
			append(&input.maps, current_map)
			current_map = make([dynamic][3]int, 0, 10)
			if err == .EOF {
				break
			}
			continue
		}

		line = line[:len(line) - 1]
		if strings.index(line, ":") > -1 {
			continue
		}

		lookup := strings.fields(line)
		defer delete(lookup)
		append(&current_map, [3]int{strconv.atoi(lookup[0]), strconv.atoi(lookup[1]), strconv.atoi(lookup[2])})
	}

	return input
}

free_input :: proc(input: ^Input) {
	for m in input.maps {
		delete(m)
	}
	delete(input.maps)
	delete(input.seeds)
}
// --- Input --- //


// --- Task 1 --- //
translate_seed :: proc(m: [][3]int, seed: int, debug: bool = false) -> int {
	s := seed
	if debug {
		fmt.printf("\ttranslating: %d\n", s)
	}
	for t in m {
		dst_a, dst_b := t[0], t[0] + t[2]
		src_a, src_b := t[1], t[1] + t[2]
		if seed >= src_a && seed < src_b {
			s = dst_a + (seed - src_a)
			if debug {
				fmt.printf("\t\t%d->%d, %d->%d, %d=%d\n", dst_a, dst_b, src_a, src_b, seed, s)
			}
			break
		}
	}
	return s
}

run_task1 :: proc(input: ^Input, debug: bool) -> Result1 {
	result: Result1 = -1

	if debug {
		print_input(input)
	}

	for seed in input.seeds {
		s := seed
		if debug {
			fmt.printf("seed: %d\n", s)
		}
		for m in input.maps {
			s = translate_seed(m[:], s, debug)
		}
		if result == -1 || Result1(s) < result {
			result = Result1(s)
		}
		if debug {
			fmt.printf("location: %d->%d\n", seed, s)
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
	result: Result2 = -1

	if debug {
		print_input(input)
	}

	for i := 0; i < len(input.seeds); i += 2 {
		start := input.seeds[i]
		end := start + input.seeds[i + 1]

		fmt.printf("scanning %d-%d, %d seeds\n", start, end, end - start + 1)
		for seed in start ..= end {
			s := seed
			if debug {
				fmt.printf("seed: %d\n", s)
			}
			for m in input.maps {
				s = translate_seed(m[:], s, debug)
			}
			if result == -1 || Result2(s) < result {
				result = Result2(s)
				fmt.printf("New result: seed %d -> location %d\n", seed, s)
			}
			if debug {
				fmt.printf("location: %d->%d\n", seed, s)
			}
		}
	}

	return result
}

print_result2 :: proc(result: ^Result2) {
	fmt.printf("Task 2: %d\n", result^)
}
// --- Task 2 --- //
