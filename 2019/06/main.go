package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Graph struct {
	Satellites []string
	Orbits     map[string][]string
}

func (g *Graph) Add(a, b string) {
	found := false
	for _, aa := range g.Satellites {
		if aa == a {
			found = true
		}
	}
	if !found {
		g.Satellites = append(g.Satellites, a)
	}
	if g.Orbits == nil {
		g.Orbits = make(map[string][]string)
	}
	g.Orbits[a] = append(g.Orbits[a], b)
}

func (g Graph) String() string {
	result := ""
	for _, a := range g.Satellites {
		result += fmt.Sprintf("%s ->", a)
		if _, ok := g.Orbits[a]; ok {
			for _, b := range g.Orbits[a] {
				result += fmt.Sprintf(" %s", b)
			}
		}
		result += fmt.Sprintf("\n")
	}
	return result
}

func CountOrbits(g Graph, a string, fromRoot int) int {
	indirect := fromRoot
	if _, ok := g.Orbits[a]; ok {
		for _, b := range g.Orbits[a] {
			indirect += CountOrbits(g, b, fromRoot+1)
		}
	}
	return indirect
}

func findNode(g Graph, parent, target string) ([]string, bool) {
	if parent == target {
		return nil, true
	}

	for _, a := range g.Orbits[parent] {
		if v, found := findNode(g, a, target); found {
			return append(v, a), found
		}
	}

	return nil, false
}

func CalcTransfers(g Graph, a, b string) int {
	aMoves, _ := findNode(g, "COM", a)
	bMoves, _ := findNode(g, "COM", b)

	for i, a := range aMoves {
		for j, b := range bMoves {
			if a == b {
				return i + j - 2
			}
		}
	}

	return -1
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("could not open input file: %w", err))
	}
	defer input.Close()

	var system Graph
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var orbit string
		fmt.Sscanf(scanner.Text(), "%s", &orbit)
		o := strings.Split(orbit, ")")
		system.Add(o[0], o[1])
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(fmt.Errorf("error while scanning input: %w", err))
	}

	fmt.Printf("Part 1: %d\n", CountOrbits(system, "COM", 0))
	fmt.Printf("Part 2: %d\n", CalcTransfers(system, "YOU", "SAN"))
}
