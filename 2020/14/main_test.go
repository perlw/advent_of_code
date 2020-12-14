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

func TestShouldMaskZeroAddr(t *testing.T) {
	tests := []string{
		"000000000000000000000000000000000000",
		"000000000000000000000000000000000001",
		"000000000000000000000000000000000010",
		"000000000000000000000000000000000011",

		"00000000000000000000000000000000000X",
		"0000000000000000000000000000000000X0",
		"0000000000000000000000000000000000XX",
		"000000000000000000000000000000000XXX",

		"00000000000000000000000000000000001X",
		"0000000000000000000000000000000000X1",
		"000000000000000000000000000000000X0X",
		"000000000000000000000000000000000X1X",
		"000000000000000000000000000000000XX1",
		"0000000000000000000000000000000001XX",
	}
	expected := [][]int{
		{0},
		{1},
		{2},
		{3},

		{0, 1},
		{0, 2},
		{0, 1, 2, 3},
		{0, 1, 2, 3, 4, 5, 6},

		{2, 3},
		{1, 3},
		{0, 1, 5, 4},
		{2, 3, 7, 6},
		{1, 3, 7, 5},
		{4, 5, 7, 6},
	}

	for i, test := range tests {
		p := PortComputer{
			Mask: test,
			Mem:  make(map[int]int),
		}
		p.SetV2(0, 42)

		for _, e := range expected[i] {
			if p.Mem[e] != 42 {
				t.Errorf("%d: got %#v, expected value in %d", i, p.Mem, expected[i])
			}
		}
	}
}

func TestShouldMaskAddr(t *testing.T) {
	tests := []string{
		"000000000000000000000000000000000000",
		"000000000000000000000000000000000001",
		"000000000000000000000000000000000010",
		"000000000000000000000000000000000011",

		"00000000000000000000000000000000000X",
		"0000000000000000000000000000000000X0",
		"0000000000000000000000000000000000X1",
		"000000000000000000000000000000000X1X",
	}
	expected := [][]int{
		{2},
		{3},
		{2},
		{3},

		{2, 3},
		{2},
		{1, 3},
		{2, 3, 7, 6},
	}

	for i, test := range tests {
		p := PortComputer{
			Mask: test,
			Mem:  make(map[int]int),
		}
		p.SetV2(2, 42)

		for _, e := range expected[i] {
			if p.Mem[e] != 42 {
				t.Errorf("%d: got %#v, expected value in %d", i, p.Mem, expected[i])
			}
		}
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
	input := []Op{
		{
			Type: OpMask,
			Mask: "000000000000000000000000000000X1001X",
		},
		{
			Type: OpMem,
			Addr: 42, Val: 100,
		},
		{
			Type: OpMask,
			Mask: "00000000000000000000000000000000X0XX",
		},
		{
			Type: OpMem,
			Addr: 26, Val: 1,
		},
	}
	expected := 208

	result := Task2(input)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}
