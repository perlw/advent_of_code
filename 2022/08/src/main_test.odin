package main

import "core:testing"

@(test)
task1_test :: proc(t: ^testing.T) {
	buffer: [4096]int
	stride, input := parse_input_file("../input_test.txt", buffer[:])
	expected :: 21

	result := task1(stride, input)
	testing.expect_value(t, result, expected)
}

@(test)
task2_test :: proc(t: ^testing.T) {
	buffer: [4096]int
	stride, input := parse_input_file("../input_test.txt", buffer[:])
	expected :: 8

	result := task2(stride, input)
	testing.expect_value(t, result, expected)
}
