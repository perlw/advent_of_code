package main

import (
	"testing"
)

func TestBusShouldHaveTimetable(t *testing.T) {
	start := 939
	tests := []Bus{7, 13, 19, 31, 59}
	expected := []int{945, 949, 950, 961, 944}

	for i, test := range tests {
		result := test.NextDeparture(start)
		if result != expected[i] {
			t.Errorf("Bus %d: got %d, expected %d", test, result, expected[i])
		}
	}
}

func TestTask1ShouldFindResult(t *testing.T) {
	start := 939
	input := []Bus{7, 13, -1, -1, 59, -1, 31, 19}
	expected := 295

	result := Task1(start, input)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	tests := [][]Bus{
		{7, 13, -1, -1, 59, -1, 31, 19},
		{17, -1, 13, 19},
		{67, 7, 59, 61},
		{67, -1, 7, 59, 61},
		{67, 7, -1, 59, 61},
		{1789, 37, 47, 1889},
	}
	expected := []int{1068781, 3417, 754018, 779210, 1261476, 1202161486}

	for i, test := range tests {
		result := Task2(test)
		if result != expected[i] {
			t.Errorf("%d: got %d, expected %d", i, result, expected)
		}
	}
}
