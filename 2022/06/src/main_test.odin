package main

import "core:testing"

@(test)
task1_test :: proc(t: ^testing.T) {
	Test :: struct {
		input:    []u8,
		expected: int,
	}

	tests := []Test{
		{input = transmute([]u8)string("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), expected = 7},
		{input = transmute([]u8)string("bvwbjplbgvbhsrlpgdmjqwftvncz"), expected = 5},
		{input = transmute([]u8)string("nppdvjthqldpwncqszvftbrmjlhg"), expected = 6},
		{input = transmute([]u8)string("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), expected = 10},
		{input = transmute([]u8)string("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), expected = 11},
	}

	for test in tests {
		result := task1(test.input)
		testing.expect_value(t, result, test.expected)
	}
}

@(test)
task2_test :: proc(t: ^testing.T) {
	Test :: struct {
		input:    []u8,
		expected: int,
	}

	tests := []Test{
		{input = transmute([]u8)string("mjqjpqmgbljsphdztnvjfqwrcgsmlb"), expected = 19},
		{input = transmute([]u8)string("bvwbjplbgvbhsrlpgdmjqwftvncz"), expected = 23},
		{input = transmute([]u8)string("nppdvjthqldpwncqszvftbrmjlhg"), expected = 23},
		{input = transmute([]u8)string("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"), expected = 29},
		{input = transmute([]u8)string("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"), expected = 26},
	}

	for test in tests {
		result := task2(test.input)
		testing.expect_value(t, result, test.expected)
	}
}
