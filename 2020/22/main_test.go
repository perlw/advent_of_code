package main

import (
	"os"
	"testing"
)

func readTestInput() []Deck {
	file, _ := os.Open("input_test.txt")
	defer file.Close()
	return readInput(file)
}

func TestShouldReadInput(t *testing.T) {
	input := readTestInput()
	expected := Deck{9, 2, 6, 3, 1}

	if len(expected) != len(input[0]) {
		t.Errorf("did not read input correctly, length does not match")
	}
	for i := range input[0] {
		if expected[i] != input[0][i] {
			t.Errorf("did not read input correctly, did not match")
		}
	}
}

func TestTask1ShouldFindResult(t *testing.T) {
	input := readTestInput()
	expect := 306

	result := Task1(input, true)
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	input := readTestInput()
	expect := 291

	result := Task2(input, true)
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}
