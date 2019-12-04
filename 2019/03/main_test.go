package main

import (
	"fmt"
	"testing"
)

func printPaths(pathA []Path, pathB []Path) {
	minX := 999999
	minY := 999999
	maxX := -999999
	maxY := -999999

	for _, p := range pathA {
		if p.Sx < minX {
			minX = p.Sx
		}
		if p.Sy < minY {
			minY = p.Sy
		}
		if p.Dx > maxX {
			maxX = p.Dx
		}
		if p.Dy > maxY {
			maxY = p.Dy
		}
	}
	for _, p := range pathB {
		if p.Sx < minX {
			minX = p.Sx
		}
		if p.Sy < minY {
			minY = p.Sy
		}
		if p.Dx > maxX {
			maxX = p.Dx
		}
		if p.Dy > maxY {
			maxY = p.Dy
		}
	}

	gridX := maxX - minX + 3
	gridY := maxY - minY + 3
	offX := -minX
	offY := -minY
	grid := make([]rune, gridX*gridY)
	for i := range grid {
		grid[i] = '.'
	}

	for _, p := range pathA {
		if p.Sx == p.Dx {
			for y := p.Sy; y <= p.Dy; y++ {
				i := ((y + offY + 1) * gridX) + p.Sx + offX + 1
				grid[i] = 'a'
			}
		} else if p.Sy == p.Dy {
			for x := p.Sx; x <= p.Dx; x++ {
				i := ((p.Sy + offY + 1) * gridX) + x + offX + 1
				grid[i] = 'a'
			}
		}
	}
	for _, p := range pathB {
		if p.Sx == p.Dx {
			for y := p.Sy; y <= p.Dy; y++ {
				i := ((y + offY + 1) * gridX) + p.Sx + offX + 1
				grid[i] = 'b'
			}
		} else if p.Sy == p.Dy {
			for x := p.Sx; x <= p.Dx; x++ {
				i := ((p.Sy + offY + 1) * gridX) + x + offX + 1
				grid[i] = 'b'
			}
		}
	}
	grid[((offY+1)*gridX)+offX+1] = 'o'

	for y := gridY - 1; y >= 0; y-- {
		i := y * gridX
		for x := 0; x < gridX; x++ {
			fmt.Printf("%c", grid[i+x])
		}
		fmt.Println()
	}
}

func TestPath(t *testing.T) {
	tests := []struct {
		pathA    []string
		pathB    []string
		expected int
	}{
		{
			pathA:    []string{"R8", "U5", "L5", "D3"},
			pathB:    []string{"U7", "R6", "D4", "L4"},
			expected: 6,
		},
		{
			pathA:    []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			pathB:    []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			expected: 159,
		},
		{
			pathA:    []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			pathB:    []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			expected: 135,
		},
	}

	for i, test := range tests {
		a := ToPaths(test.pathA)
		b := ToPaths(test.pathB)
		printPaths(a, b)
		got := FindClosestCrossing(a, b, false)
		if got != test.expected {
			t.Errorf("Test %d: Got %d, expected %d", i, got, test.expected)
		}
	}
}

func TestLength(t *testing.T) {
	tests := []struct {
		pathA    []string
		pathB    []string
		expected int
	}{
		{
			pathA:    []string{"R8", "U5", "L5", "D3"},
			pathB:    []string{"U7", "R6", "D4", "L4"},
			expected: 30,
		},
		{
			pathA:    []string{"R75", "D30", "R83", "U83", "L12", "D49", "R71", "U7", "L72"},
			pathB:    []string{"U62", "R66", "U55", "R34", "D71", "R55", "D58", "R83"},
			expected: 610,
		},
		{
			pathA:    []string{"R98", "U47", "R26", "D63", "R33", "U87", "L62", "D20", "R33", "U53", "R51"},
			pathB:    []string{"U98", "R91", "D20", "R16", "D67", "R40", "U7", "R15", "U6", "R7"},
			expected: 410,
		},
	}

	for i, test := range tests {
		a := ToPaths(test.pathA)
		b := ToPaths(test.pathB)
		printPaths(a, b)
		got := FindShortestCrossing(a, b, false)
		if got != test.expected {
			t.Errorf("Test %d: Got %d, expected %d", i, got, test.expected)
		}
	}
}
