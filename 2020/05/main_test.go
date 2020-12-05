package main

import (
	"testing"
)

func TestShouldDecode(t *testing.T) {
	test := []struct {
		Seat   Seat
		Row    int
		Column int
		ID     int
	}{
		{
			Seat:   Seat("FBFBBFFRLR"),
			Row:    44,
			Column: 5,
			ID:     357,
		},
		{
			Seat:   Seat("BFFFBBFRRR"),
			Row:    70,
			Column: 7,
			ID:     567,
		},
		{
			Seat:   Seat("FFFBBBFRRR"),
			Row:    14,
			Column: 7,
			ID:     119,
		},
		{
			Seat:   Seat("BBFFBBFRLL"),
			Row:    102,
			Column: 4,
			ID:     820,
		},
	}

	for i, in := range test {
		row, column, id := in.Seat.GetRow(), in.Seat.GetColumn(), in.Seat.GetID()
		if row != in.Row {
			t.Errorf("%d - row: got %d, expected %d", i, row, in.Row)
		}
		if column != in.Column {
			t.Errorf("%d - column: got %d, expected %d", i, column, in.Column)
		}
		if id != in.ID {
			t.Errorf("%d - id: got %d, expected %d", i, id, in.ID)
		}
	}
}

func TestTask1ShouldFindResult(t *testing.T) {
	input := []Seat{
		Seat("FBFBBFFRLR"),
		Seat("BFFFBBFRRR"),
		Seat("FFFBBBFRRR"),
		Seat("BBFFBBFRLL"),
	}
	expected := 820

	result := Task1(input)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}
