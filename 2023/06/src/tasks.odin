package main

import "core:bytes"
import "core:fmt"
import "core:math"
import "core:os"
import "core:strconv"
import "core:strings"

Race :: struct {
	time:     int,
	distance: int,
}

Input :: struct {
	races: [dynamic]Race,
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

	input.races = make([dynamic]Race, 0, 10)

	times, distances: []int
	{
		line, _ := bytes.buffer_read_string(&buf, '\n')
		line = line[:len(line) - 1]
		line_parts := strings.split(line, ":")
		defer delete(line_parts)
		numbers := strings.fields(line_parts[1])
		times = make([]int, len(numbers))
		for n, i in numbers {
			times[i] = strconv.atoi(n)
		}
	}
	defer delete(times)
	{
		line, _ := bytes.buffer_read_string(&buf, '\n')
		line = line[:len(line) - 1]
		line_parts := strings.split(line, ":")
		defer delete(line_parts)
		numbers := strings.fields(line_parts[1])
		distances = make([]int, len(numbers))
		for n, i in numbers {
			distances[i] = strconv.atoi(n)
		}
	}
	defer delete(distances)

	for i in 0 ..< len(times) {
		append(&input.races, Race{time = times[i], distance = distances[i]})
	}

	return input
}

free_input :: proc(input: ^Input) {
	delete(input.races)
}
// --- Input --- //


// --- Task 1 --- //
run_task1 :: proc(input: ^Input, debug: bool) -> Result1 {
	result: Result1 = 1

	if debug {
		print_input(input)
	}

	for race in input.races {
		min, max: f32

		s := math.sqrt(f32((race.time * race.time) - (4 * -1 * -race.distance)))
		min = (f32(-race.time) + s) / -2
		max = (f32(-race.time) - s) / -2
		min = math.ceil(min + 1) if math.ceil(min) == min else math.ceil(min)
		max = math.floor(max - 1) if math.floor(max) == max else math.floor(max)

		result *= Result1(max - min + 1)
	}

	return result
}

print_result1 :: proc(result: ^Result1) {
	fmt.printf("Task 1: %d\n", result^)
}
// --- Task 1 --- //


// --- Task 2 --- //
run_task2 :: proc(input: ^Input, debug: bool) -> Result2 {
	if debug {
		print_input(input)
	}

	time, distance: int
	for race in input.races {
		if race.time < 10 {
			time *= 10
		} else if race.time < 100 {
			time *= 100
		} else if race.time < 1000 {
			time *= 1000
		} else if race.time < 10000 {
			time *= 10000
		}
		time += race.time
		if race.distance < 10 {
			distance *= 10
		} else if race.distance < 100 {
			distance *= 100
		} else if race.distance < 1000 {
			distance *= 1000
		} else if race.distance < 10000 {
			distance *= 10000
		}
		distance += race.distance
	}
	fmt.printf("%d %d\n", time, distance)

	min, max: f64
	s := math.sqrt(f64((time * time) - (4 * -1 * -distance)))
	min = (f64(-time) + s) / -2
	max = (f64(-time) - s) / -2
	min = math.ceil(min + 1) if math.ceil(min) == min else math.ceil(min)
	max = math.floor(max - 1) if math.floor(max) == max else math.floor(max)

	return Result2(max - min + 1)
}

print_result2 :: proc(result: ^Result2) {
	fmt.printf("Task 2: %d\n", result^)
}
// --- Task 2 --- //
