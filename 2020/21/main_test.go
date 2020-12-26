package main

import (
	"os"
	"testing"
)

func readTestInput() []Food {
	file, _ := os.Open("input_test.txt")
	defer file.Close()
	return readInput(file)
}

func TestShouldReadInput(t *testing.T) {
	input := readTestInput()
	expected := Food{
		Ingredients: []string{"trh", "fvjkl", "sbzzf", "mxmxvkd"},
		Allergens:   []string{"dairy"},
	}

	var found bool
loop:
	for _, in := range input {
		if len(expected.Ingredients) != len(in.Ingredients) {
			continue
		}
		if len(expected.Allergens) != len(in.Allergens) {
			continue
		}
		found = true
		for i := range expected.Ingredients {
			if in.Ingredients[i] != expected.Ingredients[i] {
				found = false
				continue loop
			}
		}
	}
	if !found {
		t.Errorf("did not read input correctly, did not find expected")
	}
}

func TestTask1ShouldFindResult(t *testing.T) {
	input := readTestInput()
	expect := 5

	result := Task1(input, true)
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	input := readTestInput()
	expect := "mxmxvkd,sqjhc,fvjkl"

	result := Task2(input, true)
	if result != expect {
		t.Errorf("got %s, expected %s", result, expect)
	}
}
