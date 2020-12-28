package main

import (
	"os"
	"testing"
)

func readTestInput() []int {
	file, _ := os.Open("input_test.txt")
	defer file.Close()
	return readInput(file)
}

func TestTask1ShouldFindResult(t *testing.T) {
	input := readTestInput()
	expect := 14897079

	result := Task1(input, true)
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	input := readTestInput()
	expect := -1

	result := Task2(input, true)
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}
