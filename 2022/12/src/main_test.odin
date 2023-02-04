package main

import "core:testing"

@(test)
task1_test :: proc(t: ^testing.T) {
	input := parse_input_file("../input_test.txt")
	defer delete(input.data)
	expected :: 31

	result := task1(area = input, debug = true)
	testing.expect_value(t, result, expected)
}

@(test)
task2_test :: proc(t: ^testing.T) {
	input := parse_input_file("../input_test.txt")
	defer delete(input.data)
	expected :: 29

	result := task2(area = input, debug = true)
	testing.expect_value(t, result, expected)
}
