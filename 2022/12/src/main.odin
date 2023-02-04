package main

import "core:bytes"
import "core:fmt"
import "core:mem"
import "core:io"
import "core:os"

Area :: struct {
	data: []u8,
	w:    uint,
	h:    uint,
}

// NOTE: Have to free Area.data
parse_input_file :: proc(filepath: string) -> Area {
	data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	buf: bytes.Buffer
	bytes.buffer_init(&buf, data)
	defer bytes.buffer_destroy(&buf)

	area: Area
	buffer := make([dynamic]u8, 0, 10)
	defer delete(buffer)

	line: string
	err: io.Error
	for ; err != .EOF; line, err = bytes.buffer_read_string(&buf, '\n') {
		if line == "" || line == "\n" {
			continue
		}
		line = line[:len(line) - 1]

		if area.w == 0 {
			area.w = len(line)
		}
		area.h += 1
		append(&buffer, line)
	}

	area.data = mem.clone_slice(buffer[:])

	return area
}

print_area :: proc(area: Area) {
	fmt.printf("\n-=-=-=-=-\n")
	for c, i in area.data {
		if c == 'a' - 1 {
			fmt.printf("S")
		} else if c == 'z' + 1 {
			fmt.printf("E")
		} else {
			fmt.printf("%c", c)
		}
		if uint(i + 1) % area.w == 0 {
			fmt.printf("\n")
		}
	}
	fmt.printf("-=-=-=-=-\n\n")
}

print_distances :: proc(values: []int, w: uint) {
	fmt.printf("\n")
	for c, i in values {
		fmt.printf("%2d|", c if c < 999999 else -1)
		if uint(i + 1) % w == 0 {
			fmt.printf("\n")
		}
	}
	fmt.printf("\n")
}

task1 :: proc(area: Area, debug := false) -> int {
	if debug {
		print_area(area)
	}

	visited := make([]bool, len(area.data))
	defer delete(visited)
	distances := make([]int, len(area.data))
	defer delete(distances)

	start_i, end_i: int
	for c, i in area.data {
		if c == 'S' {
			start_i = i
			area.data[i] = 'a'
		} else if c == 'E' {
			area.data[i] = 'z'
			end_i = i
		}
		distances[i] = 999999
	}
	distances[start_i] = 0

	current := start_i
	for {
		visited[current] = true
		neighbors := [4]int{
			current - int(area.w) if current >= int(area.w) else -1,
			current + int(area.w) if current < len(area.data) - int(area.w) else -1,
			current - 1 if current % int(area.w) > 0 else -1,
			current + 1 if current % int(area.w) < int(area.w) - 1 else -1,
		}

		dist := distances[current]
		new_dist := dist + 1
		for n in neighbors {
			if n > -1 && area.data[n] <= area.data[current] + 1 {
				distances[n] = distances[n] if distances[n] < new_dist else new_dist
			}
		}

		least := 999999
		next_i := -1
		for d, i in distances {
			if !visited[i] && d < least {
				least = d
				next_i = i
			}
		}

		current = next_i
		if next_i == -1 {
			break
		}
	}

	if debug {
		print_area(area)
		print_distances(distances, area.w)
	}

	return distances[end_i]
}

// NOTE: Simply bruteforces the solution. There are more effecient solutions, I know.
task2 :: proc(area: Area, debug := false) -> int {
	for c, i in area.data {
		if c == 'S' {
			area.data[i] = 'a'
		}
	}

	least := 999999
	for i in 0 ..< len(area.data) {
		fmt.printf("%d/%d...", i, len(area.data))
		if area.data[i] == 'a' {
			input := clone_area(area)
			defer delete(input.data)

			input.data[i] = 'S'
			result := task1(input, debug)
			if result < least {
				least = result
			}
			fmt.printf(" %d!\n", result)
		} else {
			fmt.printf(" SKIP!\n")
		}
	}

	return least
}

clone_area :: proc(area: Area) -> Area {
	return {data = mem.clone_slice(area.data), w = area.w, h = area.h}
}

main :: proc() {
	area := parse_input_file("input.txt")
	defer delete(area.data)

	{
		input := clone_area(area)
		defer delete(input.data)

		result1 := task1(input)
		fmt.printf("Task 1 result: %d\n", result1)
	}

	{
		input := clone_area(area)
		defer delete(input.data)

		result2 := task2(input)
		fmt.printf("Task 2 result: %d\n", result2)
	}
}
