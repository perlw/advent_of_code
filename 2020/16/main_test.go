package main

import (
	"testing"
)

func TestTask1ShouldFindResult(t *testing.T) {
	rules := []Rule{
		{
			Name:   "class",
			Ranges: [][2]int{{1, 3}, {5, 7}},
		},
		{
			Name:   "rules",
			Ranges: [][2]int{{6, 11}, {33, 44}},
		},
		{
			Name:   "seat",
			Ranges: [][2]int{{13, 40}, {45, 50}},
		},
	}
	tickets := []Ticket{
		{
			Values: []int{7, 3, 47},
		},
		{
			Values: []int{40, 4, 50},
		},
		{
			Values: []int{55, 2, 20},
		},
		{
			Values: []int{38, 6, 12},
		},
	}
	expected := 71

	result := Task1(rules, tickets)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
}
