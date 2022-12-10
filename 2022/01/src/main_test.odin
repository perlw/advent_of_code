package main

import "core:testing"

@(test)
task1_test :: proc(t: ^testing.T) {
	input := []Elf{
		{total_calories = 6000},
		{total_calories = 4000},
		{total_calories = 11000},
		{total_calories = 24000},
		{total_calories = 10000},
	}
	expected :: 24000

	result := task1(input[:])
	testing.expect_value(t, result, expected)
}

@(test)
task2_test :: proc(t: ^testing.T) {
	input := []Elf{
		{total_calories = 6000},
		{total_calories = 4000},
		{total_calories = 11000},
		{total_calories = 24000},
		{total_calories = 10000},
	}
	expected :: 45000

	result := task2(input[:])
	testing.expect_value(t, result, expected)
}
