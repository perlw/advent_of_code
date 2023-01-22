package main

import "core:testing"

@(test)
task1_test :: proc(t: ^testing.T) {
	input := parse_input_file("../input_test.txt")
	defer delete(input)
	expected :: 13140

	result := task1(instructions = input[:], debug = true)
	testing.expect_value(t, result, expected)
}

// NOTE: Lazy
@(test)
task2_test :: proc(t: ^testing.T) {
	input := parse_input_file("../input_test.txt")
	defer delete(input)
	task2(instructions = input[:], debug = true)
}
