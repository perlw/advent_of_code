package main

import (
	"testing"
)

func TestCircuit(t *testing.T) {
	tests := []struct {
		ops      []int
		input    []int
		expected int
	}{
		{
			ops:      []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			input:    []int{4, 3, 2, 1, 0},
			expected: 43210,
		},
		{
			ops:      []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			input:    []int{0, 1, 2, 3, 4},
			expected: 54321,
		},
		{
			ops:      []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			input:    []int{1, 0, 4, 3, 2},
			expected: 65210,
		},
	}

	for _, test := range tests {
		got, err := RunCircuit(test.ops, test.input, false)
		if err != nil {
			t.Errorf("Error while executing test: %s", err.Error())
			continue
		}
		if got != test.expected {
			t.Errorf("%+v: Got %+v, expected %+v", test.ops, got, test.expected)
		}
	}
}

func TestPhase(t *testing.T) {
	tests := []struct {
		ops      []int
		expected int
	}{
		{
			ops:      []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0},
			expected: 43210,
		},
		{
			ops:      []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0},
			expected: 54321,
		},
		{
			ops:      []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0},
			expected: 65210,
		},
	}

	for _, test := range tests {
		got, err := FindPhaseSettings(test.ops, 0, 4)
		if err != nil {
			t.Errorf("Error while executing test: %s", err.Error())
			continue
		}
		if got != test.expected {
			t.Errorf("%+v: Got %+v, expected %+v", test.ops, got, test.expected)
		}
	}
}

/*
func TestLoop(t *testing.T) {
	tests := []struct {
		ops      []int
		expected int
	}{
		{
			ops:      []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5},
			expected: 0,
		},
	}

	for _, test := range tests {
		got, err := FindPhaseSettingsLoop(test.ops)
		if err != nil {
			t.Errorf("Error while executing test: %s", err.Error())
			continue
		}
		if got != test.expected {
			t.Errorf("%+v: Got %+v, expected %+v", test.ops, got, test.expected)
		}
	}
}
*/
