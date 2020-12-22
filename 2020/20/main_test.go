package main

import (
	"os"
	"testing"

	"github.com/perlw/advent_of_code/toolkit/grid"
)

func readTestInput() Tiles {
	file, _ := os.Open("input_test.txt")
	defer file.Close()
	return readInput(file)
}

func TestShouldReadInput(t *testing.T) {
	input := readTestInput()
	expectedToFind := Tiles{
		1489: grid.FromRunes([]rune(
			string("##.#.#......##...#...##..##.....#...#...#####...#.#..#.#.#.#...#.#.#..##.#...##...##.##.#####.##.#.."),
		),
			10, 10,
		),
	}

	for i, e := range expectedToFind {
		if in, ok := input[i]; !ok {
			t.Errorf("did not read input correctly, missing id")
		} else {
			if len(e.Cells) != len(in.Cells) {
				t.Errorf(
					"did not read input correctly, differing lengths (%d vs %d expected)",
					len(in.Cells), len(e.Cells),
				)
			}
			if in.Width != e.Width || in.Height != e.Height {
				t.Errorf("did not read input correctly, incorrect width/height")
			}
			for j := range in.Cells {
				if in.Cells[j] != e.Cells[j] {
					t.Errorf("did not read input correctly, corrupt data")
				}
			}
		}
	}
}

func TestShouldRotate90(t *testing.T) {
	input := readTestInput()
	expectedToFind := Tiles{
		1489: grid.FromRunes([]rune(
			string("#.#.##...##.#..#.#.###...####..#####..###....#....##.##..#.#.#....##..#.###...#..##..#.....#..#....."),
		),
			10, 10,
		),
	}

	g := input[1489]
	rotateGrid(&g)

	for _, e := range expectedToFind {
		for i := range g.Cells {
			if g.Cells[i] != e.Cells[i] {
				t.Errorf("did not read input correctly, corrupt data %d %c %c", i, g.Cells[i], e.Cells[i])
				return
			}
		}
	}
}

func TestShouldVFlip(t *testing.T) {
	input := readTestInput()
	expectedToFind := Tiles{
		1489: grid.FromRunes([]rune(
			string("....#.#.##..#...##.....##..##....#...#...#...######.#.#.#..#..#.#.#....##...#.####.##.##....#.##.###"),
		),
			10, 10,
		),
	}

	g := input[1489]
	vflipGrid(&g)

	for _, e := range expectedToFind {
		for i := range g.Cells {
			if g.Cells[i] != e.Cells[i] {
				t.Errorf("did not read input correctly, corrupt data %d %c %c", i, g.Cells[i], e.Cells[i])
			}
		}
	}
}

func TestShouldHFlip(t *testing.T) {
	input := readTestInput()
	expectedToFind := Tiles{
		1489: grid.FromRunes([]rune(
			string("###.##.#....##.##.####.#...##....#.#.#..#..#.#.#.######...#...#...#....##..##.....##...#..##.#.#...."),
		),
			10, 10,
		),
	}

	g := input[1489]
	hflipGrid(&g)

	for _, e := range expectedToFind {
		for i := range g.Cells {
			if g.Cells[i] != e.Cells[i] {
				t.Errorf("did not read input correctly, corrupt data %d %c %c", i, g.Cells[i], e.Cells[i])
			}
		}
	}
}

func TestShouldPerformOperations(t *testing.T) {
	input := readTestInput()
	expectedToFind := Tiles{
		1489: grid.FromRunes([]rune(
			string("##.#.#......##...#...##..##.....#...#...#####...#.#..#.#.#.#...#.#.#..##.#...##...##.##.#####.##.#.."),
		),
			10, 10,
		),
	}

	g := input[1489]
	hflipGrid(&g)
	rotateGrid(&g)
	rotateGrid(&g)
	vflipGrid(&g)

	for _, e := range expectedToFind {
		for i := range g.Cells {
			if g.Cells[i] != e.Cells[i] {
				t.Errorf("did not read input correctly, corrupt data %d %c %c", i, g.Cells[i], e.Cells[i])
				return
			}
		}
	}
}

func TestTask1ShouldFindResult(t *testing.T) {
	input := readTestInput()
	expect := 20899048083289

	result := Task1(input)
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}

func TestTask2ShouldFindResult(t *testing.T) {
	expect := -1

	result := Task2()
	if result != expect {
		t.Errorf("got %d, expected %d", result, expect)
	}
}
