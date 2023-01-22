package main

import "core:testing"

@(test)
task1_test :: proc(t: ^testing.T) {
	buffer: [128]byte
	playfield_data: [256]byte
	input := parse_input_file("../input_test.txt", buffer[:])
	expected :: 13

	playfield := Grid {
		data = playfield_data[:],
		w    = 16,
		h    = 16,
	}
	result := task1(input, &playfield)
	testing.expect_value(t, result, expected)
}

@(test)
task2_test1 :: proc(t: ^testing.T) {
	buffer: [128]byte
	playfield_data: [256]byte
	input := parse_input_file("../input_test.txt", buffer[:])
	expected :: 1

	playfield := Grid {
		data = playfield_data[:],
		w    = 16,
		h    = 16,
	}
	result := task1(instructions = input, playfield = &playfield, debug = true)
	testing.expect_value(t, result, expected)
}

@(test)
task2_test2 :: proc(t: ^testing.T) {
	buffer: [128]byte
	playfield_data: [4096]byte
	input := parse_input_file("../input_test2.txt", buffer[:])
	expected :: 36

	playfield := Grid {
		data = playfield_data[:],
		w    = 64,
		h    = 64,
	}
	result := task2(instructions = input, playfield = &playfield, debug = true)
	testing.expect_value(t, result, expected)
}
