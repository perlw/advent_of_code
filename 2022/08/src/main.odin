package main

import "core:fmt"
import "core:os"

parse_input_file :: proc(filepath: string, buffer: []int) -> (stride: int, sliced: []int) {
	data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	index: int
	for b, i in data {
		if b == '\n' {
			stride = i if stride == 0 else stride
			continue
		}
		buffer[index] = cast(int)(b - 48)
		index += 1
	}
	sliced = buffer[:index]

	return
}

Vec2 :: [2]int

check_visible :: proc(start: Vec2, dir: Vec2, stride: int, heights: []int) -> (bool, int) {
	max := heights[(start[1] * stride) + start[0]]
	current := start + dir
	steps: int
	for {
		if current[0] < 0 || current[0] >= stride {
			return true, steps
		}
		if current[1] < 0 || current[1] >= stride {
			return true, steps
		}
		i := (current[1] * stride) + current[0]
		if heights[i] >= max {
			return false, steps + 1
		}

		current += dir
		steps += 1
	}
	panic("unreachable!")
}

task1 :: proc(stride: int, heights: []int) -> int {
	result: int

	for y in 0 ..< stride {
		for x in 0 ..< stride {
			pos := Vec2{x, y}
			dirs := []Vec2{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
			visible: bool
			for d in dirs {
				if found_edge, _ := check_visible(pos, d, stride, heights); found_edge {
					visible = true
					break
				}
			}
			if visible {
				result += 1
			}
		}
	}

	return result
}

task2 :: proc(stride: int, heights: []int) -> int {
	result: int

	for y in 0 ..< stride {
		for x in 0 ..< stride {
			pos := Vec2{x, y}
			dirs := []Vec2{{0, -1}, {-1, 0}, {1, 0}, {0, 1}}
			score := 1
			for d in dirs {
				_, distance := check_visible(pos, d, stride, heights)
				score *= distance
			}
			result = score if score > result else result
		}
	}

	return result
}

main :: proc() {
	buffer: [65353]int
	stride, heights := parse_input_file("input.txt", buffer[:])

	result1 := task1(stride, heights)
	fmt.printf("Task 1 result: %d\n", result1)

	result2 := task2(stride, heights)
	fmt.printf("Task 2 result: %d\n", result2)
}
