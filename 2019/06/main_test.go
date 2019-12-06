package main

import (
	"fmt"
	"strings"
	"testing"
)

func TestBasicGraph(t *testing.T) {
	test := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
	}

	var g Graph
	for _, a := range test {
		o := strings.Split(a, ")")
		g.Add(o[0], o[1])
	}

	fmt.Println(g)

	got := CountOrbits(g, "COM", 0)
	if got != 42 {
		t.Errorf("Got %d, expected %d", got, 42)
	}
}

func TestTransfer(t *testing.T) {
	test := []string{
		"COM)B",
		"B)C",
		"C)D",
		"D)E",
		"E)F",
		"B)G",
		"G)H",
		"D)I",
		"E)J",
		"J)K",
		"K)L",
		"K)YOU",
		"I)SAN",
	}
	var g Graph
	for _, a := range test {
		o := strings.Split(a, ")")
		g.Add(o[0], o[1])
	}

	fmt.Println(g)

	got := CalcTransfers(g, "YOU", "SAN")
	if got != 4 {
		t.Errorf("Got %d, expected %d", got, 4)
	}
}
