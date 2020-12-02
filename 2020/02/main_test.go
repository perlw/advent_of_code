package main

import "testing"

func TestTask1ShouldFindResult(t *testing.T) {
	input := []Input{
		{ruleMin: 1, ruleMax: 3, ruleRune: 'a', password: "abcde"},
		{ruleMin: 1, ruleMax: 3, ruleRune: 'b', password: "cdefg"},
		{ruleMin: 2, ruleMax: 9, ruleRune: 'c', password: "ccccccccc"},
	}
	expect := 2

	result := Task1(input)
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	input := []Input{
		{ruleMin: 1, ruleMax: 3, ruleRune: 'a', password: "abcde"},
		{ruleMin: 1, ruleMax: 3, ruleRune: 'b', password: "cdefg"},
		{ruleMin: 2, ruleMax: 9, ruleRune: 'c', password: "ccccccccc"},
	}
	expect := 1

	result := Task2(input)
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}
