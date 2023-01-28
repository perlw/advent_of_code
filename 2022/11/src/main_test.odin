package main

import "core:testing"

@(test)
task1_test :: proc(t: ^testing.T) {
	input := parse_input_file("../input_test.txt")
	defer delete(input)
	defer free_monkeys(input[:])
	expected :: 10605

	result := task1(monkeys = input[:], debug = true)
	testing.expect_value(t, result, expected)
}

@(test)
task2_test :: proc(t: ^testing.T) {
	input := parse_input_file("../input_test.txt")
	defer delete(input)
	defer free_monkeys(input[:])
	expected :: 2713310158

	result := task2(monkeys = input[:], debug = false)
	testing.expect_value(t, result, expected)
}
