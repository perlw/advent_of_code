package main

import (
	"testing"
)

func TestShouldMask(t *testing.T) {
	mask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"

	p := PortComputer{
		Mask: mask,
		Mem:  make(map[int]int),
	}
	p.Set(8, 11)
	if p.Mem[8] != 73 {
		t.Errorf("got %d, expected %d", p.Mem[8], 73)
	}
	p.Set(7, 101)
	if p.Mem[7] != 101 {
		t.Errorf("got %d, expected %d", p.Mem[7], 101)
	}
	p.Set(8, 0)
	if p.Mem[8] != 64 {
		t.Errorf("got %d, expected %d", p.Mem[8], 69)
	}
}

func TestTask1ShouldFindResult(t *testing.T) {
	input := []Op{
		{
			Type: OpMask,
			Mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
		},
		{
			Type: OpMem,
			Addr: 8, Val: 11,
		},
		{
			Type: OpMem,
			Addr: 7, Val: 101,
		},
		{
			Type: OpMem,
			Addr: 8, Val: 0,
		},
	}
	expected := 165

	result := Task1(input)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	expected := -1

	result := Task2()
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}
