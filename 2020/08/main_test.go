package main

import (
	"testing"
)

func TestShouldIncreaseAccumulator(t *testing.T) {
	elf := ElfCPU{
		Ops: []Op{
			{InstrAcc, 42},
		},
	}
	expected := 42

	elf.Step()
	if elf.Acc != expected {
		t.Errorf("got %d, expected %d", elf.Acc, expected)
	}
}

func TestShouldJump(t *testing.T) {
	elf := ElfCPU{
		Ops: []Op{
			{InstrAcc, 42},
			{InstrJmp, 2},
			{InstrAcc, -42},
			{InstrNop, 0},
		},
	}
	expected := 42

	elf.Step()
	elf.Step()
	elf.Step()

	if elf.Acc != expected {
		t.Errorf("got %d, expected %d", elf.Acc, expected)
	}
}

func TestTask1ShouldFindResult(t *testing.T) {
	elf := ElfCPU{
		Ops: []Op{
			{InstrNop, 0},
			{InstrAcc, 1},
			{InstrJmp, 4},
			{InstrAcc, 3},
			{InstrJmp, -3},
			{InstrAcc, -99},
			{InstrAcc, +1},
			{InstrJmp, -4},
			{InstrAcc, 6},
		},
	}
	expected := 5

	result := Task1(&elf)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	elf := ElfCPU{
		Ops: []Op{
			{InstrNop, 0},
			{InstrAcc, 1},
			{InstrJmp, 4},
			{InstrAcc, 3},
			{InstrJmp, -3},
			{InstrAcc, -99},
			{InstrAcc, +1},
			{InstrJmp, -4},
			{InstrAcc, 6},
		},
	}
	expected := 8

	result := Task2(&elf)
	if result != expected {
		t.Errorf("got %d, expected %d", result, expected)
	}
}
