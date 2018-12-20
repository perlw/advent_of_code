package main

import (
	"bufio"
	"fmt"
	"os"
)

type ArgumentType int

const (
	ArgumentTypeImmediate ArgumentType = iota
	ArgumentTypeRegister
)

type Opcode struct {
	Name string
	A    ArgumentType
	B    ArgumentType
	Exe  func(a, b int) int
}

var opcodes = []Opcode{
	{
		Name: "mulr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  mul,
	},
}

func mul(a, b int) int {
	return a * b
}

func waitEnter() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

type Test struct {
	Before [4]int
	Op     [4]int
	After  [4]int
}

type Puzzle1 struct {
	Tests []Test
}

func (p *Puzzle1) Run(pretty bool) int {
	for _, test := range p.Tests {
		for _, op := range opcodes {
			registers := [4]int{}
			copy(registers[:], test.Before[:])

			a := test.Op[1]
			b := test.Op[2]
			if op.A == ArgumentTypeRegister {
				a = registers[a]
			}
			if op.B == ArgumentTypeRegister {
				b = registers[b]
			}

			fmt.Printf("\n%s\n", op.Name)
			registers[test.Op[3]] = op.Exe(a, b)
			fmt.Println("got", registers, "expected", test.After)
		}
	}

	return 0
}

func main() {
	input, err := os.Open("input1.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	p := Puzzle1{}

	test := Test{}
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	fmt.Sscanf(scanner.Text(), "Before: [%d, %d, %d, %d]", &test.Before[0], &test.Before[1], &test.Before[2], &test.Before[3])
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	fmt.Sscanf(scanner.Text(), "%d %d %d %d", &test.Op[0], &test.Op[1], &test.Op[2], &test.Op[3])
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	fmt.Sscanf(scanner.Text(), "After: [%d, %d, %d, %d]", &test.After[0], &test.After[1], &test.After[2], &test.After[3])
	p.Tests = append(p.Tests, test)

	fmt.Println(p.Run(true))
}
