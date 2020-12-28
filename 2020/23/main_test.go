package main

import (
	"os"
	"testing"
)

func readTestInput() Game {
	file, _ := os.Open("input_test.txt")
	defer file.Close()
	return readInput(file)
}

func TestShouldReadInput(t *testing.T) {
	input := readTestInput()
	expected := []int{3, 8, 9, 1, 2, 5, 4, 6, 7}

	if input.Cups.Root == nil {
		t.Errorf("did not read input correctly, no data?")
	}
	var i int
	for e := input.Cups.Root; e != nil; e = e.Next {
		if expected[i] != e.Val {
			t.Errorf("did not read input correctly, did not match")
		}
		i++
		if i >= input.Cups.Count {
			break
		}
	}
}

func TestTask1ShouldFindResultShort(t *testing.T) {
	input := readTestInput()
	expect := "92658374"

	result := Task1(input, 10, true)
	if result != expect {
		t.Errorf("got %s, expected %s", result, expect)
	}
}

func TestTask1ShouldFindResult(t *testing.T) {
	input := readTestInput()
	expect := "67384529"

	result := Task1(input, 100, true)
	if result != expect {
		t.Errorf("got %s, expected %s", result, expect)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	input := readTestInput()
	expect := 149245887792

	result := Task2(input, false)
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}
