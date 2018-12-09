package main

import "testing"

func TestSolution1(t *testing.T) {
	tests := []struct {
		players    int
		lastMarble int
		expected   int
	}{
		{9, 25, 32},
		{10, 1618, 8317},
		{13, 7999, 146373},
		{17, 1104, 2764},
		{21, 6111, 54718},
		{30, 5807, 37305},
	}

	for _, test := range tests {
		p := Puzzle{
			Players:    test.players,
			LastMarble: test.lastMarble,
			Marbles:    NewCircular(),
		}
		result := p.Solution1(false)
		if result != test.expected {
			t.Errorf("Expected %d, got %d\n", test.expected, result)
		}
	}
}

func BenchmarkSolution1(b *testing.B) {
	p := Puzzle{
		Players:    13,
		LastMarble: 7999,
		Marbles:    NewCircular(),
	}
	for i := 0; i < b.N; i++ {
		p.Solution1(false)
	}
}
