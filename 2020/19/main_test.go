package main

import "testing"

func TestTask1ShouldFindResult(t *testing.T) {
	rules := map[int]Rule{
		0: {
			Match: [][]int{
				{4, 1, 5},
			},
		},
		1: {
			Match: [][]int{
				{2, 3},
				{3, 2},
			},
		},
		2: {
			Match: [][]int{
				{4, 4},
				{5, 5},
			},
		},
		3: {
			Match: [][]int{
				{4, 5},
				{5, 4},
			},
		},
		4: {
			R: 'a',
		},
		5: {
			R: 'b',
		},
	}
	messages := []string{
		"ababbb",
		"bababa",
		"abbbab",
		"aaabbb",
		"aaaabbb",
	}
	expected := 2

	result := Task1(rules, messages)
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
