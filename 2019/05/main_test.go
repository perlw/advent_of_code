package main

import (
	"bytes"
	"io"
	"strconv"
	"testing"
)

func TestOps(t *testing.T) {
	tests := []struct {
		ops      []int
		expected []int
	}{
		{
			ops:      []int{3, 0, 4, 0, 99},
			expected: []int{1, 0, 4, 0, 99},
		},
		{
			ops:      []int{1002, 4, 3, 4, 33},
			expected: []int{1002, 4, 3, 4, 99},
		},
		{
			ops:      []int{1, 5, 6, 7, 99, 3, 4, 0},
			expected: []int{1, 5, 6, 7, 99, 3, 4, 7},
		},
		{
			ops:      []int{1101, 5, 6, 7, 99, 0, 0, 0},
			expected: []int{1101, 5, 6, 7, 99, 0, 0, 11},
		},
	}

test_loop:
	for _, test := range tests {
		got, err := ExecuteProgram(test.ops, func() int {
			return 1
		}, nil, false)
		if err != nil {
			t.Errorf("Error while executing test: %s", err.Error())
			continue test_loop
		}
		if len(got) != len(test.expected) {
			t.Errorf("Got %+v, expected %+v", got, test.expected)
			continue test_loop
		}
		for i := range got {
			if test.expected[i] != got[i] {
				t.Errorf("Got %+v, expected %+v", got, test.expected)
				continue test_loop
			}
		}
	}
}

func TestLogical(t *testing.T) {
	tests := []struct {
		ops    []int
		input  int
		output int
	}{
		{
			ops:    []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input:  8,
			output: 1,
		},
		{
			ops:    []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8},
			input:  7,
			output: 0,
		},
		{
			ops:    []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input:  7,
			output: 1,
		},
		{
			ops:    []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8},
			input:  8,
			output: 0,
		},
		{
			ops:    []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input:  8,
			output: 1,
		},
		{
			ops:    []int{3, 3, 1108, -1, 8, 3, 4, 3, 99},
			input:  7,
			output: 0,
		},
		{
			ops:    []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input:  7,
			output: 1,
		},
		{
			ops:    []int{3, 3, 1107, -1, 8, 3, 4, 3, 99},
			input:  8,
			output: 0,
		},
		{
			ops:    []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input:  0,
			output: 0,
		},
		{
			ops:    []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9},
			input:  8,
			output: 1,
		},
		{
			ops:    []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			input:  0,
			output: 0,
		},
		{
			ops:    []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1},
			input:  8,
			output: 1,
		},
		{
			ops:    []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:  1,
			output: 999,
		},
		{
			ops:    []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:  8,
			output: 1000,
		},
		{
			ops:    []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99},
			input:  10,
			output: 1001,
		},
	}

	for _, test := range tests {
		var b bytes.Buffer
		_, err := ExecuteProgram(test.ops, func() int {
			return test.input
		}, io.Writer(&b), false)
		if err != nil {
			t.Errorf("Error while executing test: %s", err.Error())
			continue
		}
		got, _ := strconv.Atoi(b.String())
		if got != test.output {
			t.Errorf("%+v: Got %+v, expected %+v", test.ops, got, test.output)
		}
	}
}
