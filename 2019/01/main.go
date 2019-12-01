package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func calcFuel(mass int) int {
	return (mass / 3) - 2
}

type Module struct {
	Mass int
}

func (m Module) RequiredFuel() (int, int) {
	base := calcFuel(m.Mass)
	var extra int
	e := calcFuel(base)
	for e >= 0 {
		extra += e
		e = calcFuel(e)
	}
	return base, extra
}

type ModuleStack []Module

func (ms ModuleStack) RequiredFuel() (int, int) {
	var base, extra int
	for _, m := range ms {
		b, e := m.RequiredFuel()
		base += b
		extra += e
	}
	return base, extra
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("could not open input file: %w", err))
	}
	defer input.Close()

	var stack ModuleStack
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		var mass int
		fmt.Sscanf(scanner.Text(), "%d", &mass)
		stack = append(stack, Module{
			Mass: mass,
		})
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(fmt.Errorf("error while scanning input: %w", err))
	}

	b, e := stack.RequiredFuel()
	log.Printf("required fuel: %d, %d => %d", b, e, b+e)
}
