package main

import "testing"

func TestProgram(t *testing.T) {
	tests := []struct {
		ops      []int
		expected []int
	}{
		{
			ops:      []int{1, 0, 0, 0, 99},
			expected: []int{2, 0, 0, 0, 99},
		},
		{
			ops:      []int{2, 3, 0, 3, 99},
			expected: []int{2, 3, 0, 6, 99},
		},
		{
			ops:      []int{2, 4, 4, 5, 99, 0},
			expected: []int{2, 4, 4, 5, 99, 9801},
		},
		{
			ops:      []int{1, 1, 1, 4, 99, 5, 6, 0, 99},
			expected: []int{30, 1, 1, 4, 2, 5, 6, 0, 99},
		},
	}

test_loop:
	for _, test := range tests {
		got, err := ExecuteProgram(test.ops)
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
