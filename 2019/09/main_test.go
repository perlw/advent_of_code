package main

import (
	"testing"
)

func TestRelative(t *testing.T) {
	tests := []struct {
		ops      []int
		expected []int
	}{
		{
			ops:      []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
			expected: []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99},
		},
		{
			ops:      []int{1102, 34915192, 34915192, 7, 4, 7, 99, 0},
			expected: []int{1219070632396864},
		},
		{
			ops:      []int{104, 1125899906842624, 99},
			expected: []int{1125899906842624},
		},
		{
			ops:      []int{109, 1, 204, -1, 99},
			expected: []int{109},
		},
	}

	for _, test := range tests {
		com := NewComputer(test.ops)
		got := make([]int, 0)
		for !com.Halted {
			out, halt, err := com.Iterate(func() int {
				return 0
			}, false, false)
			if err != nil {
				t.Errorf("Error while executing test: %s", err.Error())
				break
			}
			if !halt {
				got = append(got, out)
			}
		}
		if len(got) != len(test.expected) {
			t.Errorf("%+v: Got %+v, expected %+v (incorrect lengths)", test.ops, got, test.expected)
			break
		}
		for i := range got {
			if got[i] != test.expected[i] {
				t.Errorf("%+v: Got %+v, expected %+v", test.ops, got, test.expected)
				break
			}
		}
	}
}
