package main

import "core:testing"

@(test)
task1_test :: proc(t: ^testing.T) {
	state, instructions := parse_input_file("../input_test.txt")
	defer free_input_data(state, instructions)
	expected := []u8{'C', 'M', 'Z'}

	result := task1(state, instructions[:])
	defer delete(result)
	if len(result) != len(expected) {
		testing.errorf(t, "incorrect result length, expected: [%s], got [%s]", expected, result)
	}
	for char, i in result {
		testing.expect_value(t, char, expected[i])
	}
}

@(test)
task2_test :: proc(t: ^testing.T) {
	state, instructions := parse_input_file("../input_test.txt")
	defer free_input_data(state, instructions)
	expected := []u8{'M', 'C', 'D'}

	result := task2(state, instructions[:])
	defer delete(result)
	if len(result) != len(expected) {
		testing.errorf(t, "incorrect result length, expected: [%s], got [%s]", expected, result)
	}
	for char, i in result {
		testing.expect_value(t, char, expected[i])
	}
}
