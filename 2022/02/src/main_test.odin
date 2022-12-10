package main

import "core:testing"

@(test)
task1_test :: proc(t: ^testing.T) {
	input := []Game{{'A', 'Y'}, {'B', 'X'}, {'C', 'Z'}}
	expected :: 15

	result := task1(input[:])
	testing.expect_value(t, result, expected)
}

@(test)
task2_test :: proc(t: ^testing.T) {
	input := []Game{{'A', 'Y'}, {'B', 'X'}, {'C', 'Z'}}
	expected :: 12

	result := task2(input[:])
	testing.expect_value(t, result, expected)
}
