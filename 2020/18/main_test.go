package main

import (
	"fmt"
	"testing"
)

func TestShouldReduceExpression(t *testing.T) {
	tests := []string{
		"1+2*3+4*5+6",
		"2*3+(4*5)",
		"1+(2*3)+(4*(5+6))",
		"(2*3)+(4*5)",
		"5+(8*3+9+3*4*3)",
		"5*9*(7*3*3+9*3+(8+6*4))",
		"((2+4*9)*(6+9*8+6)+6)+2+4*2",
	}
	expected := []int{
		71,
		26,
		51,
		26,
		437,
		12240,
		13632,
	}

	for i, test := range tests {
		result := reduce([]rune(test), true)
		fmt.Printf("[[%s => %d]]\n", test, result)
		fmt.Println()
		if result != expected[i] {
			t.Errorf("%d: got %d, expected %d", i, result, expected[i])
		}
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	expected := -1

	result := Task2()
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}
