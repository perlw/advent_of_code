package main

import (
	"testing"
)

func TestMotion(t *testing.T) {
	tests := []struct {
		steps    int
		moons    []Moon
		expected []Moon
	}{
		{
			steps: 10,
			moons: []Moon{
				{
					x: -1, y: 0, z: 2,
				},
				{
					x: 2, y: -10, z: -7,
				},
				{
					x: 4, y: -8, z: 8,
				},
				{
					x: 3, y: 5, z: -1,
				},
			},
			expected: []Moon{
				{
					x: 2, y: 1, z: -3,
					vx: -3, vy: -2, vz: 1,
				},
				{
					x: 1, y: -8, z: 0,
					vx: -1, vy: 1, vz: 3,
				},
				{
					x: 3, y: -6, z: 1,
					vx: 3, vy: 2, vz: -3,
				},
				{
					x: 2, y: 0, z: 4,
					vx: 1, vy: -1, vz: -1,
				},
			},
		},
		{
			steps: 100,
			moons: []Moon{
				{
					x: -8, y: -10, z: 0,
				},
				{
					x: 5, y: 5, z: 10,
				},
				{
					x: 2, y: -7, z: 3,
				},
				{
					x: 9, y: -8, z: -3,
				},
			},
			expected: []Moon{
				{
					x: 8, y: -12, z: -9,
					vx: -7, vy: 3, vz: 0,
				},
				{
					x: 13, y: 16, z: -3,
					vx: 3, vy: -11, vz: -5,
				},
				{
					x: -29, y: -11, z: -1,
					vx: -3, vy: 7, vz: 4,
				},
				{
					x: 16, y: -13, z: 23,
					vx: 7, vy: 1, vz: 1,
				},
			},
		},
	}

	for i, test := range tests {
		IterateMoons(test.moons, test.steps, true)
		for j := range test.moons {
			if test.moons[j] != test.expected[j] {
				t.Errorf("%d: Got %+v, expected %+v", i, test.moons, test.expected)
			}
		}
	}
}

func TestEnergy(t *testing.T) {
	tests := []struct {
		steps    int
		moons    []Moon
		expected int
	}{
		{
			steps: 10,
			moons: []Moon{
				{
					x: -1, y: 0, z: 2,
				},
				{
					x: 2, y: -10, z: -7,
				},
				{
					x: 4, y: -8, z: 8,
				},
				{
					x: 3, y: 5, z: -1,
				},
			},
			expected: 179,
		},
		{
			steps: 100,
			moons: []Moon{
				{
					x: -8, y: -10, z: 0,
				},
				{
					x: 5, y: 5, z: 10,
				},
				{
					x: 2, y: -7, z: 3,
				},
				{
					x: 9, y: -8, z: -3,
				},
			},
			expected: 1940,
		},
	}

	for i, test := range tests {
		got := IterateMoons(test.moons, test.steps, false)
		if got != test.expected {
			t.Errorf("%d: Got %+v, expected %+v", i, got, test.expected)
		}
	}
}

func TestSteps(t *testing.T) {
	tests := []struct {
		steps    int
		moons    []Moon
		expected int
	}{
		{
			steps: 10,
			moons: []Moon{
				{
					x: -1, y: 0, z: 2,
				},
				{
					x: 2, y: -10, z: -7,
				},
				{
					x: 4, y: -8, z: 8,
				},
				{
					x: 3, y: 5, z: -1,
				},
			},
			expected: 2772,
		},
		{
			steps: 100,
			moons: []Moon{
				{
					x: -8, y: -10, z: 0,
				},
				{
					x: 5, y: 5, z: 10,
				},
				{
					x: 2, y: -7, z: 3,
				},
				{
					x: 9, y: -8, z: -3,
				},
			},
			expected: 4686774924,
		},
	}

	for i, test := range tests {
		got := FindSteps(test.moons, true)
		if got != test.expected {
			t.Errorf("%d: Got %+v, expected %+v", i, got, test.expected)
		}
	}
}
