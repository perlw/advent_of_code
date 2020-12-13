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
	input := []Bus{7, 13, 19, 31, 59}
	expected := 295

	result := Task1(start, input)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
}
