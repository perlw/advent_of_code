package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type MoveDir rune

const (
	MoveDirUp    MoveDir = 'U'
	MoveDirDown          = 'D'
	MoveDirLeft          = 'L'
	MoveDirRight         = 'R'
)

type Path struct {
	Sx, Sy int
	Dx, Dy int
	Ox, Oy int
}

func ToPaths(moves []string) []Path {
	var cx, cy int
	paths := make([]Path, len(moves))
	for i := range moves {
		var path Path
		path.Sx = cx
		path.Sy = cy
		path.Dx = cx
		path.Dy = cy
		path.Ox = cx
		path.Oy = cy
		dist, _ := strconv.Atoi(string(moves[i][1:]))
		switch MoveDir(moves[i][0]) {
		case MoveDirUp:
			cy += dist
			path.Dy = cy
		case MoveDirDown:
			cy -= dist
			path.Sy = cy
		case MoveDirLeft:
			cx -= dist
			path.Sx = cx
		case MoveDirRight:
			cx += dist
			path.Dx = cx
		}
		paths[i] = path
	}

	return paths
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

type Crossing struct {
	X, Y int
}

func findCrossings(pathA, pathB []Path, debug bool) []Crossing {
	crossings := make([]Crossing, 0, 10)
	for _, a := range pathA {
		for _, b := range pathB {
			var builder strings.Builder
			fmt.Fprintf(&builder, "checking %+v vs %+v... ", a, b)
			if a.Sy == a.Dy && b.Sx == b.Dx {
				fmt.Fprintf(&builder, "potential")
				if a.Sy >= b.Sy && a.Sy <= b.Dy {
					fmt.Fprintf(&builder, ", in y")
					if b.Sx >= a.Sx && b.Sx <= a.Dx {
						fmt.Fprintf(&builder, ", collision! %d %d", b.Sx, a.Sy)
						crossings = append(crossings, Crossing{
							X: b.Sx,
							Y: a.Sy,
						})
					}
				}
			} else if a.Sx == a.Dx && b.Sy == b.Dy {
				fmt.Fprintf(&builder, "potential")
				if a.Sx >= b.Sx && a.Sx <= b.Dx {
					fmt.Fprintf(&builder, ", in x")
					if b.Sy >= a.Sy && b.Sy <= a.Dy {
						fmt.Fprintf(&builder, ", collision! %d %d", a.Sx, b.Sy)
						crossings = append(crossings, Crossing{
							X: a.Sx,
							Y: b.Sy,
						})
					}
				}
			}
			if debug {
				log.Println(builder.String())
			}
		}
	}
	return crossings
}

func FindClosestCrossing(pathA, pathB []Path, debug bool) int {
	dist := 999999

	crossings := findCrossings(pathA, pathB, debug)
	for _, c := range crossings {
		if c.X != 0 && c.Y != 0 && abs(c.X)+abs(c.Y) < dist {
			dist = abs(c.X) + abs(c.Y)
		}
	}

	return dist
}

func countSteps(path []Path, crossings []Crossing) map[Crossing]int {
	steps := make(map[Crossing]int)
	for _, c := range crossings {
		for _, p := range path {
			if p.Sx == p.Dx && p.Sx == c.X && c.Y >= p.Sy && c.Y <= p.Dy {
				steps[c] += abs(c.Y - p.Oy)
				break
			} else if p.Sy == p.Dy && p.Sy == c.Y && c.X >= p.Sx && c.X <= p.Dx {
				steps[c] += abs(c.X - p.Ox)
				break
			} else {
				steps[c] += (p.Dx - p.Sx) + (p.Dy - p.Sy)
			}
		}
	}
	return steps
}

func FindShortestCrossing(pathA, pathB []Path, debug bool) int {
	crossings := findCrossings(pathA, pathB, debug)
	cA := countSteps(pathA, crossings)
	cB := countSteps(pathB, crossings)

	dist := 999999
	for _, c := range crossings {
		if c.X != 0 && c.Y != 0 && cA[c]+cB[c] < dist {
			dist = cA[c] + cB[c]
		}
	}
	return dist
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("could not open input file: %w", err))
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	var wires [][]Path
	for scanner.Scan() {
		m := strings.Split(strings.TrimSpace(scanner.Text()), ",")
		wires = append(wires, ToPaths(m))
	}
	fmt.Printf("%+v\n", wires)

	dist := FindClosestCrossing(wires[0], wires[1], false)
	fmt.Printf("Closest: %d\n", dist)

	dist = FindShortestCrossing(wires[0], wires[1], false)
	fmt.Printf("Shortest: %d\n", dist)
}
