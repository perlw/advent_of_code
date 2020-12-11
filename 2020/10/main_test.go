package main

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func TestTask1ShouldFindResult(t *testing.T) {
	input := []int{
		28,
		33,
		18,
		42,
		31,
		14,
		46,
		20,
		48,
		47,
		24,
		23,
		49,
		45,
		19,
		38,
		39,
		11,
		1,
		32,
		25,
		35,
		8,
		17,
		7,
		9,
		4,
		2,
		34,
		10,
		3,
	}
	expected := 220

	result := Task1(input)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	tests := [][]int{
		{
			16,
			10,
			15,
			5,
			1,
			11,
			7,
			19,
			6,
			12,
			4,
		}, {
			28,
			33,
			18,
			42,
			31,
			14,
			46,
			20,
			48,
			47,
			24,
			23,
			49,
			45,
			19,
			38,
			39,
			11,
			1,
			32,
			25,
			35,
			8,
			17,
			7,
			9,
			4,
			2,
			34,
			10,
			3,
		},
	}
	expected := []int{8, 19208}

	for i, test := range tests {
		result := Task2(test)
		if result != expected[i] {
			t.Errorf("%d: got %d, expected %d", i, result, expected)
		}
	}
}

func BenchmarkTask2(b *testing.B) {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	input := make([]int, 0, 100)
	for scanner.Scan() {
		val, _ := strconv.Atoi(scanner.Text())
		input = append(input, val)
	}
	file.Close()

	for i := 0; i < b.N; i++ {
		Task2(input)
	}
}
