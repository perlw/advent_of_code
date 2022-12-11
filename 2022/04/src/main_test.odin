package main

import "core:testing"

@(test)
task1_test :: proc(t: ^testing.T) {
	input := parse_input_file("../input_test.txt")
	defer delete(input)
	expected :: 2

	result := task1(input[:])
	testing.expect_value(t, result, expected)
}

@(test)
task2_test :: proc(t: ^testing.T) {
	input := parse_input_file("../input_test.txt")
	defer delete(input)
	expected :: 4

	result := task2(input[:])
	testing.expect_value(t, result, expected)
}
