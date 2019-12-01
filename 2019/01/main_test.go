package main

import "testing"

func TestBaseFuel(t *testing.T) {
	tests := map[int]int{
		12:     2,
		14:     2,
		1969:   654,
		100756: 33583,
	}

	for input, expected := range tests {
		module := Module{
			Mass: input,
		}
		b, _ := module.RequiredFuel()
		if b != expected {
			t.Errorf("Mass: %d, got %d, expected %d", input, b, expected)
		}
	}
}

func TestExtraFuel(t *testing.T) {
	tests := map[int]int{
		12:     2,
		14:     2,
		1969:   966,
		100756: 50346,
	}

	for input, expected := range tests {
		module := Module{
			Mass: input,
		}
		b, e := module.RequiredFuel()
		if b+e != expected {
			t.Errorf("Mass: %d, got %d, expected %d", input, b+e, expected)
		}
	}
}
