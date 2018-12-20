package main

import (
	"bufio"
	"fmt"
	"os"
)

type ArgumentType int

const (
	ArgumentTypeNone ArgumentType = iota
	ArgumentTypeRegister
	ArgumentTypeValue
)

type Opcode struct {
	Name string
	A    ArgumentType
	B    ArgumentType
	Exe  func(a, b int) int
}

var opcodes = []Opcode{
	{
		Name: "addr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  add,
	},
	{
		Name: "addi",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  add,
	},
	{
		Name: "mulr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  mul,
	},
	{
		Name: "muli",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  mul,
	},
	{
		Name: "banr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  ban,
	},
	{
		Name: "bani",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  ban,
	},
	{
		Name: "borr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  bor,
	},
	{
		Name: "bori",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  bor,
	},
	{
		Name: "setr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeNone,
		Exe:  set,
	},
	{
		Name: "seti",
		A:    ArgumentTypeValue,
		B:    ArgumentTypeNone,
		Exe:  set,
	},
	{
		Name: "gtir",
		A:    ArgumentTypeValue,
		B:    ArgumentTypeRegister,
		Exe:  gt,
	},
	{
		Name: "gtri",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  gt,
	},
	{
		Name: "gtrr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  gt,
	},
	{
		Name: "eqir",
		A:    ArgumentTypeValue,
		B:    ArgumentTypeRegister,
		Exe:  eq,
	},
	{
		Name: "eqri",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  eq,
	},
	{
		Name: "eqrr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  eq,
	},
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func ban(a, b int) int {
	return a & b
}

func bor(a, b int) int {
	return a | b
}

func set(a, b int) int {
	return a
}

func gt(a, b int) int {
	if a > b {
		return 1
	}
	return 0
}

func eq(a, b int) int {
	if a == b {
		return 1
	}
	return 0
}

func waitEnter() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

type Test struct {
	Before [4]int
	Op     [4]int
	After  [4]int
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for t := range a {
		if a[t] != b[t] {
			return false
		}
	}
	return true
}

type Puzzle1 struct {
	Tests []Test
}

func (p *Puzzle1) Run() {
	total := 0
	for t, test := range p.Tests {
		potentials := 0
		fmt.Printf("Test #%d\n\tI: %+v\n\tOP:%+v\n\tT: %+v\n", t, test.Before, test.Op, test.After)
		for _, op := range opcodes {
			registers := [4]int{}
			copy(registers[:], test.Before[:])

			a := test.Op[1]
			b := test.Op[2]
			c := test.Op[3]
			if op.A == ArgumentTypeRegister {
				a = registers[a]
			}
			if op.B == ArgumentTypeRegister {
				b = registers[b]
			}

			registers[c] = op.Exe(a, b)
			fmt.Printf("\t%s -> A:%d B:%d C:%d R:%+v", op.Name, a, b, c, registers)
			if sliceEqual(registers[:], test.After[:]) {
				fmt.Printf(" OK")
				potentials++
			}
			fmt.Printf("\n")
		}
		fmt.Println("Potentials:", potentials)
		if potentials >= 3 {
			total++
		}
	}
	fmt.Println("\nTotal samples with more >=3 potentials:", total)
}

func main() {
	input, err := os.Open("input1.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	p := Puzzle1{}
	for scanner.Scan() {
		test := Test{}
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

		scanner.Scan()
		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	p.Run()
}
