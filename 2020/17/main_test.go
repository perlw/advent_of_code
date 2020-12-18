package main

import (
	"testing"
)

func TestTask1ShouldFindResult(t *testing.T) {
	input := []rune{
		'.', '#', '.',
		'.', '.', '#',
		'#', '#', '#',
	}
	expected := 112

	result := Task1(input, 3)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	expected := -1

	result := Task2()
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}
