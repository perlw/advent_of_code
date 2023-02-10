package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:os"

Item :: struct {
	list:   ^List,
	number: int,
}

List :: struct {
	parent: ^List,
	items:  [dynamic]Item,
}

// NOTE: Have to free all of Packet
parse_input_file :: proc(filepath: string) -> [dynamic]List {
	data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	buf: bytes.Buffer
	bytes.buffer_init(&buf, data)
	defer bytes.buffer_destroy(&buf)

	packets := make([dynamic]List, 0, 10)

	line: string
	err: io.Error
	for ; err != .EOF; line, err = bytes.buffer_read_string(&buf, '\n') {
		if line == "" || line == "\n" {
			continue
		}
		line = line[:len(line) - 1]

		packet: List
		current: ^List
		for i := 0; i < len(line[:]); i += 1 {
			c := line[i]
			switch (c) {
			case '[':
				if current == nil {
					current = &packet
					current.items = make([dynamic]Item, 0, 10)
				} else {
					new_list := new(List)
					new_list.parent = current
					new_list.items = make([dynamic]Item, 0, 10)
					append(&current.items, Item{list = new_list})
					current = new_list
				}
			case ']':
				current = current.parent
			case '0' ..= '9':
				num := int(c - '0')
				if (line[i + 1] >= '0' && line[i + 1] <= '9') {
					num = (num * 10) + int(line[i + 1] - '0')
					i += 1
				}
				append(&current.items, Item{number = num})
			case:
			}
		}

		append(&packets, packet)
	}

	return packets
}

free_list :: proc(list: ^List) {
	for item in list.items {
		if item.list != nil {
			free_list(item.list)
			free(item.list)
		}
	}
}

free_lists :: proc(lists: [dynamic]List) {
	for _, i in lists {
		free_list(&lists[i])
	}
	delete(lists)
}

print_list :: proc(list: []Item) {
	fmt.printf(" [")
	for item in list {
		if item.list != nil {
			print_list(item.list.items[:])
		} else {
			fmt.printf(" %d ", item.number)
		}
	}
	fmt.printf("] ")
}

compare_lists :: proc(a, b: []Item, depth := 0, debug := false) -> int {
	if debug {
		for _ in 0 ..< depth {
			fmt.printf("  ")
		}
		fmt.printf("- Compare ")
		print_list(a)
		fmt.printf(" vs ")
		print_list(b)
		fmt.printf("\n")
	}

	max_len := len(a) if len(a) < len(b) else len(b)
	for i in 0 ..< max_len {
		if a[i].list == nil && b[i].list == nil {
			x, y := a[i].number, b[i].number
			if debug {
				for _ in 0 ..< depth + 1 {
					fmt.printf("  ")
				}
				fmt.printf("- Compare %d vs %d\n", x, y)
			}
			if x < y {
				if debug {
					for _ in 0 ..< depth + 2 {
						fmt.printf("  ")
					}
					fmt.printf("- Left side is smaller, so inputs are in the right order\n")
				}
				return 1
			} else if x > y {
				if debug {
					for _ in 0 ..< depth + 2 {
						fmt.printf("  ")
					}
					fmt.printf("- Right side is smaller, so inputs are not in the right order\n")
				}
				return -1
			}
		} else {
			aa := a[i].list.items[:] if a[i].list != nil else []Item{{number = a[i].number}}
			bb := b[i].list.items[:] if b[i].list != nil else []Item{{number = b[i].number}}
			cmp := compare_lists(aa, bb, depth + 1, debug)
			if cmp != 0 {
				return cmp
			}
		}
	}
	if len(a) < len(b) {
		if debug {
			for _ in 0 ..< depth + 1 {
				fmt.printf("  ")
			}
			fmt.printf("- Left side ran out of items, so inputs are in the right order\n")
		}
		return 1
	}
	if len(a) > len(b) {
		if debug {
			for _ in 0 ..< depth + 1 {
				fmt.printf("  ")
			}
			fmt.printf("- Right side ran out of items, so inputs are not in the right order\n")
		}
		return -1
	}

	return 0
}

task1 :: proc(packets: []List, debug := false) -> int {
	result: int

	num_pairs := len(packets) / 2
	for i in 0 ..< num_pairs {
		packet_a := packets[i * 2].items[:]
		packet_b := packets[(i * 2) + 1].items[:]

		if debug {
			fmt.printf("== Pair %d ==\n", i + 1)
		}

		if compare_lists(packet_a, packet_b, 0, debug) == 1 {
			result += i + 1
			fmt.printf("%d is okay\n", i + 1)
		}

		if debug {
			fmt.printf("\n")
		}
	}

	return result
}

task2 :: proc(packets: []List, debug := false) -> int {
	return -1
}

main :: proc() {
	debug1: bool
	debug2: bool
	if len(os.args) > 1 {
		for arg in os.args {
			switch arg {
			case "-d1":
				debug1 = true
			case "-d2":
				debug2 = true
			}
		}
	}

	packets := parse_input_file("input.txt")
	defer free_lists(packets)

	result1 := task1(packets[:], debug1)
	fmt.printf("Task 1 result: %d\n", result1)

	result2 := task2(packets[:], debug2)
	fmt.printf("Task 2 result: %d\n", result2)
}
