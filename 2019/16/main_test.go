package main

import (
	"testing"
)

func TestFFT(t *testing.T) {
	tests := []struct {
		input    []int
		phases   int
		expected []int
	}{
		{
			input: []int{
				1, 2, 3, 4, 5, 6, 7, 8,
			},
			phases: 4,
			expected: []int{
				0, 1, 0, 2, 9, 4, 9, 8,
			},
		},
		{
			input: []int{
				8, 0, 8, 7, 1, 2, 2, 4, 5, 8, 5, 9, 1, 4, 5, 4, 6, 6, 1, 9, 0, 8, 3, 2, 1, 8, 6, 4, 5, 5, 9, 5,
			},
			phases: 100,
			expected: []int{
				2, 4, 1, 7, 6, 1, 7, 6,
			},
		},
		{
			input: []int{
				1, 9, 6, 1, 7, 8, 0, 4, 2, 0, 7, 2, 0, 2, 2, 0, 9, 1, 4, 4, 9, 1, 6, 0, 4, 4, 1, 8, 9, 9, 1, 7,
			},
			phases: 100,
			expected: []int{
				7, 3, 7, 4, 5, 4, 1, 8,
			},
		},
		{
			input: []int{
				6, 9, 3, 1, 7, 1, 6, 3, 4, 9, 2, 9, 4, 8, 6, 0, 6, 3, 3, 5, 9, 9, 5, 9, 2, 4, 3, 1, 9, 8, 7, 3,
			},
			phases: 100,
			expected: []int{
				5, 2, 4, 3, 2, 1, 3, 3,
			},
		},
	}

	for i, test := range tests {
		if testing.Short() && i > 0 {
			break
		}
		got := FFT(test.input, test.phases, false)
		for j := range got[:8] {
			if got[j] != test.expected[j] {
				t.Errorf("%d: Got %+v, expected %+v", i, got, test.expected)
				break
			}
		}
	}
}

func TestFFT2(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input: []int{
				0, 3, 0, 3, 6, 7, 3, 2, 5, 7, 7, 2, 1, 2, 9, 4, 4, 0, 6, 3, 4, 9, 1, 5, 6, 5, 4, 7, 4, 6, 6, 4,
			},
			expected: []int{
				8, 4, 4, 6, 2, 0, 2, 6,
			},
		},
		{
			input: []int{
				0, 2, 9, 3, 5, 1, 0, 9, 6, 9, 9, 9, 4, 0, 8, 0, 7, 4, 0, 7, 5, 8, 5, 4, 4, 7, 0, 3, 4, 3, 2, 3,
			},
			expected: []int{
				7, 8, 7, 2, 5, 2, 7, 0,
			},
		},
		{
			input: []int{
				0, 3, 0, 8, 1, 7, 7, 0, 8, 8, 4, 9, 2, 1, 9, 5, 9, 7, 3, 1, 1, 6, 5, 4, 4, 6, 8, 5, 0, 5, 1, 7,
			},
			expected: []int{
				5, 3, 5, 5, 3, 7, 3, 1,
			},
		},
	}

	for i, test := range tests {
		if testing.Short() && i > 0 {
			break
		}

		offset := digitsToInt(test.input[:7])
		got := FFT2(test.input, 100, offset, false)
		for j := range got[:8] {
			if got[j] != test.expected[j] {
				t.Errorf("%d: Got %+v, expected %+v", i, got[offset-8:offset+8], test.expected)
				break
			}
		}
	}
}
