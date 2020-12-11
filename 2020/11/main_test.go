package main

import (
	"testing"
)

func TestRayShouldFindOccupiedSeats(t *testing.T) {
	seatMap := SeatMap{
		'#', '.', '.', '#', '.', '.', '#',
		'.', '#', '.', '.', '.', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'#', '.', '.', 'L', '.', '.', '#',
		'.', '.', '.', '.', '#', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'#', '.', '.', '#', '.', '.', '#',
	}
	tests := [][2]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
	}

	for _, test := range tests {
		if _, ok := seatMap.ray(7, 3, 3, test); !ok {
			t.Errorf("%+v did not find occupied seat", test)
		}
	}
}

func TestRayShouldNotFindOccupiedSeat(t *testing.T) {
	seatMap := SeatMap{
		'#', '.', '.', '#', '.', '.', '#',
		'.', 'L', '.', '.', '.', 'L', '.',
		'.', '.', '.', 'L', '.', '.', '.',
		'.', '.', '.', 'L', '.', '.', '.',
		'.', '.', 'L', '.', 'L', '.', '.',
		'.', '.', '.', '.', '.', '.', '.',
		'#', '.', '.', '.', '.', '.', '#',
	}
	tests := [][2]int{
		{-1, -1},
		{0, -1},
		{1, -1},
		{1, 0},
		{1, 1},
		{0, 1},
		{-1, 1},
		{-1, 0},
	}

	for _, test := range tests {
		if _, ok := seatMap.ray(7, 3, 3, test); ok {
			t.Errorf("%+v found occupied seat", test)
		}
	}
}

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
	expected := 26

	result := Task2(seatMap, 10, !testing.Short())
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}
