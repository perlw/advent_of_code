package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:os"
import "core:strconv"

parse_input_file :: proc(filepath: string) -> []u8 {
	data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	return data
}

Node :: struct {
	size:     int,
	children: Node_Map,
}

Node_Map :: map[string]Node

parse_filesystem :: proc(buf: ^bytes.Buffer, parsed: ^Node_Map, current_directory: string = "") {
	line: []u8
	err: io.Error
	for ; err != .EOF; line, err = bytes.buffer_read_bytes(buf, '\n') {
		if len(line) == 0 {
			continue
		}
		line = line[:len(line) - 1]

		cmd := string(line[2:4])
		if cmd == "cd" {
			target := string(line[5:])
			if target == ".." {
				break
			}

			if !(target in parsed) {
				parsed[target] = {
					size     = 0,
					children = make(Node_Map),
				}
			}
			dir := &parsed[target]
			parse_filesystem(buf, &dir.children, target)
		} else if cmd == "ls" {
			continue
		} else {
			split: int
			for c, i in line {
				if c == ' ' {
					split = i
				}
			}

			size_part := string(line[:split])
			name_part := string(line[split + 1:])
			if !(name_part in parsed) {
				if size_part == "dir" {
					parsed[name_part] = {
						size     = 0,
						children = make(Node_Map),
					}
				} else {
					parsed[name_part] = {
						size     = strconv.atoi(size_part),
						children = nil,
					}
				}
			}
		}
	}
}

update_dir_size :: proc(parsed_filesystem: ^Node_Map) -> int {
	size: int
	for _, node in parsed_filesystem {
		if node.size == 0 {
			node.size = update_dir_size(&node.children)
		}
		size += node.size
	}
	return size
}

find_total_of_max_size :: proc(parsed_filesystem: ^Node_Map, max_size: int) -> int {
	size: int
	for _, node in parsed_filesystem {
		if node.children != nil {
			if node.size <= 100000 {
				size += node.size
			}
			size += find_total_of_max_size(&node.children, max_size)
		}
	}
	return size
}

find_size_to_save :: proc(parsed_filesystem: ^Node_Map, total_size, target_size: int) -> int {
	min_size := 70000000
	for _, node in parsed_filesystem {
		if node.children != nil {
			if total_size - node.size <= target_size {
				if node.size < min_size {
					min_size = node.size
				}
			}
			size := find_size_to_save(&node.children, total_size, target_size)
			if size < min_size {
				min_size = size
			}
		}
	}
	return min_size
}

task1 :: proc(filesystem: []u8) -> int {
	buf: bytes.Buffer
	bytes.buffer_init(&buf, filesystem)
	defer bytes.buffer_destroy(&buf)

	parsed_filesystem := make(Node_Map)
	defer delete(parsed_filesystem)
	parse_filesystem(&buf, &parsed_filesystem)
	update_dir_size(&parsed_filesystem)

	return find_total_of_max_size(&parsed_filesystem, 100000)
}

task2 :: proc(filesystem: []u8) -> int {
	buf: bytes.Buffer
	bytes.buffer_init(&buf, filesystem)
	defer bytes.buffer_destroy(&buf)

	parsed_filesystem := make(Node_Map)
	defer delete(parsed_filesystem)
	parse_filesystem(&buf, &parsed_filesystem)
	update_dir_size(&parsed_filesystem)

	return find_size_to_save(&parsed_filesystem, parsed_filesystem["/"].size, 40000000)
}

main :: proc() {
	filesystem := parse_input_file("input.txt")
	defer delete(filesystem)

	result1 := task1(filesystem)
	fmt.printf("Task 1 result: %d\n", result1)

	result2 := task2(filesystem)
	fmt.printf("Task 2 result: %d\n", result2)
}
