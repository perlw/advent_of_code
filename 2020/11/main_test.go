package main

import (
	"testing"
)

func TestTask1ShouldFindResult(t *testing.T) {
	seatMap := SeatMap{
		'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L',
		'L', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L',
		'L', '.', 'L', '.', 'L', '.', '.', 'L', '.', '.',
		'L', 'L', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L',
		'L', '.', 'L', 'L', '.', 'L', 'L', '.', 'L', 'L',
		'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L',
		'.', '.', 'L', '.', 'L', '.', '.', '.', '.', '.',
		'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L', 'L',
		'L', '.', 'L', 'L', 'L', 'L', 'L', 'L', '.', 'L',
		'L', '.', 'L', 'L', 'L', 'L', 'L', '.', 'L', 'L',
	}
	expected := 37

	result := Task1(seatMap, 10, !testing.Short())
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
}
