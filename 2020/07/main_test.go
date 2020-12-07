package main

import (
	"testing"
)

var input = [][]Rule{
	{
		{
			Color: "light red",
			Contains: map[string]int{
				"bright white": 1,
				"muted yellow": 2,
			},
		},
		{
			Color: "dark orange",
			Contains: map[string]int{
				"bright white": 3,
				"muted yellow": 4,
			},
		},
		{
			Color: "bright white",
			Contains: map[string]int{
				"shiny gold": 1,
			},
		},
		{
			Color: "muted yellow",
			Contains: map[string]int{
				"shiny gold": 2,
				"faded blue": 9,
			},
		},
		{
			Color: "shiny gold",
			Contains: map[string]int{
				"dark olive":   1,
				"vibrant plum": 2,
			},
		},
		{
			Color: "dark olive",
			Contains: map[string]int{
				"faded blue":   3,
				"dotted black": 4,
			},
		},
		{
			Color: "vibrant plum",
			Contains: map[string]int{
				"faded blue":   5,
				"dotted black": 6,
			},
		},
		{
			Color:    "faded blue",
			Contains: map[string]int{},
		},
		{
			Color:    "dotted black",
			Contains: map[string]int{},
		},
	},
	{
		{
			Color: "shiny gold",
			Contains: map[string]int{
				"dark red": 2,
			},
		},
		{
			Color: "dark red",
			Contains: map[string]int{
				"dark orange": 2,
			},
		},
		{
			Color: "dark orange",
			Contains: map[string]int{
				"dark yellow": 2,
			},
		},
		{
			Color: "dark yellow",
			Contains: map[string]int{
				"dark green": 2,
			},
		},
		{
			Color: "dark green",
			Contains: map[string]int{
				"dark blue": 2,
			},
		},
		{
			Color: "dark blue",
			Contains: map[string]int{
				"dark violet": 2,
			},
		},
		{
			Color:    "dark violet",
			Contains: map[string]int{},
		},
	},
}

func TestShouldFindContainingColors(t *testing.T) {
	expected := [][]string{
		{"bright white", "muted yellow", "dark orange", "light red"},
		{},
	}

	for i := range input {
		result := findContainingColors(input[i], "shiny gold")
		if len(result) != len(expected[i]) {
			t.Errorf("%d: got %+v, expected %+v", i, result, expected[i])
		}
		for _, r := range result {
			var found bool
			for _, e := range expected[i] {
				if r == e {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("%d: got %+v, expected %+v", i, result, expected[i])
				break
			}
		}
	}
}

func TestShouldFindColorsContained(t *testing.T) {
	expected := []map[string]int{
		{
			"faded blue":   13,
			"dotted black": 16,
			"dark olive":   1,
			"vibrant plum": 2,
		},
		{
			"dark blue":   32,
			"dark green":  16,
			"dark orange": 4,
			"dark red":    2,
			"dark violet": 64,
			"dark yellow": 8,
		},
	}

	for i := range input {
		result := findColorsContained(input[i], "shiny gold")
		if len(result) != len(expected[i]) {
			t.Errorf("%d: got %+v, expected %+v", i, result, expected[i])
		}
		for k, v := range result {
			if expected[i][k] != v {
				t.Errorf("%d: got %+v, expected %+v", i, result, expected[i])
				break
			}
		}
	}
}

func TestTask1ShouldFindResult(t *testing.T) {
	expected := []int{4, 0}

	for i := range input {
		result := Task1(input[i])
		if result != expected[i] {
			t.Errorf("%d: got %d, expected %d", i, result, expected)
		}
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	expected := []int{32, 126}

	for i := range input {
		result := Task2(input[i])
		if result != expected[i] {
			t.Errorf("%d: got %d, expected %d", i, result, expected)
		}
	}
}
