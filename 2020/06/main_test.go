package main

import (
	"testing"
)

func TestShouldCountUniqueAnswers(t *testing.T) {
	test := []struct {
		Answers []string
		Expect  int
	}{
		{
			Answers: []string{"abc"},
			Expect:  3,
		},
		{
			Answers: []string{"a", "b", "c"},
			Expect:  3,
		},
		{
			Answers: []string{"ab", "ac"},
			Expect:  3,
		},
		{
			Answers: []string{"a", "a", "a", "a"},
			Expect:  1,
		},
		{
			Answers: []string{"b"},
			Expect:  1,
		},
	}

	for i, in := range test {
		result := countUniqueAnswers(in.Answers)
		if result != in.Expect {
			t.Errorf("%d: got %d, expected %d", i, result, in.Expect)
		}
	}
}

func TestShouldCountAllSharedAnswers(t *testing.T) {
	test := []struct {
		Answers []string
		Expect  int
	}{
		{
			Answers: []string{"abc"},
			Expect:  3,
		},
		{
			Answers: []string{"a", "b", "c"},
			Expect:  0,
		},
		{
			Answers: []string{"ab", "ac"},
			Expect:  1,
		},
		{
			Answers: []string{"a", "a", "a", "a"},
			Expect:  1,
		},
		{
			Answers: []string{"b"},
			Expect:  1,
		},
	}

	for i, in := range test {
		result := countAllSharedAnswers(in.Answers)
		if result != in.Expect {
			t.Errorf("%d: got %d, expected %d", i, result, in.Expect)
		}
	}
}

func TestTask1ShouldFindResult(t *testing.T) {
	input := [][]string{
		{"abc"},
		{"a", "b", "c"},
		{"ab", "ac"},
		{"a", "a", "a", "a"},
		{"b"},
	}
	expected := 11

	result := Task1(input)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	input := [][]string{
		{"abc"},
		{"a", "b", "c"},
		{"ab", "ac"},
		{"a", "a", "a", "a"},
		{"b"},
	}
	expected := 6

	result := Task2(input)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}
