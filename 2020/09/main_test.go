package main

import (
	"testing"
)

func TestCheckNumber(t *testing.T) {
	tests := []struct {
		Input    []int
		Number   int
		Expected [][2]int
	}{
		{
			Input: []int{
				35,
				20,
				15,
				25,
				47,
			},
			Number: 40,
			Expected: [][2]int{
				{15, 25},
			},
		},
		{
			Input: []int{
				95,
				102,
				117,
				150,
				182,
			},
			Number:   127,
			Expected: [][2]int{},
		},
		{
			Input: []int{
				65,
				95,
				102,
				117,
				150,
			},
			Number: 182,
			Expected: [][2]int{
				{65, 117},
			},
		},
	}

	for i, test := range tests {
		result := checkNumber(test.Input, test.Number)
		if len(result) != len(test.Expected) {
			t.Errorf("%d: got %+v, expected %+v", i, result, test.Expected)
		}
		for j := range result {
			if result[j] != test.Expected[j] {
				t.Errorf("%d got %+v, expected %+v", i, result, test.Expected)
			}
		}
	}
}

func TestTask1ShouldFindResult(t *testing.T) {
	input := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	expected := 127

	result := Task1(input, 5)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	input := []int{
		35,
		20,
		15,
		25,
		47,
		40,
		62,
		55,
		65,
		95,
		102,
		117,
		150,
		182,
		127,
		219,
		299,
		277,
		309,
		576,
	}
	expected := 62

	result := Task2(input, 127)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}
