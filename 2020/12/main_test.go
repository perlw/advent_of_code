package main

import (
	"testing"
)

func TestTask1ShouldFindResult(t *testing.T) {
	input := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	expected := 25

	result := Task1(input)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	input := []string{
		"F10",
		"N3",
		"F7",
		"R90",
		"F11",
	}
	expected := 286

	result := Task2(input)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}
