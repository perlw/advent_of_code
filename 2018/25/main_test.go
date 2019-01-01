package main

import (
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		filepath string
		expected int
	}{
		{"test1.txt", 2},
		{"test2.txt", 4},
		{"test3.txt", 3},
		{"test4.txt", 8},
	}

	for _, test := range tests {
		p := NewPuzzle(test.filepath)
		if result := p.Part1(); result != test.expected {
			t.Errorf("Expected %d, got %d\n", test.expected, result)
		}
	}
}
