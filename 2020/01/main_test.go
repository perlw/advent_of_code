package main

import "testing"

func TestTask1ShouldFindResult(t *testing.T) {
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	expect := 514579

	result := Task1(input)
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	input := []int{
		1721,
		979,
		366,
		299,
		675,
		1456,
	}
	expect := 241861950

	result := Task2(input)
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}
