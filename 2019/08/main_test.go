package main

import "testing"

func TestChecksum(t *testing.T) {
	tests := []struct {
		img      []int
		expected int
	}{
		{
			img:      []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
			expected: 1,
		},
	}

	for _, test := range tests {
		got := CalculateChecksum(test.img, 3, 2)
		if got != test.expected {
			t.Errorf("Got %+v, expected %+v", got, test.expected)
		}
	}
}

func TestFlatten(t *testing.T) {
	tests := []struct {
		img      []int
		expected []int
	}{
		{
			img:      []int{0, 2, 2, 2, 1, 1, 2, 2, 2, 2, 1, 2, 0, 0, 0, 0},
			expected: []int{0, 1, 1, 0},
		},
	}

	for _, test := range tests {
		got := FlattenImage(test.img, 2, 2)
		for i := range got {
			if got[i] != test.expected[i] {
				t.Errorf("Got %+v, expected %+v", got, test.expected)
				break
			}
		}
	}
}
