package main

import "core:os"

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

	input := parse_input_file("input.txt")
	defer free_input(&input)

	result1 := run_task1(&input, debug1)
	print_result1(&result1)

	result2 := run_task2(&input, debug2)
	print_result2(&result2)
}
