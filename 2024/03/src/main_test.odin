package main

import "core:testing"

@(test)
task1_test :: proc(t: ^testing.T) {
	input := parse_input_file("../input_test.txt")
	defer free_input(&input)
	expected: Result1 = 161

	result := run_task1(input = &input, debug = true)
	testing.expect_value(t, result, expected)
}

@(test)
task2_test :: proc(t: ^testing.T) {
	input := parse_input_file("../input_test.txt")
	defer free_input(&input)
	expected: Result2 = 48

	result := run_task2(input = &input, debug = true)
	testing.expect_value(t, result, expected)
}
