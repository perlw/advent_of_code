package main

import (
	"fmt"
	"testing"
)

func TestTask1ShouldFindResult(t *testing.T) {
	input := []int{0, 3, 6}
	expected := 436

	result := Task1(input)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	tests := [][]int{
		{0, 3, 6},
		{1, 3, 2},
		{2, 1, 3},
		{1, 2, 3},
		{2, 3, 1},
		{3, 2, 1},
		{3, 1, 2},
	}
	expected := []int{
		175594,
		2578,
		3544142,
		261214,
		6895259,
		18,
		362,
	}

	for i, input := range tests {
		fmt.Printf("Test %d: %+v\n", i, input)
		result := Task2(input)
		if result != expected[i] {
			t.Errorf("%d: got %d, expected %d", i, result, expected[i])
		}
	}
}
