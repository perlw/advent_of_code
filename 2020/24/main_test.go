package main

import (
	"os"
	"testing"
)

func readTestInput() []string {
	file, _ := os.Open("input_test.txt")
	defer file.Close()
	return readInput(file)
}

func TestShouldReadInput(t *testing.T) {
	input := readTestInput()
	expectedToFind := "seswneswswsenwwnwse"

	var found bool
	for _, i := range input {
		if expectedToFind != i {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("did not read input correctly, did not match")
	}
}

func TestShouldFlipTiles(t *testing.T) {
	input := []string{
		"esew",
		"nwwswee",
		"eesw",
		"eeswse",
	}
	expect := []Pos{
		{0, 1},
		{0, 0},
		{1, 1},
		{2, 2},
	}
	for i := range input {
		hg := HexGrid{
			Root:  &Cell{},
			Cells: make(map[Pos]*Cell),
		}
		hg.Cells[Pos{0, 0}] = hg.Root
		hg.MakeMoves(input[i], true)
		if !hg.Cells[expect[i]].Flipped {
			t.Errorf("got %v, expected true", hg.Cells[expect[i]].Flipped)
		}
	}
}

func TestShouldIterateLife(t *testing.T) {
	input := readTestInput()
	expect := 15

	hg := HexGrid{
		Cells: make(map[Pos]*Cell),
	}
	hg.Cells[Pos{0, 0}] = &Cell{}
	for _, in := range input {
		hg.MakeMoves(in, false)
	}

	hg.IterateLife(true)
	var result int
	for _, c := range hg.Cells {
		if c.Flipped {
			result++
		}
	}
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}

func TestTask1ShouldFindResult(t *testing.T) {
	input := readTestInput()
	expect := 10

	result := Task1(input, true)
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	input := readTestInput()
	expect := 37

	result := Task2(input, 10, true)
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}
